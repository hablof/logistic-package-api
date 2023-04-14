package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hablof/logistic-package-api/internal/app/repo"
	"github.com/hablof/logistic-package-api/internal/app/retranslator"
	"github.com/hablof/logistic-package-api/internal/app/sender"
	"github.com/hablof/logistic-package-api/internal/config"
	"github.com/hablof/logistic-package-api/internal/database"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Retranslator.Debug). // !!
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Retranslator.Name)

	if cfg.Retranslator.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	db, err := database.NewPostgres(cfg.Database, cfg.Database.Driver)
	if err != nil {
		log.Error().Err(err).Msg("Failed init postgres")
		return
	}
	defer db.Close()
	log.Info().Err(err).Msg("database ping attempt succeeded")

	r := repo.NewRepository(db)

	var (
		kp *sender.KafkaProducer
		// err error
	)
	for i := 0; i < cfg.Kafka.MaxAttempts; i++ {
		kp, err = sender.NewKafkaProducer(cfg.Kafka)

		if err == nil {
			break
		}

		log.Info().Err(err).Msgf("NewKafkaProducer: failed attempt %d/%d to connect to kafka", i+1, cfg.Kafka.MaxAttempts)
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		log.Error().Err(err).Msg("Failed init kafka")
		return
	}
	log.Info().Err(err).Msg("NewKafkaProducer: connected to kafka")

	retranslatorConfig := retranslator.RetranslatorConfig{
		ChannelSize:     uint64(cfg.Retranslator.ChannelSize),
		ConsumerCount:   uint64(cfg.Retranslator.ConsumerCount),
		BatchSize:       uint64(cfg.Retranslator.BatchSize),
		ConsumeInterval: cfg.Retranslator.ConsumeInterval,
		ProducerCount:   uint64(cfg.Retranslator.ProducerCount),
		WorkerCount:     cfg.Retranslator.WorkerCount,
		CleanerRepo:     r,
		ConsumerRepo:    r,
		Sender:          kp,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	retranslator := retranslator.NewRetranslator(retranslatorConfig)

	go func() {
		retranslator.Start()
	}()

	metricsServer := createMetricsServer(&cfg)

	go func() {
		log.Info().Msgf("Metrics server is running on %s", cfg.Retranslator.MetricsAddr)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running metrics server")
			cancel()
		}
	}()

	select {
	case v := <-sigs:
		log.Info().Msgf("signal.Notify: %v", v)

	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v", done)
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("metricsServer.Shutdown")
	} else {
		log.Info().Msg("metricsServer shut down correctly")
	}

	retranslator.Close()
}

func createMetricsServer(cfg *config.Config) *http.Server {
	// addr := fmt.Sprintf("%s:%d", cfg.R.Host, cfg.Metrics.Port)

	mux := http.DefaultServeMux
	mux.Handle(cfg.Retranslator.MetricsPath, promhttp.Handler())

	metricsServer := &http.Server{
		Addr:              cfg.Retranslator.MetricsAddr,
		Handler:           mux,
		ReadHeaderTimeout: 500 * time.Millisecond,
	}

	return metricsServer
}

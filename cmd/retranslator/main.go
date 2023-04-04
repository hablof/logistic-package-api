package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hablof/logistic-package-api/internal/app/repo"
	"github.com/hablof/logistic-package-api/internal/app/retranslator"
	"github.com/hablof/logistic-package-api/internal/app/sender"
	"github.com/hablof/logistic-package-api/internal/config"
	"github.com/hablof/logistic-package-api/internal/database"

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
		log.Fatal().Err(err).Msg("Failed init postgres")
		return
	}
	defer db.Close()

	r := repo.NewRepository(db)

	kp, err := sender.NewKafkaProducer(cfg.Kafka)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init kafka")
		return
	}

	rtrCFG := retranslator.RetranslatorConfig{
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

	retranslator := retranslator.NewRetranslator(rtrCFG)
	retranslator.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	retranslator.Close()
}

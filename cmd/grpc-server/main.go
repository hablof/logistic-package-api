package main

import (
	// "github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/hablof/logistic-package-api/internal/config"
	"github.com/hablof/logistic-package-api/internal/database"
	"github.com/hablof/logistic-package-api/internal/server"
	"github.com/hablof/logistic-package-api/internal/tracer"
)

var (
	batchSize uint64 = 2
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	// migration := flag.Bool("migration", true, "Defines the migration start option")
	// flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	/*
	   In addition to global logger, in project there is
	   local logger in gRPC api,
	   that allows rise log level via gRPC metadata.

	   To provide correct work of local loggers,
	   global level should be "DebugLevel"
	*/
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// if cfg.Project.Debug {
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// } else {
	// 	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// }

	db, err := database.NewPostgres(cfg.Database, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
		return
	}
	defer db.Close()

	// *migration = false // todo: need to delete this line for homework-4
	// if *migration {
	// 	if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
	// 		log.Error().Err(err).Msg("Migration failed")

	// 		return
	// 	}
	// }

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed init tracing")

		return
	}
	defer tracing.Close()

	if err := server.NewGrpcServer(db /* , batchSize */).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}

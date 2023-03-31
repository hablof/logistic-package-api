package retranslator

import (
	"time"

	"github.com/hablof/logistic-package-api/internal/app/cleaner"
	"github.com/hablof/logistic-package-api/internal/app/consumer"
	"github.com/hablof/logistic-package-api/internal/app/producer"
	"github.com/hablof/logistic-package-api/internal/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

var (
	totalRetranslatorEvents = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "logistic_package_api_retranslator_events_processing",
		Help: "Total number of events processing in retranslator",
	})
)

type Retranslator interface {
	Start()
	Close()
}

type RetranslatorConfig struct {
	ChannelSize uint64

	ConsumerCount   uint64
	BatchSize       uint64
	ConsumeInterval time.Duration

	ProducerCount uint64
	WorkerCount   int

	CleanerRepo  cleaner.RepoEventCleaner
	ConsumerRepo consumer.RepoEventConsumer
	Sender       producer.EventSender
}

type retranslator struct {
	events   chan model.PackageEvent
	consumer consumer.Consumer
	producer producer.Producer
	cleaner  cleaner.Cleaner
}

func NewRetranslator(cfg RetranslatorConfig) Retranslator {
	eventsChannel := make(chan model.PackageEvent, cfg.ChannelSize)
	cleanerChannel := make(chan cleaner.PackageCleanerEvent, cfg.ChannelSize)

	consumerCfg := consumer.ConsumerConfig{
		ConsumeCount:    cfg.ConsumerCount,
		EventsChannel:   eventsChannel,
		Repo:            cfg.ConsumerRepo,
		BatchSize:       cfg.BatchSize,
		ConsumeInterval: cfg.ConsumeInterval,
		GaugeAddFunc:    totalRetranslatorEvents.Add,
	}

	producerCfg := producer.ProducerConfig{
		ProducerCount:  cfg.ProducerCount,
		Sender:         cfg.Sender,
		EventsChannel:  eventsChannel,
		CleanerChannel: cleanerChannel,
	}

	cleanerCfg := cleaner.CleanerConfig{
		WorkerCount:      cfg.WorkerCount,
		CleanerBatchSize: cfg.BatchSize / 2,
		Repo:             cfg.CleanerRepo,
		CleanerChannel:   cleanerChannel,
		CleanupInterval:  cfg.ConsumeInterval,
		GaugeSubFunc:     totalRetranslatorEvents.Sub,
	}

	consumer := consumer.NewDbConsumer(consumerCfg)
	producer := producer.NewKafkaProducer(producerCfg)
	cleaner := cleaner.NewDbCleaner(cleanerCfg)

	return &retranslator{
		events:   eventsChannel,
		consumer: consumer,
		producer: producer,
		cleaner:  cleaner,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
	r.cleaner.Start()
	log.Debug().Msg("retranslator started")
}

func (r *retranslator) Close() {
	// closing sequence matters to
	// implement At-least-once guarantee
	// consumer -> producer -> cleaner
	r.consumer.Close()
	r.producer.Close()
	r.cleaner.Close()
}

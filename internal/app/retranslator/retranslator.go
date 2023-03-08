package retranslator

import (
	"log"
	"time"

	"github.com/hablof/logistic-package-api/internal/app/cleaner"
	"github.com/hablof/logistic-package-api/internal/app/consumer"
	"github.com/hablof/logistic-package-api/internal/app/producer"
	"github.com/hablof/logistic-package-api/internal/app/repo"
	"github.com/hablof/logistic-package-api/internal/app/sender"
	"github.com/hablof/logistic-package-api/internal/model"
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

	Repo   repo.EventRepo
	Sender sender.EventSender
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
		Repo:            cfg.Repo,
		BatchSize:       cfg.BatchSize,
		ConsumeInterval: cfg.ConsumeInterval,
	}

	producerCfg := producer.ProducerConfig{
		ProducerCount:  cfg.ProducerCount,
		Repo:           cfg.Repo,
		Sender:         cfg.Sender,
		EventsChannel:  eventsChannel,
		CleanerChannel: cleanerChannel,
	}

	cleanerCfg := cleaner.CleanerConfig{
		WorkerCount:      cfg.WorkerCount,
		CleanerBatchSize: cfg.BatchSize / 2,
		Repo:             cfg.Repo,
		CleanerChannel:   cleanerChannel,
		CleanupInterval:  cfg.ConsumeInterval,
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
	log.Println("retranslator started")
}

func (r *retranslator) Close() {
	// closing order matters to
	// implement At-least-once guarantee
	// consumer -> producer -> cleaner
	r.consumer.Close()
	r.producer.Close()
	r.cleaner.Close()
}

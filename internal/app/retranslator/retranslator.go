package retranslator

import (
	"time"

	"github.com/hablof/logistic-package-api/internal/app/consumer"
	"github.com/hablof/logistic-package-api/internal/app/producer"
	"github.com/hablof/logistic-package-api/internal/app/repo"
	"github.com/hablof/logistic-package-api/internal/app/sender"
	"github.com/hablof/logistic-package-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Retranslator interface {
	Start()
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.PackageEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg Config) Retranslator {
	eventsChannel := make(chan model.PackageEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumerCfg := consumer.ConsumerConfig{
		ConsumeCount:  cfg.ConsumerCount,
		EventsChannel: eventsChannel,
		Repo:          cfg.Repo,
		BatchSize:     cfg.ConsumeSize,
		Timeout:       cfg.ConsumeTimeout,
	}

	producerCfg := producer.ProducerConfig{
		ProducerCount: cfg.ProducerCount,
		Repo:          cfg.Repo,
		Sender:        cfg.Sender,
		EventsChannel: eventsChannel,
		WorkerPool:    workerPool,
	}

	consumer := consumer.NewDbConsumer(consumerCfg)
	producer := producer.NewKafkaProducer(producerCfg)

	return &retranslator{
		events:     eventsChannel,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}

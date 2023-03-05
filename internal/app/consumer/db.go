package consumer

import (
	"sync"
	"time"

	"github.com/hablof/logistic-package-api/internal/app/repo"
	"github.com/hablof/logistic-package-api/internal/model"
)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	consumerCount uint64
	events        chan<- model.PackageEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	done chan bool
	wg   *sync.WaitGroup
}

type ConsumerConfig struct {
	ConsumeCount uint64
	Events       chan<- model.PackageEvent
	Repo         repo.EventRepo
	BatchSize    uint64
	Timeout      time.Duration
}

func NewDbConsumer(cfg ConsumerConfig) Consumer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &consumer{
		consumerCount: cfg.ConsumeCount,
		batchSize:     cfg.BatchSize,
		timeout:       cfg.Timeout,
		repo:          cfg.Repo,
		events:        cfg.Events,
		wg:            wg,
		done:          done, // use ctx
	}
}

func (c *consumer) Start() {
	for i := uint64(0); i < c.consumerCount; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-c.done:
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	close(c.done)
	c.wg.Wait()
}

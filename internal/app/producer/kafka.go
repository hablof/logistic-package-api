package producer

import (
	"sync"

	"github.com/hablof/logistic-package-api/internal/app/repo"
	"github.com/hablof/logistic-package-api/internal/app/sender"
	"github.com/hablof/logistic-package-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	producerCount uint64

	repo repo.EventRepo

	sender        sender.EventSender
	eventsChannel <-chan model.PackageEvent

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool // use ctx
}

type ProducerConfig struct {
	ProducerCount uint64
	Repo          repo.EventRepo
	Sender        sender.EventSender
	EventsChannel <-chan model.PackageEvent
	WorkerPool    *workerpool.WorkerPool
}

func NewKafkaProducer(cfg ProducerConfig) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		producerCount: cfg.ProducerCount,
		repo:          cfg.Repo,
		sender:        cfg.Sender,
		eventsChannel: cfg.EventsChannel,
		workerPool:    cfg.WorkerPool,
		wg:            wg,
		done:          done, // use ctx
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.producerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.eventsChannel:
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {
							// ...
						})
					} else {
						p.workerPool.Submit(func() {
							// ...
						})
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}

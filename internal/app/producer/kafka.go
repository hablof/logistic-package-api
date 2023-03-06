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

func (p *producer) handle() {
	for {
		select {
		case event := <-p.eventsChannel: // TODO end work on channel close
			switch err := p.sender.Send(&event); err {
			case nil:
				p.workerPool.Submit(func() {
					p.repo.Remove([]uint64{event.ID})
				})

			default:
				p.workerPool.Submit(func() {
					p.repo.Unlock([]uint64{event.ID})
				})
			}
		case <-p.done:
			return
		}
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.producerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.handle()
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}

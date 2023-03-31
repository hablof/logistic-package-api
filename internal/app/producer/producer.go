package producer

import (
	"context"
	"sync"
	"time"

	"github.com/hablof/logistic-package-api/internal/app/cleaner"
	"github.com/hablof/logistic-package-api/internal/model"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	Start()
	Close()
}

type EventSender interface {
	Send(subdomain *model.PackageEvent) error
}

type producer struct {
	producerCount uint64

	cleanerChannel chan<- cleaner.PackageCleanerEvent

	// maximumKeepOrderAttempts model.TimesDefered
	// orderManager             ordermanager.OrderManager
	sender        EventSender
	eventsChannel chan model.PackageEvent

	wg     *sync.WaitGroup
	cancel context.CancelFunc
}

type ProducerConfig struct {
	// maximumKeepOrderAttempts model.TimesDefered
	ProducerCount  uint64
	Sender         EventSender
	CleanerChannel chan<- cleaner.PackageCleanerEvent
	EventsChannel  chan model.PackageEvent
}

func NewKafkaProducer(cfg ProducerConfig) Producer {

	wg := &sync.WaitGroup{}
	// orderManager := ordermanager.NewOrderManager()

	return &producer{
		producerCount:  cfg.ProducerCount,
		cleanerChannel: cfg.CleanerChannel,
		// maximumKeepOrderAttempts: cfg.maximumKeepOrderAttempts,
		// orderManager:             orderManager,
		sender:        cfg.Sender,
		eventsChannel: cfg.EventsChannel,
		wg:            wg,
		cancel: func() {
		},
	}
}

// TODO check event.Entity != nil
func (p *producer) runHandler(ctx context.Context) {
	for {
		select {
		case event := <-p.eventsChannel:
			// if !p.orderManager.ApproveOrder(event) {
			// 	if event.Defered < p.maximumKeepOrderAttempts {
			// 		event.Defered++
			// 		p.eventsChannel <- event
			// 		continue // !!!
			// 	}

			// 	log.("failed to keep right order with event %v", event)
			// }

			switch err := p.sender.Send(&event); err {
			case nil:
				// if err := p.orderManager.RegisterEvent(event); err != nil {
				// 	log.("event %v registration in ordermanager error: %v", event, err)
				// }

				p.cleanerChannel <- cleaner.PackageCleanerEvent{
					Status:  cleaner.Ok,
					EventID: event.ID,
				}

			default:
				log.Debug().Err(err).Msgf("eventID: %d", event.ID)
				p.cleanerChannel <- cleaner.PackageCleanerEvent{
					Status:  cleaner.Fail,
					EventID: event.ID,
				}
			}

		case <-ctx.Done():
			return
		}
	}
}

func (p *producer) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel

	for i := uint64(0); i < p.producerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.runHandler(ctx)
		}()
	}

	log.Info().Msgf("producer started with %d workers", p.producerCount)
}

func (p *producer) Close() {

	// if producer.Close() called after consumer close finished, new entity in channel will never occures
	// len(c.eventsChannel) == 0 means handler return by <- ctx.Done() implements At-least-once guarantee

	for len(p.eventsChannel) != 0 {
		time.Sleep(250 * time.Millisecond)
	}

	p.cancel()
	p.wg.Wait()
}

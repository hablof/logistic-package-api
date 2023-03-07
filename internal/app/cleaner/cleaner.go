package cleaner

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/hablof/logistic-package-api/internal/app/repo"
)

type SenderStatus uint8

const (
	Ok SenderStatus = iota
	Fail
)

type PackageCleanerEvent struct {
	Status  SenderStatus
	EventID uint64
}

type Cleaner interface {
	Start()
	Close()
}

type cleaner struct {
	cleanerChannel        <-chan PackageCleanerEvent
	repo                  repo.EventRepo
	batchSize             uint64
	forcedCleanupInterval time.Duration
	workerPool            *workerpool.WorkerPool

	// removeQueue []uint64
	// unlockQueue []uint64

	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

type CleanerConfig struct {
	WorkerCount      int
	CleanerBatchSize uint64
	Repo             repo.EventRepo
	CleanerChannel   <-chan PackageCleanerEvent
	CleanupInterval  time.Duration
}

func NewDbCleaner(cfg CleanerConfig) Cleaner {
	wp := workerpool.New(cfg.WorkerCount)
	wg := &sync.WaitGroup{}

	return &cleaner{
		cleanerChannel:        cfg.CleanerChannel,
		repo:                  cfg.Repo,
		batchSize:             cfg.CleanerBatchSize,
		forcedCleanupInterval: cfg.CleanupInterval,
		workerPool:            wp,
		cancel: func() {
		},
		wg: wg,
	}
}

// TODO refactor this
func (c *cleaner) runHandler(ctx context.Context) {
	removeQueue := make([]uint64, 0, c.batchSize)
	unlockQueue := make([]uint64, 0, c.batchSize)
	ticker := time.NewTicker(c.forcedCleanupInterval)

	for {
		select {
		case <-ticker.C:
			if len(removeQueue) > 0 {
				c.workerPool.Submit(func() {
					if err := c.repo.Remove(removeQueue); err != nil {
						log.Println(err)
					}
				})
				removeQueue = make([]uint64, 0, c.batchSize)
			}

			if len(unlockQueue) > 0 {
				c.workerPool.Submit(func() {
					if err := c.repo.Unlock(unlockQueue); err != nil {
						log.Println(err)
					}
				})
				unlockQueue = make([]uint64, 0, c.batchSize)
			}

		case event := <-c.cleanerChannel:
			switch event.Status {
			case Ok:
				removeQueue = append(removeQueue, event.EventID)
			case Fail:
				unlockQueue = append(unlockQueue, event.EventID)
			}

			if len(removeQueue) >= int(c.batchSize) {
				c.workerPool.Submit(func() {
					if err := c.repo.Remove(removeQueue); err != nil {
						log.Println(err)
					}
				})
				removeQueue = make([]uint64, 0, c.batchSize)
			}

			if len(unlockQueue) >= int(c.batchSize) {
				c.workerPool.Submit(func() {
					if err := c.repo.Unlock(unlockQueue); err != nil {
						log.Println(err)
					}
				})
				unlockQueue = make([]uint64, 0, c.batchSize)
			}
			ticker.Reset(c.forcedCleanupInterval)

		case <-ctx.Done():
			ticker.Stop()
			if len(removeQueue) > 0 {
				c.workerPool.Submit(func() {
					if err := c.repo.Remove(removeQueue); err != nil {
						log.Println(err)
					}
				})
				removeQueue = make([]uint64, 0, c.batchSize)
			}

			if len(unlockQueue) > 0 {
				c.workerPool.Submit(func() {
					if err := c.repo.Unlock(unlockQueue); err != nil {
						log.Println(err)
					}
				})
				unlockQueue = make([]uint64, 0, c.batchSize)
			}
			return
		}
	}
}

func (c *cleaner) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	for i := 0; i < c.workerPool.Size(); i++ {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			c.runHandler(ctx)
		}()
	}

}

func (c *cleaner) Close() {

	// if cleaner.Close() called after producer close finished, new entity in channel will never occures
	// len(c.cleanerChannel) == 0 means handler return by <- ctx.Done() implements At-least-once guarantee

	for len(c.cleanerChannel) != 0 {
		time.Sleep(250 * time.Millisecond)
	}
	c.cancel()
	c.wg.Wait()

	c.workerPool.StopWait()
}

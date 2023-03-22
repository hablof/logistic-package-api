package cleaner

import (
	"context"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/rs/zerolog/log"
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

type RepoEventCleaner interface {
	Unlock(eventIDs []uint64) error
	Remove(eventIDs []uint64) error
}

type Cleaner interface {
	Start()
	Close()
}

type cleaner struct {
	cleanerCount          int
	cleanerChannel        <-chan PackageCleanerEvent
	repo                  RepoEventCleaner
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
	Repo             RepoEventCleaner
	CleanerChannel   <-chan PackageCleanerEvent
	CleanupInterval  time.Duration
}

func NewDbCleaner(cfg CleanerConfig) Cleaner {
	wp := workerpool.New(cfg.WorkerCount)
	wg := &sync.WaitGroup{}

	return &cleaner{
		cleanerCount:          cfg.WorkerCount,
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
				c.submitToRemove(removeQueue)
				removeQueue = make([]uint64, 0, c.batchSize)
			}

			if len(unlockQueue) > 0 {
				c.submitToUnlock(unlockQueue)
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
				c.submitToRemove(removeQueue)
				removeQueue = make([]uint64, 0, c.batchSize)
			}

			if len(unlockQueue) >= int(c.batchSize) {
				c.submitToUnlock(unlockQueue)
				unlockQueue = make([]uint64, 0, c.batchSize)
			}
			ticker.Reset(c.forcedCleanupInterval)

		case <-ctx.Done():
			ticker.Stop()
			if len(removeQueue) > 0 {
				c.submitToRemove(removeQueue)
			}

			if len(unlockQueue) > 0 {
				c.submitToUnlock(unlockQueue)
			}
			return
		}
	}
}

func (c *cleaner) submitToUnlock(unlockQueue []uint64) {
	sending := make([]uint64, len(unlockQueue))
	copy(sending, unlockQueue)
	c.workerPool.Submit(func() {
		if err := c.repo.Unlock(sending); err != nil {
			log.Debug().Err(err).Msg("cleaner.submitToUnlock failed")
		}
	})
}

func (c *cleaner) submitToRemove(removeQueue []uint64) {
	sending := make([]uint64, len(removeQueue))
	copy(sending, removeQueue)
	c.workerPool.Submit(func() {
		if err := c.repo.Remove(sending); err != nil {
			log.Debug().Err(err).Msg("cleaner.submitToRemove failed")
		}
	})
}

func (c *cleaner) Start() {
	ctx, cf := context.WithCancel(context.Background())
	c.cancel = cf
	for i := 0; i < c.cleanerCount; i++ {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			c.runHandler(ctx)
		}()
	}

	log.Info().Msgf("cleaner started with %d workers", c.cleanerCount)
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

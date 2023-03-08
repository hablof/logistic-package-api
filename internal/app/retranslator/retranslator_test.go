package retranslator

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hablof/logistic-package-api/internal/mocks"
	"github.com/hablof/logistic-package-api/internal/model"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestStart(t *testing.T) {

	t.Run("correct run and stop", func(t *testing.T) {
		repo, _, retranslator := setup(t, 2)

		repo.EXPECT().Lock(gomock.Any()).AnyTimes()

		retranslator.Start()
		retranslator.Close()
		t.Log("correct run and stop PASSED")
	})

	t.Run("correctly read all events and send", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Run(fmt.Sprintf("attempt %d", i), func(t *testing.T) {
				batchSize := 32
				repo, sender, retranslator := setup(t, uint64(batchSize))

				controlChannel := make(chan struct{}, 1)

				eventsCount := 5000
				db := generate(eventsCount)

				offsetMutex := sync.Mutex{}
				offset := uint64(0)

				sendCount := int32(0)

				removedIDsMutex := sync.Mutex{}
				removedIDs := make([]uint64, 0, eventsCount)

				repo.EXPECT().Lock(gomock.Any()).DoAndReturn(func(size uint64) ([]model.PackageEvent, error) {
					offsetMutex.Lock()
					defer offsetMutex.Unlock()
					if offset >= uint64(len(db)) {
						return make([]model.PackageEvent, 0), nil
					}
					// chunk := db[offset : offset+size : offset+size]
					if offset+size >= uint64(len(db)) {
						chunk := db[offset:]
						offset += size
						return chunk, nil
					}
					chunk := db[offset : offset+size]
					offset += size
					return chunk, nil
				}).AnyTimes()

				repo.EXPECT().Remove(gomock.Any()).AnyTimes().Do(func(arr []uint64) {
					removedIDsMutex.Lock()
					removedIDs = append(removedIDs, arr...)
					if len(removedIDs) == eventsCount {
						controlChannel <- struct{}{}
					}
					removedIDsMutex.Unlock()
				})

				sender.EXPECT().Send(gomock.Any()).Times(eventsCount).Do(func(ptr *model.PackageEvent) {
					atomic.AddInt32(&sendCount, 1)
				}).Return(nil)

				retranslator.Start()
				go func() {
					time.Sleep(10 * time.Second)
					controlChannel <- struct{}{}
				}()
				<-controlChannel
				// removedIDsMutex.Lock()
				// sort.Slice(removedIDs, func(i, j int) bool { return removedIDs[i] < removedIDs[j] })
				// removedIDsMutex.Unlock()
				retranslator.Close()

				assert.Equal(t, int32(eventsCount), sendCount)
				assert.Equal(t, eventsCount, len(removedIDs), "len of removedIDs")
				assert.Equal(t, true, checkArr(removedIDs, uint64(eventsCount)))
			})
		}
	})

	// t.Run("correctly reprocess events again if error", func(t *testing.T) {
	// 	chuckSize := 10
	// 	repo, sender, retranslator := setup(t, uint64(chuckSize))

	// 	// arrange
	// 	eventsCount := 100

	// 	db := generate(eventsCount)
	// 	offset := uint64(0)

	// 	allRead := make(chan bool, 10)

	// 	sendCount := int32(0)

	// 	reprocessCount := int32(0)

	// 	repo.
	// 		EXPECT().
	// 		Lock(gomock.Any()).
	// 		DoAndReturn(func(size uint64) ([]model.PackageEvent, error) {
	// 			if offset >= uint64(len(db)) {
	// 				return make([]model.PackageEvent, 0), nil
	// 			}
	// 			maxIndex := uint64(math.Min(float64(offset+size), float64(len(db))))
	// 			chunk := db[offset:maxIndex:maxIndex]
	// 			atomic.AddUint64(&offset, maxIndex-offset)
	// 			return chunk, nil
	// 		}).
	// 		AnyTimes()

	// 	repo.
	// 		EXPECT().
	// 		Unlock(gomock.Any()).
	// 		DoAndReturn(func(eventIDs []uint64) error {
	// 			atomic.AddUint64(&offset, -uint64(len(eventIDs)))
	// 			atomic.AddInt32(&reprocessCount, 1)
	// 			return nil
	// 		}).AnyTimes()

	// 	repo.EXPECT().Remove(gomock.Any()).AnyTimes()

	// 	sender.
	// 		EXPECT().
	// 		Send(gomock.Any()).
	// 		DoAndReturn(func(event *model.PackageEvent) error {
	// 			atomic.AddInt32(&sendCount, 1)
	// 			for sendCount <= int32(eventsCount) {
	// 				return errors.New("Error has occurred when send to kafka")
	// 			}

	// 			if sendCount == int32(eventsCount*2) {
	// 				allRead <- true
	// 			}
	// 			return nil
	// 		}).
	// 		AnyTimes()

	// 	retranslator.Start()

	// 	wg := sync.WaitGroup{}
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for {
	// 			<-allRead
	// 			retranslator.Close()
	// 			return
	// 		}
	// 	}()

	// 	wg.Wait()
	// 	assert.Equal(t, int32(eventsCount*2), sendCount)
	// 	assert.Equal(t, int32(eventsCount), reprocessCount)
	// 	t.Log("correctly reprocess events again if error PASSED")
	// })
}

func generate(count int) []model.PackageEvent {
	result := make([]model.PackageEvent, 0, count)
	for i := 0; i < count; i++ {
		event := model.PackageEvent{
			ID:      uint64(i),
			Type:    model.Created,
			Status:  0,
			Defered: 0,
			Entity: &model.Package{
				ID:            uint64(i),
				Title:         "",
				Material:      "",
				MaximumVolume: 0,
				Reusable:      false,
			},
		}
		result = append(result, event)
	}

	return result
}

func setup(t *testing.T, batchSize uint64) (*mocks.MockEventRepo, *mocks.MockEventSender, Retranslator) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := RetranslatorConfig{
		ChannelSize:     512,
		ConsumerCount:   16,
		BatchSize:       batchSize,
		ConsumeInterval: 100 * time.Millisecond,
		ProducerCount:   8,
		WorkerCount:     4,
		Repo:            repo,
		Sender:          sender,
	}

	retranslator := NewRetranslator(cfg)

	return repo, sender, retranslator
}

func checkArr(arr []uint64, n uint64) bool {
e:
	for i := uint64(0); i < n; i++ {
		for _, elem := range arr {
			if elem == i {
				continue e
			}
		}
		log.Printf("arr missed %d", i)
		return false
	}

	return true
}

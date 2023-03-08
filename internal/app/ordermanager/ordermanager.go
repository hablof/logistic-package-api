package ordermanager

import (
	"errors"
	"sync"

	"github.com/hablof/logistic-package-api/internal/model"
)

var (
	ErrIncorrectOrder = errors.New("has registred with incorrect order")
)

type OrderManager interface {
	ApproveOrder(incomingEvent model.PackageEvent) bool
	RegisterEvent(incomingEvent model.PackageEvent) error
}

func NewOrderManager() OrderManager {
	return &orderManager{
		mu:       sync.Mutex{},
		ordermap: map[uint64]model.EventType{},
	}
}

type orderManager struct {
	mu       sync.Mutex
	ordermap map[uint64]model.EventType // PackageEvent.Entity.ID -> PackageEvent.EventType
}

func (o *orderManager) ApproveOrder(incomingEvent model.PackageEvent) bool {
	o.mu.Lock()
	defer o.mu.Unlock()

	prevEventType, ok := o.ordermap[incomingEvent.Entity.ID]

	switch incomingEvent.Type {
	case model.Created:
		if !ok {
			return true
		}

	case model.Updated, model.Removed:
		if ok && (prevEventType == model.Created || prevEventType == model.Updated) {
			return true
		}
	}

	return false
}

func (o *orderManager) RegisterEvent(incomingEvent model.PackageEvent) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	prevEventType, ok := o.ordermap[incomingEvent.Entity.ID]

	switch incomingEvent.Type {
	case model.Created:
		o.ordermap[incomingEvent.Entity.ID] = model.Created
		if !ok {
			return nil
		}

	case model.Updated:
		o.ordermap[incomingEvent.Entity.ID] = model.Updated
		if ok && (prevEventType == model.Created || prevEventType == model.Updated) {
			return nil
		}

	case model.Removed:
		delete(o.ordermap, incomingEvent.Entity.ID)
		if ok && (prevEventType == model.Created || prevEventType == model.Updated) {
			return nil
		}
	}

	return ErrIncorrectOrder
}

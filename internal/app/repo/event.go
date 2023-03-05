package repo

import (
	"github.com/hablof/logistic-package-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.PackageEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.PackageEvent) error // not used
	Remove(eventIDs []uint64) error
}

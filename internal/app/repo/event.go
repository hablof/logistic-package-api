package repo

import (
	"github.com/ozonmp/omp-demo-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.PackageEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.PackageEvent) error
	Remove(eventIDs []uint64) error
}

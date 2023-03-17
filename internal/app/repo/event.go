package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/hablof/logistic-package-api/internal/api"
	"github.com/hablof/logistic-package-api/internal/app/cleaner"
	"github.com/hablof/logistic-package-api/internal/app/consumer"
	"github.com/hablof/logistic-package-api/internal/model"
)

// type RepoEvent interface {
// 	Lock(n uint64) ([]model.PackageEvent, error)
// 	Unlock(eventIDs []uint64) error

// 	Add(event []model.PackageEvent) error // not used
// 	Remove(eventIDs []uint64) error
// }

var _ api.RepoCRUD = repository{}
var _ cleaner.RepoEventCleaner = repository{}
var _ consumer.RepoEventConsumer = repository{}

type repository struct {
	db        *sqlx.DB
	batchSize uint
	initQuery sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB, batchSize uint) *repository {
	return &repository{
		db:        db,
		batchSize: batchSize,
		initQuery: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), // Postgres format $1, $2....
	}
}

// CreatePackage implements api.RepoCRUD
func (repository) CreatePackage(ctx context.Context, pack *model.Package) (uint64, error) {

	panic("unimplemented")
}

// DescribePackage implements api.RepoCRUD
func (repository) DescribePackage(ctx context.Context, packageID uint64) (*model.Package, error) {
	panic("unimplemented")
}

// ListPackages implements api.RepoCRUD
func (repository) ListPackages(ctx context.Context, offset uint64) ([]*model.Package, error) {
	panic("unimplemented")
}

// RemovePackage implements api.RepoCRUD
func (repository) RemovePackage(ctx context.Context, packageID uint64) error {
	panic("unimplemented")
}

// Lock implements consumer.RepoEventConsumer
func (repository) Lock(n uint64) ([]model.PackageEvent, error) {
	panic("unimplemented")
}

// Remove implements cleaner.RepoEventCleaner
func (repository) Remove(eventIDs []uint64) error {
	panic("unimplemented")
}

// Unlock implements cleaner.RepoEventCleaner
func (repository) Unlock(eventIDs []uint64) error {
	panic("unimplemented")
}

package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/hablof/logistic-package-api/internal/api"
	"github.com/hablof/logistic-package-api/internal/app/cleaner"
	"github.com/hablof/logistic-package-api/internal/app/consumer"
)

// type RepoEvent interface {
// 	Lock(n uint64) ([]model.PackageEvent, error)
// 	Unlock(eventIDs []uint64) error

// 	Add(event []model.PackageEvent) error // not used
// 	Remove(eventIDs []uint64) error
// }

var _ api.RepoCRUD = &repository{}
var _ cleaner.RepoEventCleaner = &repository{}
var _ consumer.RepoEventConsumer = &repository{}

type repository struct {
	db *sqlx.DB
	// batchSize uint64
	initQuery sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB /* , batchSize uint64 */) *repository {
	return &repository{
		db: db,
		// batchSize: batchSize,
		initQuery: sq.StatementBuilder.PlaceholderFormat(sq.Dollar), // Postgres format $1, $2....
	}
}

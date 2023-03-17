package repo

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/hablof/logistic-package-api/internal/model"
)

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) *repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) CreatePackage(ctx context.Context, pack *model.Package) (uint64, error) {
	return 0, errors.New("not implemented")
}

func (r *repo) DescribePackage(ctx context.Context, packageID uint64) (*model.Package, error) {
	return nil, errors.New("not implemented")
}

func (r *repo) ListPackages(ctx context.Context, offset uint64) ([]*model.Package, error) {
	return nil, errors.New("not implemented")
}

func (r *repo) RemovePackage(ctx context.Context, packageID uint64) error {
	return errors.New("not implemented")
}

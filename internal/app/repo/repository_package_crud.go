package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/hablof/logistic-package-api/internal/api"
	"github.com/hablof/logistic-package-api/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

const (
	DefaultLimit = 10
)

// CreatePackage implements api.RepoCRUD
func (r *repository) CreatePackage(ctx context.Context, pack *model.Package) (uint64, error) {

	log.Debug().Msgf("repository.CreatePackage has called with package title: %s", pack.Title)

	crudQuery, crudArgs, err := r.initQuery.
		Insert("package").
		Columns("title", "material", "max_volume", "reusable", "created_at").
		Values(pack.Title, pack.Material, pack.MaximumVolume, pack.Reusable, "now()").
		Suffix("RETURNING package_id").
		ToSql()
	if err != nil {
		return 0, err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	log.Debug().Msgf("crud query: %s; args: %v", crudQuery, crudArgs)

	returningID := uint64(0)
	row := tx.QueryRowxContext(ctx, crudQuery, crudArgs...)
	if err := row.Scan(&returningID); err != nil {
		return 0, err
	}

	unit := model.Package{
		ID:            returningID,
		Title:         pack.Title,
		Material:      pack.Material,
		MaximumVolume: pack.MaximumVolume,
		Reusable:      pack.Reusable,
		Created:       pack.Created,
	}

	jsonbytes, err := json.Marshal(unit)
	if err != nil {
		return 0, err
	}

	eventQuery, eventArgs, err := r.initQuery.
		Insert("package_event").
		Columns("package_id", "event_type", "payload").
		Values(returningID, "Created", jsonbytes).
		ToSql()
	if err != nil {
		return 0, err
	}

	log.Debug().Msgf("event query: %s; args: %v", eventQuery, eventArgs)

	if _, err := tx.ExecContext(ctx, eventQuery, eventArgs...); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	log.Debug().Msgf("repository.CreatePackage: package_id %d inserted", returningID)

	return returningID, nil

}

// DescribePackage implements api.RepoCRUD
func (r *repository) DescribePackage(ctx context.Context, packageID uint64) (*model.Package, error) {

	log.Debug().Msgf("repository.DescribePackage has called with ID: %d", packageID)

	query, args, err := r.initQuery.
		Select("package_id", "title", "material", "max_volume", "reusable", "created_at").
		From("package").
		Where(sq.Eq{"package_id": packageID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("query: %s; args: %v", query, args)

	row := r.db.QueryRowxContext(ctx, query, args...)

	unit := model.Package{}
	switch err := row.StructScan(&unit); {
	case errors.Is(err, sql.ErrNoRows):
		return nil, api.ErrRepoEntityNotFound

	case err != nil:
		return nil, err
	}

	log.Debug().Msgf("repository.DescribePackage: package_id %d has read", packageID)

	return &unit, nil
}

// ListPackages implements api.RepoCRUD
func (r *repository) ListPackages(ctx context.Context, offset uint64) ([]model.Package, error) {

	log.Debug().Msgf("repository.ListPackages has called with offset %d", offset)

	query, args, err := r.initQuery.
		Select("package_id", "title", "material", "max_volume", "reusable", "created_at").
		From("package").
		Limit(DefaultLimit).
		Offset(offset).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("query: %s; args: %v", query, args)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	output := make([]model.Package, 0, DefaultLimit)

	for rows.Next() {
		unit := model.Package{}
		if err := rows.StructScan(&unit); err != nil {
			return nil, err
		}
		output = append(output, unit)
	}

	/****************/
	returningIDs := make([]uint64, 0, len(output))
	for _, elem := range output {
		returningIDs = append(returningIDs, elem.ID)
	}
	log.Debug().Msgf("repository.ListPackages: returns packages with IDs: %v", returningIDs)
	/****************/

	return output, nil
}

// RemovePackage implements api.RepoCRUD
func (r *repository) RemovePackage(ctx context.Context, packageID uint64) error {

	log.Debug().Msgf("repository.RemovePackage has called with ID: %d", packageID)

	crudQuery, crudArgs, err := r.initQuery.
		Delete("package").
		Where(sq.Eq{"package_id": packageID}).
		ToSql()
	if err != nil {
		return err
	}

	eventQuery, eventArgs, err := r.initQuery.
		Insert("package_event").
		Columns("package_id", "event_type").
		Values(packageID, "Removed").
		ToSql()
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	log.Debug().Msgf("crud query: %s; args: %v", crudQuery, crudArgs)

	result, err := tx.ExecContext(ctx, crudQuery, crudArgs...)
	if err != nil {
		return err
	}

	// fetch rowsAffected to ensure that a least 1 entity was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		log.Debug().Msgf("package with id %d not found", packageID)
		return api.ErrRepoEntityNotFound
	}

	log.Debug().Msgf("event query: %s; args: %v", eventQuery, eventArgs)

	if _, err := tx.ExecContext(ctx, eventQuery, eventArgs...); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Debug().Msgf("repository.RemovePackage: removed package with ID: %d", packageID)

	return nil
}

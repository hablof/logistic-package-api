package repo

import (
	"context"
	"encoding/json"

	"github.com/hablof/logistic-package-api/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

const (
	DefaultLimit = 10
)

// CreatePackage implements api.RepoCRUD
func (r *repository) CreatePackage(ctx context.Context, pack *model.Package) (uint64, error) {
	returningID := uint64(0)

	crudQuery, crudArgs, err := r.initQuery.Insert("package").
		Columns("title", "material", "max_volume", "reusable", "created_at").
		Values(pack.Title, pack.Material, pack.MaximumVolume, pack.Reusable, "now()").
		Suffix("RETURNING package_id").ToSql()
	if err != nil {
		return 0, err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

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

	eventQuery, eventArgs, err := r.initQuery.Insert("package_event").
		Columns("package_id", "event_type", "payload").
		Values(returningID, "Created", jsonbytes).ToSql()
	if err != nil {
		return 0, err
	}

	if _, err := tx.ExecContext(ctx, eventQuery, eventArgs...); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	log.Debug().Msgf("db: package_id %d successfully inserted", returningID)

	return returningID, nil

}

// DescribePackage implements api.RepoCRUD
func (r *repository) DescribePackage(ctx context.Context, packageID uint64) (*model.Package, error) {
	query, args, err := r.initQuery.Select("package_id", "title", "material", "max_volume", "reusable", "created_at").
		From("package").
		Where(sq.Eq{"package_id": packageID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowxContext(ctx, query, args...)

	unit := model.Package{}
	if err := row.StructScan(&unit); err != nil {
		return nil, err
	}

	return &unit, nil
}

// ListPackages implements api.RepoCRUD
func (r *repository) ListPackages(ctx context.Context, offset uint64) ([]model.Package, error) {
	query, args, err := r.initQuery.Select("package_id", "title", "material", "max_volume", "reusable", "created_at").
		From("package").
		Limit(DefaultLimit).
		Offset(offset).
		ToSql()
	if err != nil {
		return nil, err
	}

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

	return output, nil
}

// RemovePackage implements api.RepoCRUD
func (r *repository) RemovePackage(ctx context.Context, packageID uint64) error {
	crudQuery, crudArgs, err := r.initQuery.Delete("package").
		Where(sq.Eq{"package_id": packageID}).
		ToSql()
	if err != nil {
		return err
	}

	eventQuery, eventArgs, err := r.initQuery.Insert("package_event").
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

	if _, err := tx.ExecContext(ctx, crudQuery, crudArgs...); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, eventQuery, eventArgs...); err != nil {
		return err
	}

	return tx.Commit()
}

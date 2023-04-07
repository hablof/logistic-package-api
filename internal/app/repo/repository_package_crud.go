package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/hablof/logistic-package-api/internal/api"
	"github.com/hablof/logistic-package-api/internal/model"
	"github.com/opentracing/opentracing-go"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

// const (
// 	DefaultLimit = 10
// )

// tables
const (
	packageTable = "package"
	eventTable   = "package_event"
)

// package table column names
const (
	packageIdCol = "package_id"
	titleCol     = "title"
	materialCol  = "material"
	maxVolumeCol = "max_volume"
	reusableCol  = "reusable"
	createdAtCol = "created_at"
	updatedAtCol = "updated_at"
)

// package_event table column names
const (
	// packageIdCol        = "package_id" // same name as in package table
	packageEventIdCol = "package_event_id"
	eventTypeCol      = "event_type"
	eventStatusCol    = "event_status"
	payloadCol        = "payload"
	//createdAtCol        = "created_at"   // same name as in package table
)

type packageModelWithSqlNull struct {
	ID            uint64       `db:"package_id"`
	Title         string       `db:"title"`
	Material      string       `db:"material"`
	MaximumVolume float32      `db:"max_volume"`
	Reusable      bool         `db:"reusable"`
	Created       time.Time    `db:"created_at"`
	Updated       sql.NullTime `db:"updated_at"`
}

// CreatePackage implements api.RepoCRUD
func (r *repository) CreatePackage(ctx context.Context, pack *model.Package, log zerolog.Logger) (uint64, error) {

	repoSpan, ctx := opentracing.StartSpanFromContext(ctx, "repository.CreatePackage")
	defer repoSpan.Finish()

	log.Debug().Msgf("repository.CreatePackage has called with package title: %s", pack.Title)

	crudQuery, crudArgs, err := r.initQuery.
		Insert(packageTable).
		Columns(titleCol, materialCol, maxVolumeCol, reusableCol, createdAtCol).
		Values(pack.Title, pack.Material, pack.MaximumVolume, pack.Reusable, "now()").
		Suffix("RETURNING package_id, created_at").
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
	timestamp := time.Time{}

	crudSpan, _ := opentracing.StartSpanFromContext(ctx, "crud query")
	row := tx.QueryRowxContext(ctx, crudQuery, crudArgs...)
	crudSpan.Finish()

	if err := row.Scan(&returningID, &timestamp); err != nil {
		return 0, err
	}

	unit := model.Package{
		ID:            returningID,
		Title:         pack.Title,
		Material:      pack.Material,
		MaximumVolume: pack.MaximumVolume,
		Reusable:      pack.Reusable,
		Created:       timestamp,
		Updated:       nil,
	}

	jsonbytes, err := json.Marshal(unit)
	if err != nil {
		return 0, err
	}

	eventQuery, eventArgs, err := r.initQuery.
		Insert(eventTable).
		Columns(packageIdCol, eventTypeCol, payloadCol).
		Values(returningID, "Created", jsonbytes).
		ToSql()
	if err != nil {
		return 0, err
	}

	log.Debug().Msgf("event query: %s; args: %v", eventQuery, eventArgs)

	eventSpan, _ := opentracing.StartSpanFromContext(ctx, "event query")
	_, err = tx.ExecContext(ctx, eventQuery, eventArgs...)
	eventSpan.Finish()

	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	log.Debug().Msgf("repository.CreatePackage: package_id %d inserted", returningID)

	return returningID, nil

}

// DescribePackage implements api.RepoCRUD
func (r *repository) DescribePackage(ctx context.Context, packageID uint64, log zerolog.Logger) (*model.Package, error) {

	repoSpan, ctx := opentracing.StartSpanFromContext(ctx, "repository.DescribePackage")
	defer repoSpan.Finish()

	log.Debug().Msgf("repository.DescribePackage has called with ID: %d", packageID)

	query, args, err := r.initQuery.
		Select(packageIdCol, titleCol, materialCol, maxVolumeCol, reusableCol, createdAtCol, updatedAtCol).
		From(packageTable).
		Where(sq.Eq{packageIdCol: packageID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("query: %s; args: %v", query, args)

	crudSpan, _ := opentracing.StartSpanFromContext(ctx, "crud query")
	row := r.db.QueryRowxContext(ctx, query, args...)
	crudSpan.Finish()

	scanUnit := packageModelWithSqlNull{}
	switch err := row.StructScan(&scanUnit); {
	case errors.Is(err, sql.ErrNoRows):
		return nil, api.ErrRepoEntityNotFound

	case err != nil:
		return nil, err
	}

	unit := model.Package{
		ID:            scanUnit.ID,
		Title:         scanUnit.Title,
		Material:      scanUnit.Material,
		MaximumVolume: scanUnit.MaximumVolume,
		Reusable:      scanUnit.Reusable,
		Created:       scanUnit.Created,
	}

	if scanUnit.Updated.Valid {
		unit.Updated = &scanUnit.Updated.Time
	} else {
		unit.Updated = nil
	}

	log.Debug().Msgf("repository.DescribePackage: package_id %d has read", packageID)

	return &unit, nil
}

// ListPackages implements api.RepoCRUD
func (r *repository) ListPackages(ctx context.Context, offset uint64, limit uint64, log zerolog.Logger) ([]model.Package, error) {

	repoSpan, ctx := opentracing.StartSpanFromContext(ctx, "repository.ListPackages")
	defer repoSpan.Finish()

	repoSpan.SetTag("query offset", offset)

	log.Debug().Msgf("repository.ListPackages has called with offset %d", offset)

	query, args, err := r.initQuery.
		Select(packageIdCol, titleCol, materialCol, maxVolumeCol, reusableCol, createdAtCol, updatedAtCol).
		From(packageTable).
		OrderBy(packageIdCol).
		Limit(limit).
		Offset(offset).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("query: %s; args: %v", query, args)

	crudSpan, _ := opentracing.StartSpanFromContext(ctx, "crud query")
	rows, err := r.db.QueryxContext(ctx, query, args...)
	crudSpan.Finish()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	output := make([]model.Package, 0, limit)

	for rows.Next() {
		unit := model.Package{}
		if err := rows.StructScan(&unit); err != nil {
			return nil, err
		}
		output = append(output, unit)
	}

	if log.GetLevel() == zerolog.DebugLevel {
		returningIDs := make([]uint64, 0, len(output))
		for _, elem := range output {
			returningIDs = append(returningIDs, elem.ID)
		}
		log.Debug().Msgf("repository.ListPackages: returns packages with IDs: %v", returningIDs)
	}

	return output, nil
}

// RemovePackage implements api.RepoCRUD
func (r *repository) RemovePackage(ctx context.Context, packageID uint64, log zerolog.Logger) error {

	repoSpan, ctx := opentracing.StartSpanFromContext(ctx, "repository.RemovePackage")
	defer repoSpan.Finish()

	log.Debug().Msgf("repository.RemovePackage has called with ID: %d", packageID)

	crudQuery, crudArgs, err := r.initQuery.
		Delete(packageTable).
		Where(sq.Eq{packageIdCol: packageID}).
		ToSql()
	if err != nil {
		return err
	}

	eventQuery, eventArgs, err := r.initQuery.
		Insert(eventTable).
		Columns(packageIdCol, eventTypeCol).
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

	crudSpan, _ := opentracing.StartSpanFromContext(ctx, "crud query")
	result, err := tx.ExecContext(ctx, crudQuery, crudArgs...)
	crudSpan.Finish()

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

	eventSpan, _ := opentracing.StartSpanFromContext(ctx, "event query")
	_, err = tx.ExecContext(ctx, eventQuery, eventArgs...)
	eventSpan.Finish()

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Debug().Msgf("repository.RemovePackage: removed package with ID: %d", packageID)

	return nil
}

// UpdatePackage implements api.RepoCRUD
func (r *repository) UpdatePackage(ctx context.Context, packageID uint64, changes map[api.FieldName]interface{}, log zerolog.Logger) error {

	repoSpan, ctx := opentracing.StartSpanFromContext(ctx, "repository.UpdatePackage")
	defer repoSpan.Finish()

	ub := r.initQuery.Update(packageTable).Set(updatedAtCol, "now()")

	for key, value := range changes {
		switch key {
		case api.Title:
			ub = ub.Set(titleCol, value)

		case api.Material:
			ub = ub.Set(materialCol, value)

		case api.MaxVolume:
			ub = ub.Set(maxVolumeCol, value)

		case api.Reusable:
			ub = ub.Set(reusableCol, value)
		}
	}

	crudQuery, crudArgs, err := ub.Where(sq.Eq{packageIdCol: packageID}).
		Suffix("RETURNING package_id, title, material, max_volume, reusable, created_at, updated_at").
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

	scanUnit := packageModelWithSqlNull{}
	// unit := model.Package{}

	crudSpan, _ := opentracing.StartSpanFromContext(ctx, "crud query")
	row := tx.QueryRowxContext(ctx, crudQuery, crudArgs...)
	crudSpan.Finish()
	if err := row.StructScan(&scanUnit); err != nil {
		return err
	}

	payload, err := json.Marshal(scanUnit)
	if err != nil {
		return err
	}

	eventQuery, eventArgs, err := r.initQuery.
		Insert(eventTable).
		Columns(packageIdCol, eventTypeCol, payloadCol).
		Values(packageID, "Updated", payload).
		ToSql()
	if err != nil {
		return err
	}

	log.Debug().Msgf("event query: %s; args: %v", eventQuery, eventArgs)

	eventSpan, _ := opentracing.StartSpanFromContext(ctx, "event query")
	_, err = tx.ExecContext(ctx, eventQuery, eventArgs...)
	eventSpan.Finish()

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	log.Debug().Msgf("repository.UpdatePackage: package_id %d updated", packageID)

	return nil
}

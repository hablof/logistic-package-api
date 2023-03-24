package repo

import (
	"context"
	"time"

	"github.com/hablof/logistic-package-api/internal/model"
	"github.com/rs/zerolog/log"

	sq "github.com/Masterminds/squirrel"
)

type scanPackageStruct struct {
	ID        uint64    `db:"package_event_id"`
	PackageID uint64    `db:"package_id"`
	Type      string    `db:"event_type"`
	Status    string    `db:"event_status"`
	Created   time.Time `db:"created_at"`
	Payload   []byte    `db:"payload"`
}

const (
	defaultTimeout = 5 * time.Second
)

// Lock implements consumer.RepoEventConsumer
func (r *repository) Lock(limit uint64) ([]model.PackageEvent, error) { // use r.batchsize instead argument limit ?

	log.Debug().Msgf("repository.Lock was called to lock %d entries", limit)

	query, args, err := r.initQuery.Update("package_event").
		Set("event_status", "Locked").
		Where("package_event_id IN (SELECT package_event_id FROM package_event WHERE event_status = ? LIMIT ?)", "Unlocked", limit).
		Suffix("RETURNING package_event_id, package_id, event_type, event_status, payload, created_at").
		ToSql()
	if err != nil {
		return nil, err
	}

	ctx, cf := context.WithTimeout(context.Background(), defaultTimeout)
	defer cf()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	scanUnit := scanPackageStruct{}

	output := make([]model.PackageEvent, 0, limit)
	for rows.Next() {
		if err := rows.StructScan(&scanUnit); err != nil {
			return nil, err
		}
		unit := r.decodeScanStruct(scanUnit)

		output = append(output, unit)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// ****************** //
	returningIDs := make([]uint64, 0, len(output))
	for _, elem := range output {
		returningIDs = append(returningIDs, elem.ID)
	}
	log.Debug().Msgf("repository.Lock locked events with IDs: %v", returningIDs)
	// ****************** //

	return output, nil
}

// Postgres ENUM to golang types
func (*repository) decodeScanStruct(scanUnit scanPackageStruct) model.PackageEvent {
	unit := model.PackageEvent{
		ID:        scanUnit.ID,
		PackageID: scanUnit.PackageID,
		Created:   scanUnit.Created,
		Payload:   scanUnit.Payload,
	}

	switch scanUnit.Status {
	case "Locked":
		unit.Status = model.Locked
	case "Unlocked":
		unit.Status = model.Unlocked
	}

	switch scanUnit.Type {
	case "Created":
		unit.Type = model.Created
	case "Updated":
		unit.Type = model.Updated
	case "Removed":
		unit.Type = model.Removed
	}

	return unit
}

// Remove implements cleaner.RepoEventCleaner
func (r *repository) Remove(eventIDs []uint64) error {

	log.Debug().Msgf("repository.Remove was called with arg: %v", eventIDs)

	query, args, err := r.initQuery.Delete("package_event").
		Where(sq.Eq{"package_event_id": eventIDs}).ToSql()
	if err != nil {
		return err
	}

	ctx, cf := context.WithTimeout(context.Background(), defaultTimeout)
	defer cf()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

// Unlock implements cleaner.RepoEventCleaner
func (r *repository) Unlock(eventIDs []uint64) error {

	log.Debug().Msgf("repository.Unlock was called with arg: %v", eventIDs)

	query, args, err := r.initQuery.Update("package_event").
		Set("event_status", "Unlocked").
		Where(sq.Eq{"package_event_id": eventIDs}).ToSql()
	if err != nil {
		return err
	}

	ctx, cf := context.WithTimeout(context.Background(), defaultTimeout)
	defer cf()

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

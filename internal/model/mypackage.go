package model

import "time"

type EventType uint8

type EventStatus uint8

type TimesDefered uint8

const (
	_ EventType = iota
	Created
	Updated
	Removed
)

const (
	_ EventStatus = iota
	Deferred
	Processed
)

type Package struct {
	ID            uint64  `db:"package_id"`
	Title         string  `db:"title"`
	Material      string  `db:"material"`
	MaximumVolume float32 `db:"max_volume"`
	Reusable      bool    `db:"reusable"`
}

type PackageEvent struct {
	ID        uint64       `db:"package_event_id"`
	PackageID uint64       `db:"package_id"`
	Type      EventType    `db:"event_type"`
	Status    EventStatus  `db:"event_status"`
	Created   time.Time    `db:"created_at"`
	Payload   []byte       `db:"payload"`
	Defered   TimesDefered // retranslator param
}

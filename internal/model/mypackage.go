package model

import "time"

type EventType uint8

type EventStatus uint8

// type TimesDefered uint8

const (
	_ EventType = iota
	Created
	Updated
	Removed
)

const (
	_ EventStatus = iota
	Locked
	Unlocked
)

type Package struct {
	ID            uint64     `db:"package_id"`
	Title         string     `db:"title"`
	Material      string     `db:"material"`
	MaximumVolume float32    `db:"max_volume"`
	Reusable      bool       `db:"reusable"`
	Created       time.Time  `db:"created_at"`
	Updated       *time.Time `db:"updated_at"`
}

type PackageEvent struct {
	ID        uint64      // `db:"package_event_id"`
	PackageID uint64      // `db:"package_id"`
	Type      EventType   // `db:"event_type"`
	Status    EventStatus // `db:"event_status"` // Это даже нигде не используется, кроме как внутри базы
	Created   time.Time   // `db:"created_at"`
	Payload   []byte      // `db:"payload"`
	// Defered   TimesDefered // retranslator param for keeping correct order sending to kafka
}

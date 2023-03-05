package model

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed
)

const (
	Deferred EventStatus = iota
	Processed
)

type Package struct {
	Title         string
	Material      string
	MaximumVolume float32 //cm^3
	Reusable      bool
}

type PackageEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Package
}

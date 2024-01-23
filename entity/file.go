package entity

import "time"

type File struct {
	ID        uint64
	GUID      string
	Path      string
	CreatedAt time.Time
}

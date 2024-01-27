package entity

import "time"

type Clinic struct {
	ID          uint64
	GUID        string
	Location_ID uint64
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

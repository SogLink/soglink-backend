package entity

import "time"

type Clinic struct {
	ID          uint64
	GUID        string
	Location_ID uint64
	Name        string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

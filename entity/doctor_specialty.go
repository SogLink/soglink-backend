package entity

import "time"

type Doctor_specialty struct {
	DoctorID    uint64
	SpecialtyID uint64
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

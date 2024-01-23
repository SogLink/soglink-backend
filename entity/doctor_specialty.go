package entity

import "time"

type Doctor_specialty struct {
	Specialty string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

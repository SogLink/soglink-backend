package entity

import "time"

type User struct {
	ID        uint64
	GUID      string
	Username  string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

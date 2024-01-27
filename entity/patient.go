package entity

import "time"

type Patient struct {
	User       *User
	Patient_ID uint64
	Name       string
	Surname    string
	Gender     string
	Birthday   time.Time
	Pinfl      uint64
}

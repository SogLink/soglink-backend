package entity

import "time"

type Doctor struct {
	Doctor_ID    uint64
	Clinic_ID    uint64
	Name         string
	Surname      string
	Birthday     time.Time
	Gender       string
	Education    string
	Certificates string
}

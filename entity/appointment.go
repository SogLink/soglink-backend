package entity

import "time"

type Appointment struct {
	Doctor_ID         uint64
	Patient_ID        uint64
	ID                uint64
	GUID              string
	AppointmentAt     time.Time
	AppointmentReason string
	Price             float64
	Status            string
	EmrID             uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

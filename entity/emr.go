package entity

import "time"

type Emr struct {
	Diagnoses_text     string
	Prescriptions_text string
	Doctor_ID          uint64
	Patient_ID         uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

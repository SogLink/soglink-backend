package usecase

import (
	"context"
	"time"

	"github.com/SogLink/soglink-backend/entity"
	"github.com/SogLink/soglink-backend/repository/patient"
)

type Patient interface {
	Create(ctx context.Context, patient *entity.Patient) error
}

type patientUsecase struct {
	BaseUsecase
	patientRepo patient.Repository
	ctxTimeout  time.Duration
}

func NewPatientUsecase(ctxTimeout time.Duration, patientRepo patient.Repository) Patient {
	return &patientUsecase{
		patientRepo: patientRepo,
		ctxTimeout:  ctxTimeout,
	}
}

func (u patientUsecase) Create(ctx context.Context, patient *entity.Patient) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.patientRepo.Create(ctx, patient)
}

package v1

import (
	"net/http"

	"github.com/SogLink/soglink-backend/api/handler"
	"github.com/SogLink/soglink-backend/pkg/config"
	clinic "github.com/SogLink/soglink-backend/repository/Clinic"
	"github.com/SogLink/soglink-backend/repository/doctor"
	doctorspeciality "github.com/SogLink/soglink-backend/repository/doctor_speciality"
	"github.com/SogLink/soglink-backend/repository/location"
	"github.com/SogLink/soglink-backend/repository/specialty"
	usecase "github.com/SogLink/soglink-backend/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type doctorHandler struct {
	reshreshTokenUsecase usecase.RefreshToken
	userUsecase          usecase.User
	doctorRepo           doctor.Repository
	clinicRepo           clinic.Repository
	locationRepo         location.Repository
	docotrSpecRepo       doctorspeciality.Repository
	specialtyRepo        specialty.Repository
	logger               *zap.Logger
	config               *config.Config
}

func NewDoctorHandler(option handler.HandlerArguments) http.Handler {
	handler := doctorHandler{
		reshreshTokenUsecase: option.ReshreshTokenUsecase,
		logger:               option.Logger,
		config:               option.Config,
		userUsecase:          option.UserUsecase,
	}

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		// public apis
		r.Get("/", handler.GetDoctorList())

	})
	return router
}

// GetDoctorList
// @Security ApiKeyAuth
// @Router /v1/doctor [GET]
// @Summary Get one dentist
// @Description Get one dentist by ID
// @Tags dentists
// @Accept json
// @Produce json
// @Param limit path int false "limit"
// @Param page path int false "page"
// @Param specialty path string true "specialty"
// @Param doctor_name path string true "doctor_name"
// @Param clinic_id path string true "clinic_id"
// @Success 200 {object} models.GetDoctorListResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h doctorHandler) GetDoctorList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// GetSpecialty
// @Security ApiKeyAuth
// @Router /v1/doctor/specialty [GET]
// @Summary Get doctor specialty
// @Description Get  doctor specialty
// @Tags dentists
// @Accept json
// @Produce json
// @Success 200 {object} models.GetSpecialty
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h doctorHandler) GetDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

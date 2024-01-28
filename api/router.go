package api

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"

	_ "github.com/SogLink/soglink-backend/api/docs"
	"github.com/SogLink/soglink-backend/api/handler"
	v1 "github.com/SogLink/soglink-backend/api/handler/v1"
	"github.com/SogLink/soglink-backend/api/middleware"
	"github.com/SogLink/soglink-backend/pkg/config"
	usecase "github.com/SogLink/soglink-backend/services"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RouteArguments struct {
	ReshreshTokenUsecase usecase.RefreshToken
	UserUsecase          usecase.User
	PatientUsecase       usecase.Patient
	Logger               *zap.Logger
	Config               *config.Config
}

// NewRoute
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(args RouteArguments) http.Handler {
	handlersArgs := handler.HandlerArguments{
		Config:               args.Config,
		ReshreshTokenUsecase: args.ReshreshTokenUsecase,
		UserUsecase:          args.UserUsecase,
		Logger:               args.Logger,
		PatientUsecase:       args.PatientUsecase,
	}

	router := chi.NewRouter()
	router.Use(chimiddleware.RealIP, chimiddleware.Logger, chimiddleware.Recoverer)
	// router.Use(chimiddleware.Timeout(args.ContextTimeout))
	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/v1", func(r chi.Router) {
		// protected group api
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthContext(args.Config.Token.Secret))
		})

		// public group api
		r.Group(func(r chi.Router) {
			r.Mount("/auth", v1.NewAuthHandler(handlersArgs))
			r.Mount("/doctor", v1.NewDoctorHandler(handlersArgs))
		})
	})

	// declare swagger api route
	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

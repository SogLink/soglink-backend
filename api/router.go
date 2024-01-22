package api

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	// _ "github.com/SogLink/soglink-backend/api/docs"
	// "github.com/SogLink/soglink-backend/api/handlers"
	// v1 "github.com/SogLink/soglink-backend/api/handlers/v1"
	// "github.com/AsaHero/abclinic/api/middleware"
	"github.com/SogLink/soglink-backend/pkg/config"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RouteArguments struct {
	Config *config.Config
}

// NewRoute
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(args RouteArguments) http.Handler {
	// handlersArgs := handlers.HandlerArguments{
	// 	Config: args.Config,
	// }

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
	})

	// declare swagger api route
	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/SogLink/soglink-backend/api"
	"github.com/SogLink/soglink-backend/pkg/config"
	"github.com/SogLink/soglink-backend/pkg/database"
	"github.com/SogLink/soglink-backend/pkg/logger"
	"github.com/SogLink/soglink-backend/repository/patient"
	refreshtoken "github.com/SogLink/soglink-backend/repository/refresh_token"
	"github.com/SogLink/soglink-backend/repository/user"
	usecase "github.com/SogLink/soglink-backend/services"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
	DB     *postgres.PostgresDB
	server *http.Server
}

func NewApp(cfg *config.Config) *App {
	// logger init
	logger, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		log.Fatalf("error on logger init: %v", err)
	}

	// db init
	db, err := postgres.NewPostgresDB(*cfg)
	if err != nil {
		log.Fatalf("error on db init: %v", err)
	}

	return &App{
		Logger: logger,
		Config: cfg,
		DB:     db,
	}
}

func (a *App) Run() error {
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error while parsing context timeout: %v", err)
	}

	// repo init
	patientRepo := patient.NewPatientRepo(a.DB)
	userRepo := user.NewUserRepo(a.DB)
	refreshTokenRepo := refreshtoken.NewRefreshTokenRepo(a.DB)

	// usecase init
	userUsecase := usecase.NewUserUsecase(contextTimeout, userRepo)
	patientUsecase := usecase.NewPatientUsecase(contextTimeout, patientRepo)
	refreshTokenUsecase := usecase.NewRefreshTokenService(contextTimeout, refreshTokenRepo)

	routerArgs := api.RouteArguments{
		Config:               a.Config,
		ReshreshTokenUsecase: refreshTokenUsecase,
		UserUsecase:          userUsecase,
		Logger:               a.Logger,
		PatientUsecase:       patientUsecase,
	}

	// router init
	handlers := api.NewRouter(routerArgs)

	// server init
	a.server, err = api.NewServer(a.Config, handlers)
	if err != nil {
		return fmt.Errorf("error while server init: %v", err)
	}

	return a.server.ListenAndServe()
}

func (a *App) Stop() {

	// close db pool
	a.DB.Close()

	// shutdown server http
	if err := a.server.Shutdown(context.Background()); err != nil {
		a.Logger.Error("shutdown server http", zap.Error(err))
	}

	// logger sync
	a.Logger.Sync()
}

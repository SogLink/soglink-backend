package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SogLink/soglink-backend/app"
	"github.com/SogLink/soglink-backend/pkg/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// load dot env file
	godotenv.Load()
	// config init
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("cannot load env: %v", err)
	}

	// app init
	app := app.NewApp(cfg)

	// app runs
	go func() {
		app.Logger.Info("Listen:", zap.String("address", cfg.Server.Host+cfg.Server.Port))
		if err := app.Run(); err != nil {
			app.Logger.Error("error while running server", zap.Error(err))
		}
	}()

	// wait for sigint
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// app stops
	app.Logger.Info("soglink sevrver stops")
	app.Stop()
}

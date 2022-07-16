package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/berikarg/fortune-wheel/api"
	"github.com/berikarg/fortune-wheel/internal/config"
	"github.com/berikarg/fortune-wheel/internal/server"
	"github.com/berikarg/fortune-wheel/pkg/handler"
	"github.com/berikarg/fortune-wheel/pkg/repository"
	"github.com/berikarg/fortune-wheel/pkg/service"
)

var configPath = flag.String("c", "./configs/config.yml", "config file path")

func main() {
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer logger.Sync() //nolint

	cfg, err := config.New(*configPath)
	if err != nil {
		logger.Fatal("parse config", zap.Error(err))
	}

	db, err := repository.OpenPostgres(cfg.Database)
	if err != nil {
		logger.Fatal("open db connection", zap.Error(err))
	}
	defer db.Close()

	initWheel := api.Wheel{Fields: []string{"1", "2", "3", "4"}}

	repo := repository.NewSpinResultRepository(db)
	services := service.NewSpinResultService(repo)
	handlers := handler.NewHandler(services, logger, initWheel)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
			logger.Fatal("run http server", zap.Error(err))
		}
	}()

	logger.Info("Fortune wheel server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Info("Fortune wheel server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error("server shut down", zap.Error(err))
	}
}

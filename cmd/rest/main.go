package main

import (
	"context"
	"github.com/elsyarif/go-auth-api/internal/infrastructure/database"
	"github.com/elsyarif/go-auth-api/internal/infrastructure/http"
	"github.com/elsyarif/go-auth-api/pkg/config"
	"github.com/elsyarif/go-auth-api/pkg/helper/log"
	"github.com/sirupsen/logrus"
)

func main() {
	err := database.LoadConfig()
	if err != nil {
		log.Fatal("load configuration failed", logrus.Fields{"error": err.Error()})
	}

	log.Info("Config", logrus.Fields{
		"DBName": config.Conf.DBName,
	})
	postgres, err := database.NewConnectPostgres()
	if err != nil {
		log.Fatal("Database configuration failed", logrus.Fields{"error": err.Error()})
	}
	c := context.Background()
	ctx, svr := http.NewServer(c, config.Conf.Host, config.Conf.Port, config.Conf.ShutdownTimeout, postgres)
	err = svr.Run(ctx)
	if err != nil {
		log.Fatal("server run failed", logrus.Fields{"error": err.Error()})
	}
}

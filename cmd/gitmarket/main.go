package main

import (
	"context"
	app "github.com/Warh40k/gitmarket/pkg/app"
	"github.com/Warh40k/gitmarket/pkg/handler"
	"github.com/Warh40k/gitmarket/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// TODO: перенести pkg в internal (в pkg оставить сторонние либы, для работы с REST, логгер, например)
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars %s", err.Error())
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)
	srv := new(app.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured during shutting down %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

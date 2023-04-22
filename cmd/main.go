package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"Forum"
	"Forum/logger"
	"Forum/pkg/handlers"
	"Forum/pkg/repository"
	"Forum/pkg/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	l := logger.Logger(ctx)
	db, err := repository.NewSQLiteDB()
	config := repository.ReadConfig()
	if err != nil {
		log.Fatal("error with read config")
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewHandler(services)

	l.Info("App starting")
	log.Printf("Starting server...\nhttps://localhost%v/\n", ":"+config.Port)

	srv := new(Forum.Server)
	if err := srv.Run(config.Port, handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}

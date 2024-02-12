package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/kaitsubaka/hexagonal-architecture-go/docs"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/app/rest"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/config"
	todosvc "github.com/kaitsubaka/hexagonal-architecture-go/internal/core/services/todo"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/adapters/driven/repository"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/adapters/driving/rest/controllers"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/adapters/driving/rest/routes"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/platform/database"
)

// @title Hexagonal arquitecture rest API

// @version 1.0

// @description This is a todo sample server.

// @contact.name API Support

// @contact.email kaitsubaka@gmail.com

// @license.name MIT

// @BasePath /api
func main() {
	cfg, err := config.Load(config.WithEnv(
		os.Getenv("APP_ENV"),
	))
	if err != nil {
		log.Fatal(err)
	}

	baseContext, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	psqlDBClient, err := database.NewPsqlDB(&database.PsqlDBConfig{
		Host:     cfg.Postgres.Host,
		Password: cfg.Postgres.Password,
		User:     cfg.Postgres.User,
		DBName:   cfg.Postgres.DBName,
		Port:     cfg.Postgres.Port,
		Ctx:      baseContext,
	})
	if err != nil {
		log.Fatal(err)
	}

	restApp, err := rest.
		NewAppBuilder(cfg).
		WithContext(baseContext).
		WithGroup("/todos",
			routes.TodoRoutes(
				controllers.NewTodoController(
					todosvc.NewTodoService(
						repository.NewTodoRepository(psqlDBClient, "todos"),
					),
				),
			)).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := restApp.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	log.Println("Press CTRL+C to exit...")

	<-baseContext.Done()

	// graceful shutdown
	if err := restApp.Shutdown(); err != nil {
		log.Fatal(err)
	}

	if err := psqlDBClient.Close(); err != nil {
		log.Fatal(err)
	}
}

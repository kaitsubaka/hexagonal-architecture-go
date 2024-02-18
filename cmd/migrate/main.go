package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/kaitsubaka/hexagonal-architecture-go/internal/common"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/config"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/platform/database"
	migratecli "github.com/kaitsubaka/hexagonal-architecture-go/internal/platform/database/migrate"
	"github.com/kaitsubaka/hexagonal-architecture-go/scripts/database/migrations"
	"github.com/uptrace/bun/migrate"
)

func main() {
	cfg, err := config.Load(config.WithEnv(os.Getenv(common.AppEnvEnvar)))
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	psqlDB, err := database.NewPsqlDB(&database.PsqlDBConfig{
		Host:     cfg.Postgres.Host,
		Password: cfg.Postgres.Password,
		User:     cfg.Postgres.User,
		DBName:   cfg.Postgres.DBName,
		Port:     cfg.Postgres.Port,
		Ctx:      ctx,
	})
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := migratecli.NewRootCmd()

	migratecli.NewMigrateCommand(
		rootCmd,
		migrate.NewMigrator(psqlDB, migrations.Migrations),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

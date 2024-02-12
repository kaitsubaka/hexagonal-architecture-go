package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/pubdtos"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

type PsqlDB struct {
	dbConn     bun.IDB
	shutDownFn func() error
}

type PsqlDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Ctx      context.Context
}

func NewPsqlDB(cfg *PsqlDBConfig) (*bun.DB, error) {

	sqldb, err := sql.Open("pg", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	if err := db.PingContext(cfg.Ctx); err != nil {
		return nil, err
	}
	if err := db.ResetModel(cfg.Ctx, (*pubdtos.TodoPsqlDTO)(nil)); err != nil {
		return nil, err
	}

	return db, nil
}

func (pdbc *PsqlDB) Connection() bun.IDB {
	return pdbc.dbConn
}

func (pdbc *PsqlDB) Shutdown() error {
	return pdbc.shutDownFn()
}

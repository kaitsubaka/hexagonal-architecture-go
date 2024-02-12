package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/dtos"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/pubdtos"
	"github.com/uptrace/bun"
)

type TodoReadRepository struct {
	dbClient  bun.IDB
	tableName string
}

func NewTodoRepository(dbClient bun.IDB, tableName string) *TodoReadRepository {
	return &TodoReadRepository{
		dbClient:  dbClient,
		tableName: tableName,
	}
}

func (tr *TodoReadRepository) table() string {
	return fmt.Sprintf("%s AS t", tr.tableName)
}

func (tr *TodoReadRepository) FindById(
	ctx context.Context,
	id uuid.UUID,
) (*dtos.TodoDTO, error) {
	queryModel := &pubdtos.TodoPsqlDTO{
		ID: id,
	}
	if err := tr.dbClient.NewSelect().
		Model(queryModel).
		ModelTableExpr(tr.table()).
		WherePK().
		Scan(ctx, queryModel); err != nil {
		return nil, err
	}
	return &dtos.TodoDTO{
		ID:   queryModel.ID,
		Note: queryModel.Note,
	}, nil
}

func (tr *TodoReadRepository) Create(
	ctx context.Context,
	todo *dtos.TodoDTO,
) (*dtos.TodoDTO, error) {
	psqlTodoModel := &pubdtos.TodoPsqlDTO{
		ID:   todo.ID,
		Note: todo.Note,
	}
	if _, err := tr.dbClient.
		NewInsert().
		Model(psqlTodoModel).
		ModelTableExpr(tr.table()).
		Exec(ctx); err != nil {
		return nil, err
	}
	return &dtos.TodoDTO{
		ID:   psqlTodoModel.ID,
		Note: psqlTodoModel.Note,
	}, nil
}

package driving

import (
	"context"

	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/dtos"
)

type TodoService interface {
	GetTodoById(context.Context, string) (*dtos.TodoDTO, error)
	CreateTodo(ctx context.Context, dto *dtos.CreateTodoDTO) (*dtos.TodoDTO, error)
}

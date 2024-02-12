package driven

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/dtos"
)

type TodoRepository interface {
	Create(
		ctx context.Context,
		todo *dtos.TodoDTO,
	) (*dtos.TodoDTO, error)
	FindById(
		ctx context.Context,
		id uuid.UUID,
	) (*dtos.TodoDTO, error)
}

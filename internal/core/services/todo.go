package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/dtos"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/entities"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/ports/driven"
)

type TodoService struct {
	todoRepository driven.TodoRepository
}

func NewTodoService(todoRepository driven.TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (ts *TodoService) GetTodoById(ctx context.Context, id string) (*dtos.TodoDTO, error) {
	uID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return ts.todoRepository.FindById(ctx, uID)
}

func (ts *TodoService) CreateTodo(ctx context.Context, dto *dtos.CreateTodoDTO) (*dtos.TodoDTO, error) {
	todo, err := entities.NewTodoFromCreateTodoDTO(dto)
	if err != nil {
		return nil, err
	}
	return ts.todoRepository.Create(ctx, &dtos.TodoDTO{
		ID:   todo.ID,
		Note: todo.Note.String(),
	})
}

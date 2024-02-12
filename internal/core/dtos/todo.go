package dtos

import "github.com/google/uuid"

// TodoDTO represents the model for an order
type TodoDTO struct {
	ID   uuid.UUID
	Note string
}

type CreateTodoDTO struct {
	Note string `example:"name_example"`
}

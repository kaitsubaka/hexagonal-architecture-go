package pubdtos

import "github.com/google/uuid"

// GetTodoByIDResponseDTO represents the model for an order
type GetTodoByIDResponseDTO struct {
	ID   uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000"`
	Note string    `example:"name_example"`
}

// CreateTodoResponseDTO represents the model for an order
type CreateTodoResponseDTO struct {
	ID   uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000"`
	Note string    `example:"name_example"`
}

// CreateTodoRequestDTO represents the model for an order
type CreateTodoRequestDTO struct {
	Note string `example:"name_example"`
}

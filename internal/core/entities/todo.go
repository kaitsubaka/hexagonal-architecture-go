package entities

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/dtos"
)

type Note string

func NewNoteFromString(in string) (*Note, error) {
	if in == "" {
		return nil, errors.New("NewNoteFromString: empty note is not allowed")
	}
	note := Note(in)
	return &note, nil
}

func (n *Note) String() string {
	if n == nil {
		return "<nil>"
	}
	return string(*n)
}

type Todo struct {
	ID   uuid.UUID
	Note *Note
}

func NewTodoFromCreateTodoDTO(in *dtos.CreateTodoDTO) (*Todo, error) {
	if in == nil {
		return nil, errors.New("NewTodoFromCreateTodoDTO: creation input error")
	}
	note, err := NewNoteFromString(in.Note)
	if err != nil {
		return nil, fmt.Errorf("NewTodoFromCreateTodoDTO: error parsing note caused by (%w)", err)
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Todo{
		ID:   id,
		Note: note,
	}, nil
}

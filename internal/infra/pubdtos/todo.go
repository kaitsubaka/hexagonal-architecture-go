package pubdtos

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TodoPsqlDTO struct {
	bun.BaseModel `bun:"table:todos,alias:t"`
	ID            uuid.UUID `bun:"type:uuid,pk"`
	Note          string
}

package controllers

import (
	"context"
	"net/http"

	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/dtos"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/core/ports/driving"
	"github.com/kaitsubaka/hexagonal-architecture-go/pkg/pubdtos"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	todoService driving.TodoService
}

func NewTodoController(todoService driving.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

// GetTodoById godoc
// @Summary Get Todo detail by ID
// @Description get string by ID
// @ID get-todo-detail-by-todoid
// @Accept  json
// @Produce  json
// @Param id path string true "Todo ID"
// @Success 200 {object} pubdtos.GetTodoByIDResponseDTO
// @Router /todos/{id} [get]
func (th *TodoController) GetTodoById() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto, err := th.todoService.GetTodoById(c.Request().Context(), c.Param("id"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, pubdtos.GetTodoByIDResponseDTO{
			ID:   dto.ID,
			Note: dto.Note,
		})
	}
}

// CreateTodo godoc
// @Summary Get Todo detail by ID
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param todo body pubdtos.CreateTodoRequestDTO true "Create order"
// @Success 200 {object} pubdtos.CreateTodoResponseDTO
// @Router /todos [post]
func (th *TodoController) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		createTodoRequestDTO := new(pubdtos.CreateTodoRequestDTO)

		if err := c.Bind(createTodoRequestDTO); err != nil {
			return err
		}

		dto, err := th.todoService.CreateTodo(context.Background(), &dtos.CreateTodoDTO{
			Note: createTodoRequestDTO.Note,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, dto)
	}
}

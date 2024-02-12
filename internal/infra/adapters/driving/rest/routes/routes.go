package routes

import (
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/adapters/driving/rest/controllers"

	"github.com/labstack/echo/v4"
)

type RegistrationFunc func(router *echo.Group)

func TodoRoutes(c *controllers.TodoController) RegistrationFunc {
	return func(router *echo.Group) {
		router.GET("/:id", c.GetTodoById())
		router.POST("", c.CreateTodo())
	}
}

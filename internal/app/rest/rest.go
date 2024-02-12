package rest

import (
	"context"
	"fmt"
	"time"

	"github.com/kaitsubaka/hexagonal-architecture-go/internal/config"
	"github.com/kaitsubaka/hexagonal-architecture-go/internal/infra/adapters/driving/rest/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	cfg     *config.Configuration
	ctx     context.Context
	address string
	server  *echo.Echo
}

func NewAppBuilder(cfg *config.Configuration) *AppBuilder {
	return &AppBuilder{
		cfg:    cfg,
		routes: make(map[string]routes.RegistrationFunc),
	}
}

type AppBuilder struct {
	cfg    *config.Configuration
	ctx    context.Context
	routes map[string]routes.RegistrationFunc
}

func (a *AppBuilder) WithGroup(path string, regFn routes.RegistrationFunc) *AppBuilder {
	a.routes[path] = regFn
	return a
}
func (a *AppBuilder) WithContext(ctx context.Context) *AppBuilder {
	a.ctx = ctx
	return a
}

func (a *AppBuilder) Build() (*App, error) {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	baseGroup := e.Group("/api")
	for path, regFn := range a.routes {
		regFn(baseGroup.Group(path))
	}
	if a.cfg.IsLowEnv() {
		baseGroup.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	return &App{
		cfg:     a.cfg,
		server:  e,
		ctx:     a.ctx,
		address: fmt.Sprintf(":%s", a.cfg.App.Port),
	}, nil
}

func (a *App) Start() error {
	return a.server.Start(a.address)
}

func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(a.ctx, 10*time.Second)
	defer cancel()
	return a.server.Shutdown(ctx)
}

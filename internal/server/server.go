package server

import (
	"solid-task-management/internal/handler"
	"solid-task-management/internal/middleware"
	"solid-task-management/internal/server/routes"

	"solid-task-management/pkg/logger"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func NewServer(
	h *handler.HandlerRegistry,
) *Server {
	logger, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	e := echo.New()

	e.Use(middleware.RequestID)
	e.Use(middleware.ZapLogger(logger))
	e.Use(middleware.RecoverWithZap(logger))
	routes.RegisterTaskRoutes(e.Group("/tasks"), h.TaskHandler)

	return &Server{e: e}
}

func (s *Server) Run() error {
	return s.e.Start(":8080")
}

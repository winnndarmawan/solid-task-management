package routes

import (
	"solid-task-management/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterTaskRoutes(e *echo.Group, h *handler.TaskHandler) {
	e.POST("/create", h.Create)
	e.GET("/list", h.Get)
	e.PATCH("/list", h.Update)

}

package handler

import (
	"fmt"
	"net/http"
	dom "solid-task-management/internal/domain"
	"solid-task-management/internal/service/task"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.Service
}

func TaskHandlerProvider(
	taskService task.Service,
) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) Create(ec echo.Context) error {
	var req CreateTaskRequest
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, "BAD_REQUEST")
	}

	ctx := ec.Request().Context()
	task, err := h.taskService.CreateNewTask(ctx, dom.CreateReq{
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "CREATE_FAILED")
	}

	return ec.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Get(ec echo.Context) error {
	var req FetchTaskRequest

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, "BAD_REQUEST")
	}
	fmt.Print(req)
	ctx := ec.Request().Context()
	task, err := h.taskService.FetchTasks(ctx, dom.FetchTasksReq{
		Title:       req.Title,
		Description: req.Description,
		Page:        req.Page,
		PerPage:     req.PerPage,
	})

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "CREATE_FAILED")
	}

	return ec.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Update(ec echo.Context) error {
	var req UpateTaskRequest
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, "BAD_REQUEST")
	}

	ctx := ec.Request().Context()
	task, err := h.taskService.UpdateTask(ctx, dom.UpdateReq{
		Title:       req.Title,
		Description: req.Description,
		ID:          req.ID,
		Status:      dom.TaskStatus(req.Status),
	})

	if err != nil {
		return ec.JSON(http.StatusInternalServerError, "CREATE_FAILED")
	}

	return ec.JSON(http.StatusCreated, task)
}

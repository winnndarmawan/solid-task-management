package task

import (
	"context"
	"fmt"

	task "solid-task-management/internal/domain"
	"solid-task-management/internal/mongoose"
)

type (
	Service interface {
		CreateNewTask(ctx context.Context, req task.CreateReq) (*task.Task, error)
		FetchTasks(ctx context.Context, req task.FetchTasksReq) ([]*task.Task, error)
	}

	serviceImpl struct {
		TaskAdapter mongoose.TaskRepository
	}
)

func TaskServiceProvider(
	taskAdapter mongoose.TaskRepository,
) Service {
	return serviceImpl{TaskAdapter: taskAdapter}
}

func (s serviceImpl) CreateNewTask(ctx context.Context, req task.CreateReq) (*task.Task, error) {
	print("cek1")
	if err := req.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return nil, err
	}
	print("cek2")
	return s.TaskAdapter.Create(ctx, req)
}

func (s serviceImpl) FetchTasks(ctx context.Context, req task.FetchTasksReq) ([]*task.Task, error) {
	fmt.Print(req)
	return s.TaskAdapter.FindAll(ctx, req)
}

package task

import (
	"context"

	task "solid-task-management/internal/domain"
	"solid-task-management/internal/mongoose"
)

type (
	Service interface {
		CreateNewTask(ctx context.Context, req task.CreateReq) (*task.Task, error)
		FetchTasks(ctx context.Context, req task.FetchTasksReq) ([]*task.Task, error)
		UpdateTask(ctx context.Context, req task.UpdateReq) (*task.Task, error)
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
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return s.TaskAdapter.Create(ctx, req)
}

func (s serviceImpl) FetchTasks(ctx context.Context, req task.FetchTasksReq) ([]*task.Task, error) {
	return s.TaskAdapter.FindAll(ctx, req)
}

func (s serviceImpl) UpdateTask(ctx context.Context, req task.UpdateReq) (*task.Task, error) {
	return s.TaskAdapter.Update(ctx, req)
}

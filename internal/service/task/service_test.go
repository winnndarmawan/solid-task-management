package task

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	task "solid-task-management/internal/domain"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) FindOne(ctx context.Context, id string) (*task.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) Create(ctx context.Context, req task.CreateReq) (*task.Task, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) FindAll(ctx context.Context, req task.FetchTasksReq) ([]*task.Task, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*task.Task), args.Error(1)
}

func TestCreateNewTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := TaskServiceProvider(mockRepo)

	ctx := context.Background()
	req := task.CreateReq{
		Title:       "Test Task",
		Description: "Test Description",
	}

	expectedTask := &task.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "Test Description",
	}

	mockRepo.On("Create", ctx, req).Return(expectedTask, nil)

	result, err := service.CreateNewTask(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateNewTask_ValidationError(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := TaskServiceProvider(mockRepo)

	ctx := context.Background()
	req := task.CreateReq{
		Title:       "", // Invalid title
		Description: "Test Description",
	}

	result, err := service.CreateNewTask(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "Create")
}

func TestFetchTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := TaskServiceProvider(mockRepo)

	ctx := context.Background()
	req := task.FetchTasksReq{
		Title:       "Test Task",
		Description: "Test Description",
	}

	expectedTasks := []*task.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Test Task",
			Description: "Test Description",
		},
	}

	mockRepo.On("FindAll", ctx, req).Return(expectedTasks, nil)

	result, err := service.FetchTasks(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, result)
	mockRepo.AssertExpectations(t)
}

func TestFetchTasks_Error(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := TaskServiceProvider(mockRepo)

	ctx := context.Background()
	req := task.FetchTasksReq{
		Title:       "Test Task",
		Description: "Test Description",
	}

	mockRepo.On("FindAll", ctx, req).Return(nil, errors.New("database error"))

	result, err := service.FetchTasks(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

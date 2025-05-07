//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"solid-task-management/internal/db"
	"solid-task-management/internal/handler"
	"solid-task-management/internal/mongoose"
	"solid-task-management/internal/server"
	"solid-task-management/internal/service/task"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		db.ConnectMongo,

		// Tasks
		mongoose.TaskRepositoryProvider,
		task.TaskServiceProvider,
		handler.TaskHandlerProvider,

		// Group all handlers
		handler.HandlerRegistryProvider,

		// Server
		server.NewServer,
	)
	return &server.Server{}, nil
}

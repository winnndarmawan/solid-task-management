package handler

type HandlerRegistry struct {
	TaskHandler *TaskHandler
}

func HandlerRegistryProvider(
	taskHandler *TaskHandler,
) *HandlerRegistry { // Change return type to *HandlerRegistry
	return &HandlerRegistry{ // Return a pointer
		TaskHandler: taskHandler,
	}
}

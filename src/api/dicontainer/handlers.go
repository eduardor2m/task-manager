package dicontainer

import "github.com/eduardor2m/task-manager/src/api/handlers"

func GetTaskHandlers() *handlers.TaskHandlers {
	return handlers.NewTaskHandlers(
		GetTaskServices(),
	)
}

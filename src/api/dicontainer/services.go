package dicontainer

import (
	"github.com/eduardor2m/task-manager/src/core/interfaces/primary"
	"github.com/eduardor2m/task-manager/src/core/services"
)

func GetTaskServices() primary.TaskManager {
	return services.NewTaskServices(
		GetTaskRepository(),
	)
}

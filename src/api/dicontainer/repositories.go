package dicontainer

import (
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/eduardor2m/task-manager/src/infra/sqlite"
)

func GetTaskRepository() repository.TaskLoader {
	return sqlite.NewTaskSQLiteRepository(GetSQLiteConnectionManager())
}

func GetSQLiteConnectionManager() *sqlite.DatabaseConnectionManager {
	return &sqlite.DatabaseConnectionManager{}
}

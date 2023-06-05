package dicontainer

import (
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/eduardor2m/task-manager/src/infra/postgres"
)

func GetTaskRepository() repository.TaskLoader {
	return postgres.NewTaskSQLiteRepository(GetSQLiteConnectionManager())
}

func GetUserRepository() repository.UserLoader {
	return postgres.NewUserSQLiteRepository(GetSQLiteConnectionManager())
}

func GetSQLiteConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}

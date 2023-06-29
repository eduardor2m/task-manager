package router

import (
	"github.com/eduardor2m/task-manager/src/api/dicontainer"
	"github.com/labstack/echo/v4"
)

func loadTaskRoutes(group *echo.Group) {
	taskGroup := group.Group("/task")

	taskHandlers := dicontainer.GetTaskHandlers()

	taskGroup.GET("", taskHandlers.GetTasks)
	taskGroup.GET("/:id", taskHandlers.GetTask)
	taskGroup.POST("", taskHandlers.CreateTask)
	taskGroup.PUT("/:id", taskHandlers.UpdateTask)
	taskGroup.PUT("/:id/status", taskHandlers.UpdateTaskStatus)
	taskGroup.DELETE("/:id", taskHandlers.DeleteTask)
	taskGroup.DELETE("", taskHandlers.DeleteTasks)
}

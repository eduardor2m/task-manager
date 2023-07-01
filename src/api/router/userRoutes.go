package router

import (
	"github.com/eduardor2m/task-manager/src/api/dicontainer"
	"github.com/labstack/echo/v4"
)

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")

	userHandlers := dicontainer.GetUserHandlers()

	userGroup.POST("/signup", userHandlers.SignUp)
	userGroup.POST("/signin", userHandlers.SignIn)
	userGroup.DELETE("/delete", userHandlers.DeleteUserByEmail)
}

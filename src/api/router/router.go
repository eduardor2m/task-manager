package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/eduardor2m/task-manager/src/api/docs"
)

type Router interface {
	Load(*echo.Group)
}

type router struct{}

func New() Router {
	return &router{}
}

func (instance *router) Load(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)

	loadTaskRoutes(group)
	loadUserRoutes(group)
}

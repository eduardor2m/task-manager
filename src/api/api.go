package api

import (
	"net/http"

	"github.com/eduardor2m/task-manager/src/api/middlewares"
	"github.com/eduardor2m/task-manager/src/api/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API interface {
	Serve()
	loadRoutes()
}

type Options struct{}

type api struct {
	options      *Options
	group        *echo.Group
	echoInstance *echo.Echo
}

func NewApi(options *Options) API {
	echoInstance := echo.New()
	return &api{
		options:      options,
		group:        echoInstance.Group("/api"),
		echoInstance: echoInstance,
	}
}

func (instance *api) Serve() {
	instance.loadRoutes()
	instance.echoInstance.Use(instance.getCORSSettings())
	instance.echoInstance.Use(middlewares.GuardMiddleware)
	instance.echoInstance.Logger.Fatal(instance.echoInstance.Start(":9090"))
}

func (instance *api) loadRoutes() {
	router := router.New()
	router.Load(instance.group)
}

func (instance *api) getCORSSettings() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:         middlewares.OriginInspectSkipper,
		AllowOriginFunc: middlewares.VerifyOrigin,
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
	})
}

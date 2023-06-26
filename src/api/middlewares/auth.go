package middlewares

import (
	"strings"

	token "github.com/eduardor2m/task-manager/src/api/handlers/utils"
	"github.com/labstack/echo/v4"
)

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	// se a rota for api/docs algo, passa

	return func(c echo.Context) error {
		docsURL := strings.HasPrefix(c.Request().URL.Path, "/api/docs")
		if c.Request().URL.Path == "/api/user/signin" {
			return next(c)
		} else if c.Request().URL.Path == "/api/user/signup" {
			return next(c)

		} else if docsURL {
			return next(c)
		}

		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		splitToken := strings.Split(tokenString, "Bearer ")

		if len(splitToken) != 2 {
			return echo.ErrUnauthorized
		}

		tokenString = splitToken[1]

		token := token.VerifyToken(tokenString)

		if !token {
			return echo.ErrUnauthorized
		}

		return next(c)
	}

}

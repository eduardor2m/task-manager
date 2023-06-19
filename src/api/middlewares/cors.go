package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// VerifyOrigin verifies if the request origin is included on the defined server
// allowed hosts.
func VerifyOrigin(origin string) (bool, error) {
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:9090",
	}

	for _, allowedOrigin := range allowedOrigins {
		if strings.Compare(origin, allowedOrigin) == 0 {
			return true, nil
		}
	}

	return false, &echo.HTTPError{Code: 401, Message: "you're not allowed to access this API"}
}

// OriginInspectSkipper verifies the request context and skip the origin verification.
// It's useful to allow access for any origin when a route (e.g. public images routes)
// should be accessed by anyone.
func OriginInspectSkipper(context echo.Context) bool {
	return false
}

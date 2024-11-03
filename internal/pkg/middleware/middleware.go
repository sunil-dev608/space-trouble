package middleware

import (
	"github.com/labstack/echo/v4"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/sunil-dev608/space-trouble/internal/pkg/response"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			_ = token
			// if token == "" {
			// 	return response.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", nil)
			// }

			// Validate JWT token...

			return next(c)
		}
	}
}

package middleware

import (
	"net/http"
	"project2/helpers"

	"github.com/labstack/echo"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return err
		}
		c.Set("userData", verifyToken)
		return next(c)
	}
}

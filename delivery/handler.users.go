package delivery

import (
	"context"
	"net/http"
	"project2/usecases"

	"github.com/labstack/echo"
)

func RegisterUser(ctx context.Context, uc usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.RegisterUserRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		res, httpcode, err := uc.RegisterUser(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToRegister,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

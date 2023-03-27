package delivery

import (
	"context"
	"project2/usecases"

	"github.com/labstack/echo"
)

const (
	FailedToUnmarshall = "Failed to Unmarshall"
	TokenIsRequired    = "Token must be provided"
	SuccessMsg         = "Success"
	WelcomeMsg         = "welcome"
)

type echoObject struct {
	*echo.Echo
	UseCase
}

type UseCase struct {
	usecases.UserUseCase
}

func NewEchoHandler(ctx context.Context, c *echo.Echo, uc UseCase) {
	obj := &echoObject{c, uc}
	obj.initRoute(ctx)
}

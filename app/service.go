package app

import (
	"context"
	"project2/delivery"
	"project2/usecases"
	users_case "project2/usecases/users"
	"time"
)

func (x *App) initService(ctx context.Context) {
	timeout := 55 * time.Second

	usercase := users_case.New(
		users_case.Configuration{
			Timeout: timeout,
		},
		users_case.Depencency{
			Postgresql: x.DB,
		},
	)

	delivery.NewEchoHandler(ctx, x.Echo, struct {
		usecases.UserUseCase
	}{
		usercase,
	})
}

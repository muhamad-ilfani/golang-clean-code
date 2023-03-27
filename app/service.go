package app

import (
	"context"
	"project2/delivery"
	"project2/usecases"
	users_case "project2/usecases/users"
)

func (x *App) initService(ctx context.Context) {

	usercase := users_case.New(
		users_case.Configuration{},
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

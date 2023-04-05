package app

import (
	"context"
	"project2/delivery"
	"project2/delivery/kafka"
	"project2/usecases"
	users_case "project2/usecases/users"
	"time"
)

func (x *App) initService(ctx context.Context) (err error) {
	timeout := 55 * time.Second

	kafka.InitSubscriptions(ctx, kafka.Usecase{})

	usercase := users_case.New(
		users_case.Configuration{
			Timeout: timeout,
		},
		users_case.Depencency{
			Postgresql:    x.DB,
			KafkaProducer: x.kafkaProducer,
		},
	)

	delivery.NewEchoHandler(ctx, x.Echo, struct {
		usecases.UserUseCase
	}{
		usercase,
	})

	return nil
}

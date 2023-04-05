package repository

import "context"

type NotifyRegistration interface {
	NotifyRegistration(
		ctx context.Context, req NotifyRegistrationRequest) (
		err error,
	)
}

type NotifyRegistrationRequest struct {
	UserName string
	Email    string
}

type NotifyRegistrationResponse struct{}

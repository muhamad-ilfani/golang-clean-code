package repository

import "context"

type UserRepo interface {
	RegisterUser(
		ctx context.Context, req RegisterUserRequest) (
		res RegisterUserResponse, httpcode int, err error,
	)
	LoginUser(
		ctx context.Context, req LoginUserRequest) (
		res LoginUserResponse, httpcode int, err error,
	)
}

type RegisterUserRequest struct {
}

type RegisterUserResponse struct {
}

type LoginUserRequest struct {
}

type LoginUserResponse struct {
}

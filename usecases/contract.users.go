package usecases

import "context"

type UserUseCase interface {
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
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
}

type LoginUserResponse struct {
}

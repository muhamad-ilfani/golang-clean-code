package users_case

import (
	"context"
	"project2/usecases"
)

func (x *usecase) RegisterUser(
	ctx context.Context, req usecases.RegisterUserRequest) (
	res usecases.RegisterUserResponse, httpcode int, err error,
) {

	return res, httpcode, err
}

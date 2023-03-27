package users_case

import (
	"context"
	"project2/usecases"
)

func (x *usecase) LoginUser(
	ctx context.Context, req usecases.LoginUserRequest) (
	res usecases.LoginUserResponse, httpcode int, err error,
) {

	return res, httpcode, err
}

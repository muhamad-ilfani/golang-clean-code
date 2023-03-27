package user_repo

import (
	"context"
	"project2/repository"
)

func (x *PostgreSQLConn) RegisterUser(
	ctx context.Context, req repository.RegisterUserRequest) (
	res repository.RegisterUserResponse, httpcode int, err error,
) {
	return res, httpcode, err
}

func (x *PostgreSQLConn) LoginUser(
	ctx context.Context, req repository.LoginUserRequest) (
	res repository.LoginUserResponse, httpcode int, err error,
) {
	return res, httpcode, err
}

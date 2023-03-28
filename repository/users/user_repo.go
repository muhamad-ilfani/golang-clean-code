package user_repo

import (
	"context"
	"net/http"
	"project2/repository"
	"project2/repository/users/query"
)

func (x *PostgreSQLConn) RegisterUser(
	ctx context.Context, req repository.RegisterUserRequest) (
	res repository.RegisterUserResponse, httpcode int, err error,
) {
	query := query.InsertUserRegistration
	args := List{
		req.UserName,
		req.Email,
		req.Password,
	}

	_, err = x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) LoginUser(
	ctx context.Context, req repository.LoginUserRequest) (
	res repository.LoginUserResponse, httpcode int, err error,
) {
	return res, httpcode, err
}

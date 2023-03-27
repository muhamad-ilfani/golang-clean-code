package user_repo

import (
	"project2/repository"

	"github.com/jmoiron/sqlx"
)

type PostgreSQLConn struct {
	tc *sqlx.DB
}

type Repository interface {
	repository.UserRepo
}

func NewRepository(tc *sqlx.DB) Repository { return &PostgreSQLConn{tc} }

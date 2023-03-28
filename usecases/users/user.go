package users_case

import (
	"errors"
	"project2/usecases"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	EmailNotValid    = errors.New("Email not valid")
	PasswordNotValid = errors.New("Password not valid")
)

func New(c Configuration, d Depencency) usecases.UserUseCase {
	return &usecase{c, d}
}

type Configuration struct {
	Timeout time.Duration
}

type Depencency struct {
	Postgresql *sqlx.DB
}

type usecase struct {
	Configuration
	Depencency
}

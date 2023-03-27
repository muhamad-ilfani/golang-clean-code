package users

import (
	"project2/usecases"
	"time"

	"github.com/jmoiron/sqlx"
)

func New(c Configuration, d Depencency) usecases.UserUseCase {
	return &usecase{c, d}
}

type Configuration struct {
	Timeout time.Time
}

type Depencency struct {
	Postgresql *sqlx.DB
}

type usecase struct {
	Configuration
	Depencency
}

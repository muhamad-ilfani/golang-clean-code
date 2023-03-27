package app

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type App struct {
	DB   *sqlx.DB
	Echo *echo.Echo
}

func Run(ctx context.Context) {
	app := new(App)

	if err := app.initDB(ctx); err != nil {
		//error
	}

	app.initService(ctx)
}

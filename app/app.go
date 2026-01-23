package app

import (
	"go-backend/handler"
	"go-backend/repository"
	"go-backend/service"
	"database/sql"
)

type App struct {
	UserHandler     handler.UserHandler
	LoginHandler    handler.LoginHandler
	RegisterHandler handler.RegisterHandler
}

func NewApp(db *sql.DB) *App {
	userRepo := repository.NewUserRepoPostgres(db)

	userService := service.UserService{
		Repo: userRepo,
	}

	return &App{
		UserHandler: handler.UserHandler{
			Service: userService,
		},
		LoginHandler: handler.LoginHandler{
			Service: userService,
		},
		RegisterHandler: handler.RegisterHandler{
			Service: userService,
		},
	}
}

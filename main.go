package main

import (
	"go-backend/handler"
	"go-backend/repository"
	"go-backend/service"
	"os"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go-backend/routes"

)

func main() {
	

	
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	dns := os.Getenv("DATABASE_URL")
	if dns == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	db, err := repository.NewPostgresDB(dns)
	if  err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	userRepo := repository.NewUserRepoPostgres(db)
	
	UserService := service.UserService{
		Repo: userRepo,
	}
	UserHandler := handler.UserHandler{
		Service:  UserService,
	}
	LoginHandler := handler.LoginHandler{
		Service: UserService,
	}
	RegisterHandler := handler.RegisterHandler{
		Service: UserService,
	}

	muxWithMiddleware := routes.SetupRoutes(
		UserHandler,
		LoginHandler,
		RegisterHandler,
	)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", muxWithMiddleware))
}

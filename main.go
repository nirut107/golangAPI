package main

import (
	"go-backend/handler"
	"go-backend/repository"
	"go-backend/service"
	"os"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go-backend/middleware"
)

func main() {
	mux := http.NewServeMux()

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
	userRepo := repository.PostgresUserRepo{DB: db}
	
	UserService := service.UserService{
		Repo: userRepo,
	}
	UserHandler := handler.UserHandler{
		Service:  UserService,
	}

	mux.HandleFunc("/users", UserHandler.Users)
	muxWithMiddleware := middleware.LoggingMiddleware(mux)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", muxWithMiddleware))
}

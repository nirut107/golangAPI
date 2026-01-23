package routes

import (
	"net/http"

	"go-backend/handler"
	"go-backend/middleware"
)

func SetupRoutes(
	userHandler handler.UserHandler,
	loginHandler handler.LoginHandler,
	registerHandler handler.RegisterHandler,
) http.Handler {

	// ===== public routes (ไม่ต้อง auth) =====
	publicMux := http.NewServeMux()
	publicMux.HandleFunc("/login", loginHandler.Login)
	publicMux.HandleFunc("/register", registerHandler.Register)

	// ===== protected routes (ต้อง auth) =====
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/users", userHandler.Users)


	protectedHandler := middleware.AuthMiddleware(protectedMux)

	// ===== root mux =====
	rootMux := http.NewServeMux()
	rootMux.Handle("/login", publicMux)
	rootMux.Handle("/register", publicMux)
	rootMux.Handle("/", protectedHandler)

	// ===== global middleware =====
	finalHandler := middleware.LoggingMiddleware(rootMux)

	return finalHandler
}

package handler

import (
	"encoding/json"
	"go-backend/model"
	"go-backend/service"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	Service service.UserService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "username and password required", http.StatusBadRequest)
		return
	}
	
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	req.Password = string(hashedPassword)	
	newUser := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	
	createdUser, err := h.Service.Create(newUser)
	if err != nil {
		http.Error(w,  "User already exists", http.StatusConflict)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"id":      createdUser.ID,
		"user":    createdUser,

	})
}

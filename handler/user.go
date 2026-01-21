package handler

import (
	"net/http"
	"strconv"

	"encoding/json"
	"go-backend/model"
	"go-backend/repository"
	"go-backend/service"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	Service service.UserService
}

// Users handles HTTP requests for user operations
func (h UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method{
	// Get all users or get user by ID
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")

		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			user, err := h.Service.GetByID(id)
			if err != nil {
				if err == repository.ErrUserNotFound {
					http.Error(w, "user not found", http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(user)
			return
		}
		users, err := h.Service.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	// Create a new user
	case http.MethodPost:
		var u model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		create, err := h.Service.Create(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(create)
	// Delete a user by ID	
	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.Service.Delete(id)
		if err != nil {
			if err == repository.ErrUserNotFound {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	// Update a user
	case http.MethodPut:
		var u model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedUser, err := h.Service.Update(u)
		if err != nil {
			if err == repository.ErrUserNotFound {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedUser)
	// 	Method not allowed
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}




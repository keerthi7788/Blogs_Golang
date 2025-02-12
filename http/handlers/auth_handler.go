package handlers

import (
	"Blogs/modals"
	"Blogs/service"
	"encoding/json"
	"net/http"
)

type Authhandler struct {
	service service.UserService
}

func NewAuthHandlers(service service.UserService) *UserHandler {
	return &UserHandler{service: service}

}
func (a *Authhandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var loginRequest modals.LoginModal
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	validationErrors := loginRequest.Validate()
	if validationErrors != nil {
		json.NewEncoder(w).Encode(validationErrors)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var user modals.Users
	_, err = a.service.GetUserByEmail(r.Context(), user.Email)
	if err != nil {
		// handle this in repo layer
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

}

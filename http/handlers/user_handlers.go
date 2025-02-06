package handlers

import (
	"Blogs/modals"
	"Blogs/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandlers(service service.UserService) *UserHandler {
	return &UserHandler{service: service}

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var users modals.Users
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "invalid response body", http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateUser(r.Context(), users)
	if err != nil {
		http.Error(w, "unable to crete the user", http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.Hex()})
}
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "unable to get the users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	users, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "No users found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)

}
func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	deletedUser, err := h.service.DeleteUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Books NotfOUND ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedUser)

}
func (h *UserHandler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.DeleteAllUsers(r.Context())
	if err != nil {
		http.Error(w, "unable to delete the users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)

}
func (h *UserHandler) UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}
	var userDetails modals.Users
	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		http.Error(w, "invalid response", http.StatusBadRequest)
	}
	updatedUserDetails, err := h.service.UpdateUserDetails(r.Context(), id, userDetails)
	if err != nil {
		http.Error(w, "unable to get the users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUserDetails)

}
func (h *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	//  id := primitive.ObjectIDFromHex("id")
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	var userDetails modals.Users
	err :=
		json.NewDecoder(r.Body).Decode(&userDetails)

	if err != nil {
		http.Error(w, "invalid JSON request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.UpdateUserDetails(r.Context(), id, userDetails)
	if err != nil {
		http.Error(w, "internalserver error:", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contant-Type", "applicaton/json")
	json.NewEncoder(w).Encode(&user)

}

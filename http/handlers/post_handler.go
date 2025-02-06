package handlers

import (
	"Blogs/modals"
	"Blogs/service"
	"encoding/json"
	"net/http"
)

type PostHandler struct {
	service service.PostService
}

func NewpostHandlers(service service.PostService) *PostHandler {
	return &PostHandler{service: service}

}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var posts modals.Post
	err := json.NewDecoder(r.Body).Decode(&posts)
	if err != nil {
		http.Error(w, "invalid response body", http.StatusBadRequest)
		return
	}
	id, err := h.service.CreatePost(r.Context(), posts)
	if err != nil {
		http.Error(w, "unable to crete the post", http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.ID.Hex()})
}
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, "unable to get the posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}
func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	posts, err := h.service.GetPostByID(r.Context(), id)
	if err != nil {
		http.Error(w, "No posts found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&posts)

}
func (h *PostHandler) DeletePostById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	deletedUser, err := h.service.DeletePostById(r.Context(), id)
	if err != nil {
		http.Error(w, "Books NotfOUND ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedUser)

}
func (h *PostHandler) DeleteAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.DeleteAllPosts(r.Context())
	if err != nil {
		http.Error(w, "unable to delete the books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}
func (h *PostHandler) UpdateAllPosts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}
	var userDetails modals.Post
	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		http.Error(w, "invalid response", http.StatusBadRequest)
	}
	updatedUserDetails, err := h.service.UpdateAllPosts(r.Context(), id, userDetails)
	if err != nil {
		http.Error(w, "unable to get the posts", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUserDetails)

}
func (h *PostHandler) UpdatePostById(w http.ResponseWriter, r *http.Request) {
	//  id := primitive.ObjectIDFromHex("id")
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	var userDetails modals.Post
	err :=
		json.NewDecoder(r.Body).Decode(&userDetails)

	if err != nil {
		http.Error(w, "invalid JSON request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.UpdatePostById(r.Context(), id, userDetails.Content, userDetails.Title_Post)
	if err != nil {
		http.Error(w, "internalserver error:", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contant-Type", "applicaton/json")
	json.NewEncoder(w).Encode(&user)

}

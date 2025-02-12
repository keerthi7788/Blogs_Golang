package handlers

import (
	"Blogs/modals"
	"Blogs/service"
	"encoding/json"
	"net/http"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandlers(service service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}

}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comments modals.Comment
	err := json.NewDecoder(r.Body).Decode(&comments)
	if err != nil {
		http.Error(w, "invalid response body", http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateComment(r.Context(), comments)
	if err != nil {
		http.Error(w, "unable to crete thecommant", http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.ID.Hex()})
}
func (h *CommentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAllComments(r.Context())
	if err != nil {
		http.Error(w, "unable to get the books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}
func (h *CommentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	comments, err := h.service.GetCommentByID(r.Context(), id)
	if err != nil {
		http.Error(w, "No books found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)

}
func (h *CommentHandler) DeleteCommentById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	deletedcomments, err := h.service.DeleteCommentById(r.Context(), id)
	if err != nil {
		http.Error(w, "Books NotfOUND ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedcomments)

}
func (h *CommentHandler) DeleteAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.service.DeleteAllComments(r.Context())
	if err != nil {
		http.Error(w, "unable to delete the books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)

}
func (h *CommentHandler) UpdateAllComments(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}
	var commentDetails modals.Comment
	err := json.NewDecoder(r.Body).Decode(&commentDetails)
	if err != nil {
		http.Error(w, "invalid response", http.StatusBadRequest)
	}
	updatedcommentDetails, err := h.service.UpdateAllComments(r.Context(), id, commentDetails)
	if err != nil {
		http.Error(w, "unable to get the books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedcommentDetails)

}
func (h *CommentHandler) UpdateCommentById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	var userDetails modals.Comment
	err :=
		json.NewDecoder(r.Body).Decode(&userDetails)

	if err != nil {
		http.Error(w, "invalid JSON request body", http.StatusBadRequest)
		return
	}

	comment, err := h.service.UpdateCommentById(r.Context(), id, userDetails.Content)
	if err != nil {
		http.Error(w, "internalserver error:", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contant-Type", "applicaton/json")
	json.NewEncoder(w).Encode(&comment)

}

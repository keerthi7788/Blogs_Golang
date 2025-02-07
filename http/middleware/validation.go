package middleware

import (
	"Blogs/modals"
	"encoding/json"
	"net/http"
	"regexp"
)

func validate(Next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		var user modals.Users
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if user.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}
		if user.Email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}
		if !isValidEmail(user.Email) {
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}
		Next.ServeHTTP(w, r)
	})
}
func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

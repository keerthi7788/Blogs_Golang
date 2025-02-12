package middleware

import (
	"Blogs/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthJWT(Next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusBadRequest)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("my_secret_key"), nil

		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// check if the token is expired
		if err == jwt.ErrTokenExpired {
			NewToken, err := utils.CreateJwtToken(claims.UserID)
			if err != nil {
				http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
				return
			}
			// Return new token in response
			Response := map[string]string{
				"message":   "token expire and generated new token",
				"new_token": NewToken,
			}
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(Response)

		}

	})
}

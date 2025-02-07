package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var Secretekey = []byte("Config.jwt.secret")

type Claims struct {
	UserID string `json:"userid,omitempty" bson:"userid,omitempty"`
	jwt.RegisteredClaims
}

func CreateJwtToken(userID string) (string, error) {
	ExpirationTime := time.Now().Add(15 * time.Minute)
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secretekey)
}

// JWT Middleware with Expiration Handling
func AuthJWT(Next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
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
		// check if the token is expired
		if err == jwt.ErrTokenExpired {
			NewToken, err := CreateJwtToken(claims.UserID)
			if err != nil {
				http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
				return
			}
			// Return new token in response
			Response := map[string]string{"message": "token expire and generated new token", "new_token": NewToken}
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(Response)

		}
		http.Error(w, "invalid token", http.StatusUnauthorized)
	})
}

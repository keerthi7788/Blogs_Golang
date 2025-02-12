package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key (should be read from config)
var Secretekey = []byte("your_secret_key") // Replace with actual key from config

// Claims structure
type Claims struct {
	UserID string `json:"userid,omitempty" bson:"userid,omitempty"`
	jwt.RegisteredClaims
}

// CreateJwtToken generates a JWT token and sets it in an HTTP-only cookie
func CreateJwtToken(userID string) (string, error) {
	// Set expiration time (15 minutes)
	ExpirationTime := time.Now().Add(15 * time.Minute)
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpirationTime),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Secretekey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func RefreshToken(userID string) (string, error) {
	ExpirationTime := time.Now().Add(15 * time.Minute)
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpirationTime),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Secretekey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// JWT Middleware with Expiration Handling

func HomePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenstring := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenstring, claims,
		func(t *jwt.Token) (interface{}, error) {
			return Secretekey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello %s", claims.UserID)))

}

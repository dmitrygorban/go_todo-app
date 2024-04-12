package middlewares

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := os.Getenv("TODO_PASSWORD")
		secret := os.Getenv("TODO_SECRET")

		if secret == "" {
			secret = "dafault_secret" // only for dev purposes
		}

		if len(password) > 0 {
			var token string

			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}

			if token == "" {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !jwtToken.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := jwtToken.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			hashClaim := claims["hash"]

			hashClaim, ok = hashClaim.(string)
			if !ok {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			passwordHash := sha256.Sum256([]byte(password))
			passwordHashStr := hex.EncodeToString(passwordHash[:])

			if hashClaim != passwordHashStr {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
		}

		next(w, r)
	})
}

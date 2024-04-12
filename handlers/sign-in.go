package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type SignInRequest struct {
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := os.Getenv("TODO_PASSWORD")
	if req.Password != password {
		handleError(w, "wrong password", http.StatusUnauthorized)
		return
	}
	hash := sha256.Sum256([]byte(password))
	hashStr := hex.EncodeToString(hash[:])
	claims := jwt.MapClaims{
		"hash": hashStr,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("TODO_SECRET")
	if secret == "" {
		secret = "default_secret"
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("token generated", tokenString)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(SignInResponse{Token: tokenString})
}

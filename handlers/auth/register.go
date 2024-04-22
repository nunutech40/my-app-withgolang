package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nunutech40/my-app-withgolang/common/response"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct { // handler can call Register
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler { // this return Handler
	return &Handler{DB: db}
}

// Struct for get body
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) { // with (h *Handler), kita bisa call Register dari struct handler

	// check the method use POST for register
	if r.Method != http.MethodPost {
		response.SendJsonResponse(w, http.StatusMethodNotAllowed, "Method not Allowed", nil)
		return
	}

	// write the body, body will have three input like user
	// body will have three input like user struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.SendJsonResponse(w, http.StatusBadRequest, "Invalid Request Body", nil)
		return
	}

	// check for empty fields
	if user.Username == "" || user.Password == "" || user.Email == "" {
		response.SendJsonResponse(w, http.StatusBadRequest, "Missing fields: username, password, or email", nil)
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.SendJsonResponse(w, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	// Insert the user into the database
	stmt := `INSERT INTO users (username, password_hash, email, created_at, last_login) VALUES (?, ?, ?, ?, ?)`
	_, err = h.DB.Exec(stmt, user.Username, hashedPassword, user.Email, time.Now(), nil)
	if err != nil {
		response.SendJsonResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err), nil)
		return
	}

	// Successfull registration
	response.SendJsonResponse(w, http.StatusCreated, "User registered successfully", nil)

}

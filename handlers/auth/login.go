package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/nunutech40/my-app-withgolang/common/response"
	"github.com/nunutech40/my-app-withgolang/handlers"
	"golang.org/x/crypto/bcrypt"
)

func Login(h *handlers.Handler, w http.ResponseWriter, r *http.Request) {
	// check method, and only can use post method
	if r.Method != http.MethodPost {
		response.SendJsonResponse(w, http.StatusMethodNotAllowed, "Method not Allowed", nil)
		return
	}

	// credential struct for body request
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// create body request from credential
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		response.SendJsonResponse(w, http.StatusBadRequest, "Invalid Request Body", nil)
		return
	}

	// validation input credential
	if credentials.Username == "" || credentials.Password == "" {
		response.SendJsonResponse(w, http.StatusBadRequest, "Missing username or password", nil)
		return
	}

	// fetch credential (usrname & password) is match with database (user table)
	var storedHash string
	var userID int
	stmt := `SELECT id, password_hash FROM users WHERE username=?`
	err = h.DB.QueryRow(stmt, credentials.Username).Scan(&userID, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			response.SendJsonResponse(w, http.StatusUnauthorized, "Invalid username or password", nil)
		} else {
			response.SendJsonResponse(w, http.StatusInternalServerError,
				"Database Error", nil)
		}
		return
	}

	// compared hashes password
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(credentials.Password))
	if err != nil {
		response.SendJsonResponse(w, http.StatusUnauthorized, "Invalid username and password", nil)
		return
	}

	// update last login
	stmt = "UPDATE users SET last_login=? WHERE id=?"
	_, err = h.DB.Exec(stmt, time.Now(), userID)
	if err != nil {
		response.SendJsonResponse(w, http.StatusInternalServerError, "Failed to update last login", nil)
		return
	}

	// send json response to client
	response.SendJsonResponse(w, http.StatusOK, "Login successful", nil)

}

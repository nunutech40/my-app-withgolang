package auth

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Handler struct { // handler can call Register
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler { // this return Handler
	return &Handler{DB: db}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) { // with (h *Handler), kita bisa call Register dari struct handler
	fmt.Fprintf(w, "Testing register API")
}

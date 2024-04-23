package handlers

import (
	"database/sql"
)

// handler for DB
type Handler struct { // handler can call Register
	DB *sql.DB
}

// Handler for injection
func NewHandler(db *sql.DB) *Handler { // this return Handler
	return &Handler{DB: db}
}

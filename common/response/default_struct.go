package response

import (
	"encoding/json"
	"net/http"
)

// struct for response
type JSONResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendJsonResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := JSONResponse{
		Status:  code,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

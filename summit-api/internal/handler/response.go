package handler

import (
	"encoding/json"
	"net/http"

	"github.com/summit/summit-api/pkg/apperror"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		writeJSON(w, appErr.Code, map[string]string{"message": appErr.Message})
		return
	}
	writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "internal server error"})
}

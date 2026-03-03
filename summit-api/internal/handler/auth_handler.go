package handler

import (
	"net/http"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/service"
	"github.com/summit/summit-api/pkg/apperror"
	"github.com/summit/summit-api/pkg/validator"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Service() *service.AuthService {
	return h.service
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	resp, err := h.service.Login(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	resp, err := h.service.Register(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

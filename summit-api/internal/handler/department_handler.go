package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/summit/summit-api/internal/service"
	"github.com/summit/summit-api/pkg/apperror"
)

type DepartmentHandler struct {
	service *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: s}
}

func (h *DepartmentHandler) List(w http.ResponseWriter, r *http.Request) {
	depts, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, depts)
}

func (h *DepartmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid department id"))
		return
	}

	dept, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, apperror.NotFound("department not found"))
		return
	}
	writeJSON(w, http.StatusOK, dept)
}

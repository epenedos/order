package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/summit/summit-api/internal/service"
	"github.com/summit/summit-api/pkg/apperror"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) List(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid employee id"))
		return
	}

	employee, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, apperror.NotFound("employee not found"))
		return
	}

	writeJSON(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) ListSalesReps(w http.ResponseWriter, r *http.Request) {
	reps, err := h.service.ListSalesReps(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, reps)
}

func (h *EmployeeHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid employee id"))
		return
	}

	customers, err := h.service.GetCustomersBySalesRep(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, customers)
}

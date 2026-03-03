package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/service"
	"github.com/summit/summit-api/pkg/apperror"
	"github.com/summit/summit-api/pkg/pagination"
	"github.com/summit/summit-api/pkg/validator"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (h *CustomerHandler) List(w http.ResponseWriter, r *http.Request) {
	filter := models.CustomerFilter{
		SortBy: r.URL.Query().Get("sort"),
	}
	if country := r.URL.Query().Get("country"); country != "" {
		filter.Country = &country
	}
	if repID := r.URL.Query().Get("sales_rep_id"); repID != "" {
		id, _ := strconv.Atoi(repID)
		filter.SalesRepID = &id
	}

	pg := pagination.FromRequest(r)
	customers, total, err := h.service.List(r.Context(), filter, pg)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, pagination.NewPagedResponse(customers, pg, total))
}

func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid customer id"))
		return
	}

	customer, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, apperror.NotFound("customer not found"))
		return
	}

	writeJSON(w, http.StatusOK, customer)
}

func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCustomerRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	customer, err := h.service.Create(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, customer)
}

func (h *CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid customer id"))
		return
	}

	var req models.UpdateCustomerRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	customer, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, customer)
}

func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid customer id"))
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CustomerHandler) GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := h.service.GetDistinctCountries(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, countries)
}

func (h *CustomerHandler) GetByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "country")
	customers, err := h.service.GetByCountry(r.Context(), country)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, customers)
}

func (h *CustomerHandler) GetTree(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")
	if mode == "" {
		mode = "by_country"
	}

	tree, err := h.service.GetCustomerTree(r.Context(), mode)
	if err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}
	writeJSON(w, http.StatusOK, tree)
}

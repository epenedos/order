package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/summit/summit-api/internal/service"
	"github.com/summit/summit-api/pkg/apperror"
	"github.com/summit/summit-api/pkg/pagination"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	pg := pagination.FromRequest(r)

	products, total, err := h.service.List(r.Context(), search, pg)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, pagination.NewPagedResponse(products, pg, total))
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid product id"))
		return
	}

	product, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, apperror.NotFound("product not found"))
		return
	}

	writeJSON(w, http.StatusOK, product)
}

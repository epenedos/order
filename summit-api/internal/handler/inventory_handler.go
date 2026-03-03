package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/summit/summit-api/internal/service"
	"github.com/summit/summit-api/pkg/apperror"
)

type InventoryHandler struct {
	service *service.InventoryService
}

func NewInventoryHandler(s *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: s}
}

func (h *InventoryHandler) GetByProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid product id"))
		return
	}

	inventory, err := h.service.GetByProduct(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, inventory)
}

func (h *InventoryHandler) List(w http.ResponseWriter, r *http.Request) {
	inventory, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, inventory)
}

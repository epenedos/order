package handler

import (
	"net/http"

	"github.com/summit/summit-api/internal/service"
)

type WarehouseHandler struct {
	service *service.WarehouseService
}

func NewWarehouseHandler(s *service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service: s}
}

func (h *WarehouseHandler) List(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, warehouses)
}

package handler

import (
	"net/http"

	"github.com/summit/summit-api/internal/service"
)

type RegionHandler struct {
	service *service.RegionService
}

func NewRegionHandler(s *service.RegionService) *RegionHandler {
	return &RegionHandler{service: s}
}

func (h *RegionHandler) List(w http.ResponseWriter, r *http.Request) {
	regions, err := h.service.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, regions)
}

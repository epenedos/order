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

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	var customerID *int
	if cid := r.URL.Query().Get("customer_id"); cid != "" {
		id, _ := strconv.Atoi(cid)
		customerID = &id
	}

	pg := pagination.FromRequest(r)
	orders, total, err := h.service.List(r.Context(), customerID, pg)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, pagination.NewPagedResponse(orders, pg, total))
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid order id"))
		return
	}

	order, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, apperror.NotFound("order not found"))
		return
	}

	writeJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateOrderRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	order, err := h.service.CreateOrder(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, order)
}

func (h *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid order id"))
		return
	}

	var req models.UpdateOrderRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	order, err := h.service.UpdateOrder(r.Context(), id, req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid order id"))
		return
	}

	if err := h.service.DeleteOrder(r.Context(), id); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Order Items

func (h *OrderHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid order id"))
		return
	}

	items, err := h.service.GetItems(r.Context(), orderID)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, items)
}

func (h *OrderHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, apperror.BadRequest("invalid order id"))
		return
	}

	var req models.CreateOrderItemRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	item, err := h.service.AddItem(r.Context(), orderID, req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, item)
}

func (h *OrderHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	orderID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	itemID, _ := strconv.Atoi(chi.URLParam(r, "itemId"))

	var req models.UpdateOrderItemRequest
	if err := validator.DecodeAndValidate(r, &req); err != nil {
		writeError(w, apperror.BadRequest(err.Error()))
		return
	}

	if err := h.service.UpdateItem(r.Context(), orderID, itemID, req); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *OrderHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	orderID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	itemID, _ := strconv.Atoi(chi.URLParam(r, "itemId"))

	if err := h.service.DeleteItem(r.Context(), orderID, itemID); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/service"
	"github.com/rseigha/goecomapi/pkg/response"
)

type OrderHandler struct {
	svc service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{svc: s}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var o domain.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: "invalid request"})
		return
	}
	// Extract user id from context (set by auth middleware)
	uid, ok := ctx.Value("user_id").(string)
	if !ok || uid == "" {
		response.JSON(w, http.StatusUnauthorized, response.APIResponse{Status: "error", Error: "unauthorized"})
		return
	}
	o.UserID = uid
	if err := h.svc.CreateOrder(ctx, &o); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusCreated, response.APIResponse{Status: "success", Data: o})
}

func (h *OrderHandler) ListByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, ok := ctx.Value("user_id").(string)
	if !ok || uid == "" {
		response.JSON(w, http.StatusUnauthorized, response.APIResponse{Status: "error", Error: "unauthorized"})
		return
	}
	orders, err := h.svc.GetByUser(ctx, uid)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success", Data: orders})
}
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/service"
	"github.com/rseigha/goecomapi/pkg/response"
)

type ProductHandler struct {
	svc service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{svc: s}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p domain.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: "invalid request"})
		return
	}
	if err := h.svc.Create(ctx, &p); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusCreated, response.APIResponse{Status: "success", Data: p})
}

func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	p, err := h.svc.GetByID(ctx, id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, response.APIResponse{Status: "error", Error: "product not found"})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success", Data: p})
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	page, _ := strconv.Atoi(q.Get("page"))
	products, total, err := h.svc.List(ctx, limit, page)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success", Data: map[string]interface{}{
		"items": products, "total": total, "page": page, "limit": limit,
	}})
}
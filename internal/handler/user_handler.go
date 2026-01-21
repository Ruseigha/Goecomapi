package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rseigha/goecomapi/internal/domain"
	"github.com/rseigha/goecomapi/internal/service"
	"github.com/rseigha/goecomapi/pkg/response"
)

type UserHandler struct {
	svc service.UserService
}



func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	u, err := h.svc.GetByID(ctx, id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, response.APIResponse{Status: "error", Error: "user not found"})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success", Data: u})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: "invalid request"})
		return
	}
	u.ID = id
	if err := h.svc.Update(ctx, &u); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success"})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	if err := h.svc.Delete(ctx, id); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success"})
}
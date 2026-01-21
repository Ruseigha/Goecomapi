package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rseigha/goecomapi/internal/service"
	"github.com/rseigha/goecomapi/pkg/response"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService service.AuthService
	logger *zap.Logger
}


func NewAuthHandler(as service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: as,
		logger: logger,
	}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: "invalid request"})
		return
	}
	user, err := h.authService.Register(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: err.Error()})
		return
	}
	response.JSON(w, http.StatusCreated, response.APIResponse{Status: "success", Data: user})
}


func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Status: "error", Error: "invalid request"})
		return
	}
	token, err := h.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, response.APIResponse{Status: "error", Error: "invalid credentials"})
		return
	}
	response.JSON(w, http.StatusOK, response.APIResponse{Status: "success", Data: map[string]string{"access_token": token}})
}
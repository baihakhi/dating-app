package handler

import (
	"github.com/baihakhi/dating-app/internal/services"
)

type Handler struct {
	service services.Services
}

func InitiHandler(services services.Services) *Handler {
	return &Handler{
		service: services,
	}
}

package handler

import (
	"my_app_name/internal/store/repository"

	"github.com/socle-lab/pkg/http/handler"
)

// Handlers is the type for handlers, and gives access to Socle and models
type Handler struct {
	handler.Handler
	Store repository.Store
}

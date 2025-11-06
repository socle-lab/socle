package handler

import (
	"my_app_name/internal/store"

	socle "github.com/socle-lab/core"
)

// Handlers is the type for handlers, and gives access to Socle and models
type Handler struct {
	Core  *socle.Socle
	Store store.Store
}

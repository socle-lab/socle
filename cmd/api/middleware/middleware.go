package middleware

import (
	"my_app_name/internal/store/repository"

	"github.com/socle-lab/pkg/http/middleware"
)

type Middleware struct {
	middleware.Middleware
	Store repository.Store
}

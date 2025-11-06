package middleware

import (
	"my_app_name/internal/store"

	socle "github.com/socle-lab/core"
)

type Middleware struct {
	Core  *socle.Socle
	Store store.Store
}

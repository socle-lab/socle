package handler

import (
	"context"
	"net/http"

	socle "github.com/socle-lab/core"
	"github.com/socle-lab/render"
)

func (h *Handler) render(w http.ResponseWriter, r *http.Request, opts render.PageOptions) error {
	return h.Core.Render.Page(w, r, opts)
}

func (h *Handler) log(tag string, args ...any) {
	switch tag {
	case "info":
		h.Core.Log.InfoLog.Println(args...)
	default:
		h.Core.Log.ErrorLog.Println(args...)
	}
}

func (h *Handler) putSession(ctx context.Context, key string, val interface{}) {
	h.Core.Log.InfoLog.Printf("Putting session data: %v", val)
	h.Core.Session.Put(ctx, key, val)
}

func (h *Handler) hasSession(ctx context.Context, key string) bool {
	return h.Core.Session.Exists(ctx, key)
}

func (h *Handler) getSession(ctx context.Context, key string) interface{} {
	return h.Core.Session.Get(ctx, key)
}

func (h *Handler) getSessionString(ctx context.Context, key string) string {
	return h.Core.Session.GetString(ctx, key)
}

func (h *Handler) removeSession(ctx context.Context, key string) {
	h.Core.Session.Remove(ctx, key)
}

func (h *Handler) renewSession(ctx context.Context) error {
	return h.Core.Session.RenewToken(ctx)
}

func (h *Handler) destroySession(ctx context.Context) error {
	return h.Core.Session.Destroy(ctx)
}

func (h *Handler) randomString(n int) string {
	return h.Core.RandomString(n)
}

func (h *Handler) encrypt(text string) (string, error) {
	enc := socle.Encryption{Key: []byte(h.Core.EncryptionKey)}

	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (h *Handler) decrypt(crypto string) (string, error) {
	enc := socle.Encryption{Key: []byte(h.Core.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}

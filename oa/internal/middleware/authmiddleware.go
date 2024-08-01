package middleware

import (
	"context"
	"net/http"
)

type AuthMiddleware struct {
	ID string `json:"id"`
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		// val := r.Header.Get("User-Agent")
		ctx := context.WithValue(r.Context(), "u", "123")
		next(w, r.WithContext(ctx))
	}
}

package middlewares

import (
	"context"
	"go-edash/config"
	"net/http"
	"strings"
)

// VerifyTokenMiddleware is a middleware function that verifies the JWT token in the request header.
// If the token is valid, it allows the request to proceed to the next handler.
// If the token is invalid or missing, it returns a 401 Unauthorized response.
func VerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request header
		authorization := r.Header.Get("Authorization")
		token := strings.Replace(authorization, "Bearer", "", -1)
		token = strings.TrimSpace(token)

		// Verify the token using the VerifyToken function from the libs package
		verify, err := config.VerifyToken(token)

		// If the token is invalid, return a 401 Unauthorized response
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		claims := verify.Claims
		ctx := context.WithValue(r.Context(), "claims", claims)

		// If the token is valid, allow the request to proceed to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//func (h *Handler) VerifyToken() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//	}
//}

package middlewares

import (
	"net/http"
	"strings"
)

// AuthorizationCheck is a middleware function that checks for a valid authorization token in the request header.
// If the token is missing or invalid, it returns a 401 Unauthorized or 400 Bad Request response respectively.
// Otherwise, it allows the request to proceed to the next handler in the chain.
func AuthorizationCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// Check if the request has a valid authorization token
		// If not, return a 401 Unauthorized response
		if authorization == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Check if the authorization token is of the "Bearer" type
		// If not, return a 400 Bad Request response
		if !strings.Contains(authorization, "Bearer") {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		// If the authorization token is valid, allow the request to proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

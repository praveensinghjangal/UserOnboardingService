package middlewares

import (
	"context"
	"net/http"
	"strings"

	"Onboarding_Service/services"
	"Onboarding_Service/utils"
)

func InternalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		email := claims.Email

		// Access the exported Mutex and TokenStore
		services.Mutex.Lock()
		latestToken, exists := services.TokenStore[email]
		services.Mutex.Unlock()

		if !exists || latestToken != token {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid or expired session")
			return
		}

		// Add email to request context
		ctx := context.WithValue(r.Context(), "email", email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

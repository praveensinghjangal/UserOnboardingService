package middlewares

import (
	"encoding/json"
	"net/http"

	"Onboarding_Service/utils"
)

func ValidateSignUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Decode and validate payload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := utils.ValidateEmail(payload.Email); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := utils.ValidatePassword(payload.Password); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		next.ServeHTTP(w, r) 
	})
}

func ValidateSignIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Decode and validate payload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := utils.ValidateEmail(payload.Email); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Onboarding_Service/services"
	"Onboarding_Service/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&payload)

	if err := services.SignUp(payload.Email, payload.Password); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&payload)

	token, err := services.SignIn(payload.Email, payload.Password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)
	fmt.Println(email)
	services.Logout(email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	token, err := services.RefreshToken(email)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"Onboarding_Service/controllers" 
    "Onboarding_Service/middlewares"  
)

func RegisterRoutes(router *mux.Router) {
	// Public routes
	router.Handle("/signup",
		middlewares.ChainMiddleware(middlewares.ValidateSignUp)(http.HandlerFunc(controllers.SignUp)),
	).Methods("POST")
	router.Handle("/signin",
		middlewares.ChainMiddleware(middlewares.ValidateSignIn)(http.HandlerFunc(controllers.SignIn)),
	).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middlewares.InternalAuth)

	protected.HandleFunc("/logout", controllers.Logout).Methods("POST")
	protected.HandleFunc("/refresh", controllers.Refresh).Methods("POST")
}
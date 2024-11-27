package services

import (
	"errors"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"Onboarding_Service/utils"
)

var (
	UserStore      = make(map[string]string)
	ActiveSessions = make(map[string]struct{})
	TokenStore     = make(map[string]string) // Exported for access in middleware
	Mutex          = sync.Mutex{}            // Exported for access in middleware
)

func SignUp(email, password string) error {
	Mutex.Lock()
	defer Mutex.Unlock()

	if _, exists := UserStore[email]; exists {
		return errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	UserStore[email] = string(hashedPassword)
	return nil
}

func SignIn(email, password string) (string, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	storedPassword, exists := UserStore[email]
	if !exists {
		return "", errors.New("invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate a new token
	token, err := utils.GenerateJWT(email, 24*time.Hour)
	if err != nil {
		return "", err
	}

	// Track the token
	ActiveSessions[email] = struct{}{}
	TokenStore[email] = token
	return token, nil
}

func Logout(email string) {
	Mutex.Lock()
	defer Mutex.Unlock()

	delete(ActiveSessions, email)
	delete(TokenStore, email)
}

func RefreshToken(email string) (string, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	if _, exists := ActiveSessions[email]; !exists {
		return "", errors.New("session not active")
	}

	// Generate a new token
	newToken, err := utils.GenerateJWT(email, 24*time.Hour)
	if err != nil {
		return "", err
	}

	// Update the tokenStore with the new token
	TokenStore[email] = newToken
	return newToken, nil
}
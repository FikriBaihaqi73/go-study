package main

import (
	"github.com/gorilla/mux"

	"github.com/FikriBaihaqi73/go-study/internal/user"
)

// registerUserRoutes mendaftarkan semua route untuk modul user
func registerUserRoutes(router *mux.Router) {
	// Initialize user dependencies
	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// User routes - akan jadi /api/users
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserById).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
}

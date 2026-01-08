package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/FikriBaihaqi73/go-study/internal/user"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	r.Use(LoggingMiddleware)

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	return r
}

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter membuat dan mengkonfigurasi HTTP router utama
func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Global middleware
	r.Use(LoggingMiddleware)

	// API subrouter dengan prefix /api
	api := r.PathPrefix("/api").Subrouter()

	// Register all module routes
	// Setiap modul punya file routes_*.go sendiri
	registerUserRoutes(api)
	// registerProductRoutes(api)  // ← Tinggal tambah di routes_product.go
	// registerOrderRoutes(api)    // ← Tinggal tambah di routes_order.go

	// Health check endpoint (di luar /api prefix)
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	return r
}

// healthCheckHandler adalah endpoint sederhana untuk cek status server
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

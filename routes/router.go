package routes

import (
	"github.com/FikriBaihaqi73/go-study/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Error examples routes
	errorHandler := &handlers.ErrorExamplesHandler{}
	r.HandleFunc("/examples", errorHandler.AllErrorExamples).Methods("GET")
	r.HandleFunc("/examples/400", errorHandler.Example400BadRequest).Methods("GET")
	r.HandleFunc("/examples/401", errorHandler.Example401Unauthorized).Methods("GET")
	r.HandleFunc("/examples/403", errorHandler.Example403Forbidden).Methods("GET")
	r.HandleFunc("/examples/404", errorHandler.Example404NotFound).Methods("GET")
	r.HandleFunc("/examples/409", errorHandler.Example409Conflict).Methods("GET")
	r.HandleFunc("/examples/422", errorHandler.Example422UnprocessableEntity).Methods("GET")
	r.HandleFunc("/examples/429", errorHandler.Example429TooManyRequests).Methods("GET")
	r.HandleFunc("/examples/500", errorHandler.Example500InternalServerError).Methods("GET")
	r.HandleFunc("/examples/503", errorHandler.Example503ServiceUnavailable).Methods("GET")

	return r
}

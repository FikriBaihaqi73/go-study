package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ErrorExamplesHandler demonstrates various HTTP error codes
type ErrorExamplesHandler struct{}

// Example400BadRequest godoc
// @Summary Example of 400 Bad Request error
// @Description Demonstrates when client sends invalid data
// @Tags error-examples
// @Produce json
// @Param invalid query string false "Send 'invalid' to trigger error"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Router /examples/400 [get]
func (h *ErrorExamplesHandler) Example400BadRequest(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("invalid")

	if param == "invalid" {
		sendJSONError(w, "Bad Request: The parameter 'invalid' cannot have value 'invalid'", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Request is valid",
		"status":  "success",
	})
}

// Example401Unauthorized godoc
// @Summary Example of 401 Unauthorized error
// @Description Demonstrates when authentication is required but not provided
// @Tags error-examples
// @Produce json
// @Param Authorization header string false "Bearer token"
// @Success 200 {object} map[string]string
// @Failure 401 {object} ErrorResponse
// @Router /examples/401 [get]
func (h *ErrorExamplesHandler) Example401Unauthorized(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		sendJSONError(w, "Unauthorized: Missing authentication token", http.StatusUnauthorized)
		return
	}

	if authHeader != "Bearer valid-token" {
		sendJSONError(w, "Unauthorized: Invalid authentication token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Authentication successful",
		"status":  "authenticated",
	})
}

// Example403Forbidden godoc
// @Summary Example of 403 Forbidden error
// @Description Demonstrates when user is authenticated but doesn't have permission
// @Tags error-examples
// @Produce json
// @Param role query string false "User role (admin or user)"
// @Success 200 {object} map[string]string
// @Failure 403 {object} ErrorResponse
// @Router /examples/403 [get]
func (h *ErrorExamplesHandler) Example403Forbidden(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")

	if role != "admin" {
		sendJSONError(w, "Forbidden: You don't have permission to access this resource. Admin role required.", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Access granted",
		"status":  "authorized",
	})
}

// Example404NotFound godoc
// @Summary Example of 404 Not Found error
// @Description Demonstrates when requested resource doesn't exist
// @Tags error-examples
// @Produce json
// @Param id query string false "Resource ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} ErrorResponse
// @Router /examples/404 [get]
func (h *ErrorExamplesHandler) Example404NotFound(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Simulate resource not found
	if id != "123" {
		sendJSONError(w, "Not Found: Resource with the specified ID does not exist", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Resource found",
		"id":      id,
	})
}

// Example409Conflict godoc
// @Summary Example of 409 Conflict error
// @Description Demonstrates when there's a conflict with current state
// @Tags error-examples
// @Produce json
// @Param email query string false "Email to check"
// @Success 200 {object} map[string]string
// @Failure 409 {object} ErrorResponse
// @Router /examples/409 [get]
func (h *ErrorExamplesHandler) Example409Conflict(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	// Simulate duplicate email
	existingEmails := []string{"john@example.com", "jane@example.com"}
	for _, existing := range existingEmails {
		if email == existing {
			sendJSONError(w, "Conflict: Email already exists in the system", http.StatusConflict)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Email is available",
		"email":   email,
	})
}

// Example422UnprocessableEntity godoc
// @Summary Example of 422 Unprocessable Entity error
// @Description Demonstrates when server understands the request but can't process it
// @Tags error-examples
// @Produce json
// @Param age query int false "User age"
// @Success 200 {object} map[string]string
// @Failure 422 {object} ErrorResponse
// @Router /examples/422 [get]
func (h *ErrorExamplesHandler) Example422UnprocessableEntity(w http.ResponseWriter, r *http.Request) {
	ageStr := r.URL.Query().Get("age")

	if ageStr == "" {
		sendJSONError(w, "Unprocessable Entity: Age is required", http.StatusUnprocessableEntity)
		return
	}

	// Simulate business logic validation
	age := 0
	_, err := fmt.Sscanf(ageStr, "%d", &age)
	if err != nil || age < 18 {
		sendJSONError(w, "Unprocessable Entity: Age must be 18 or older", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Age is valid",
		"age":     ageStr,
	})
}

// Example429TooManyRequests godoc
// @Summary Example of 429 Too Many Requests error
// @Description Demonstrates rate limiting
// @Tags error-examples
// @Produce json
// @Param requests query int false "Number of requests made"
// @Success 200 {object} map[string]string
// @Failure 429 {object} ErrorResponse
// @Router /examples/429 [get]
func (h *ErrorExamplesHandler) Example429TooManyRequests(w http.ResponseWriter, r *http.Request) {
	requestsStr := r.URL.Query().Get("requests")

	requests := 0
	if requestsStr != "" {
		fmt.Sscanf(requestsStr, "%d", &requests)
	}

	// Simulate rate limit (max 10 requests)
	if requests > 10 {
		w.Header().Set("Retry-After", "60")
		sendJSONError(w, "Too Many Requests: Rate limit exceeded. Please try again later.", http.StatusTooManyRequests)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Request processed",
		"requests": requestsStr,
	})
}

// Example500InternalServerError godoc
// @Summary Example of 500 Internal Server Error
// @Description Demonstrates unexpected server error
// @Tags error-examples
// @Produce json
// @Param trigger query string false "Set to 'error' to trigger 500"
// @Success 200 {object} map[string]string
// @Failure 500 {object} ErrorResponse
// @Router /examples/500 [get]
func (h *ErrorExamplesHandler) Example500InternalServerError(w http.ResponseWriter, r *http.Request) {
	trigger := r.URL.Query().Get("trigger")

	if trigger == "error" {
		// Simulate unexpected error
		sendJSONError(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Server is working properly",
		"status":  "healthy",
	})
}

// Example503ServiceUnavailable godoc
// @Summary Example of 503 Service Unavailable error
// @Description Demonstrates when service is temporarily unavailable
// @Tags error-examples
// @Produce json
// @Param maintenance query string false "Set to 'true' for maintenance mode"
// @Success 200 {object} map[string]string
// @Failure 503 {object} ErrorResponse
// @Router /examples/503 [get]
func (h *ErrorExamplesHandler) Example503ServiceUnavailable(w http.ResponseWriter, r *http.Request) {
	maintenance := r.URL.Query().Get("maintenance")

	if maintenance == "true" {
		w.Header().Set("Retry-After", "300") // 5 minutes
		sendJSONError(w, "Service Unavailable: System is under maintenance. Please try again later.", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Service is available",
		"uptime":  time.Now().Format(time.RFC3339),
	})
}

// AllErrorExamples godoc
// @Summary List all error code examples
// @Description Returns a list of all available error examples
// @Tags error-examples
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /examples [get]
func (h *ErrorExamplesHandler) AllErrorExamples(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	examples := map[string]interface{}{
		"available_examples": []map[string]string{
			{
				"code":        "400",
				"name":        "Bad Request",
				"endpoint":    "/examples/400?invalid=invalid",
				"description": "Client sent invalid data",
			},
			{
				"code":        "401",
				"name":        "Unauthorized",
				"endpoint":    "/examples/401",
				"description": "Authentication required",
			},
			{
				"code":        "403",
				"name":        "Forbidden",
				"endpoint":    "/examples/403?role=user",
				"description": "Insufficient permissions",
			},
			{
				"code":        "404",
				"name":        "Not Found",
				"endpoint":    "/examples/404?id=999",
				"description": "Resource not found",
			},
			{
				"code":        "409",
				"name":        "Conflict",
				"endpoint":    "/examples/409?email=john@example.com",
				"description": "Resource conflict",
			},
			{
				"code":        "422",
				"name":        "Unprocessable Entity",
				"endpoint":    "/examples/422?age=15",
				"description": "Validation failed",
			},
			{
				"code":        "429",
				"name":        "Too Many Requests",
				"endpoint":    "/examples/429?requests=15",
				"description": "Rate limit exceeded",
			},
			{
				"code":        "500",
				"name":        "Internal Server Error",
				"endpoint":    "/examples/500?trigger=error",
				"description": "Unexpected server error",
			},
			{
				"code":        "503",
				"name":        "Service Unavailable",
				"endpoint":    "/examples/503?maintenance=true",
				"description": "Service temporarily unavailable",
			},
		},
		"message": "Test these endpoints to see different error responses",
	}

	json.NewEncoder(w).Encode(examples)
}

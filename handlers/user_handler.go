package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/FikriBaihaqi73/go-study/models"
	"github.com/FikriBaihaqi73/go-study/storage"
	"github.com/gorilla/mux"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// sendJSONError sends a structured JSON error response
func sendJSONError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
		Code:    code,
	})
}

// validateUser validates user input data
func validateUser(user models.User) (bool, string) {
	// Check if name is empty
	if strings.TrimSpace(user.Name) == "" {
		return false, "Name is required and cannot be empty"
	}

	// Check if name is too short
	if len(strings.TrimSpace(user.Name)) < 3 {
		return false, "Name must be at least 3 characters long"
	}

	// Check if email is empty
	if strings.TrimSpace(user.Email) == "" {
		return false, "Email is required and cannot be empty"
	}

	// Simple email validation
	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return false, "Email format is invalid"
	}

	return true, ""
}

// GetUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate potential error (uncomment to test 500 error)
	// sendJSONError(w, "Internal server error while fetching users", http.StatusInternalServerError)
	// return

	err := json.NewEncoder(w).Encode(storage.Users)
	if err != nil {
		sendJSONError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /users/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Validate ID parameter
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		sendJSONError(w, "Invalid ID format. ID must be a number", http.StatusBadRequest)
		return
	}

	// Validate ID is positive
	if id <= 0 {
		sendJSONError(w, "Invalid ID. ID must be greater than 0", http.StatusBadRequest)
		return
	}

	for _, user := range storage.Users {
		if user.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	sendJSONError(w, "User not found", http.StatusNotFound)
}

// CreateUser godoc
// @Summary Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse "Bad Request - Invalid input"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		sendJSONError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate user data
	valid, message := validateUser(newUser)
	if !valid {
		sendJSONError(w, message, http.StatusBadRequest)
		return
	}

	// Check for duplicate email
	for _, user := range storage.Users {
		if strings.EqualFold(user.Email, newUser.Email) {
			sendJSONError(w, "Email already exists", http.StatusBadRequest)
			return
		}
	}

	storage.LastID++
	newUser.ID = storage.LastID
	storage.Users = append(storage.Users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated Data"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse "Bad Request - Invalid input"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Validate ID parameter
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		sendJSONError(w, "Invalid ID format. ID must be a number", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		sendJSONError(w, "Invalid ID. ID must be greater than 0", http.StatusBadRequest)
		return
	}

	// Decode update data
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		sendJSONError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate updated user data
	valid, message := validateUser(updatedUser)
	if !valid {
		sendJSONError(w, message, http.StatusBadRequest)
		return
	}

	// Check if user exists and update
	for i, user := range storage.Users {
		if user.ID == id {
			// Check for duplicate email (excluding current user)
			for j, u := range storage.Users {
				if i != j && strings.EqualFold(u.Email, updatedUser.Email) {
					sendJSONError(w, "Email already exists", http.StatusBadRequest)
					return
				}
			}

			storage.Users[i] = updatedUser
			storage.Users[i].ID = id
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(storage.Users[i])
			return
		}
	}

	sendJSONError(w, "User not found", http.StatusNotFound)
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse "Bad Request - Invalid ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Validate ID parameter
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		sendJSONError(w, "Invalid ID format. ID must be a number", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		sendJSONError(w, "Invalid ID. ID must be greater than 0", http.StatusBadRequest)
		return
	}

	for i, user := range storage.Users {
		if user.ID == id {
			storage.Users = append(storage.Users[:i], storage.Users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	sendJSONError(w, "User not found", http.StatusNotFound)
}

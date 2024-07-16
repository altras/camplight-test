package http

import (
	"backend/internal/application"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
    userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
    // Implement pagination logic
    users, err := h.userService.ListUsers(1, 10)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		JSONError(w, errors.NewAppError(http.StatusNotFound, "User not found", err), h.logger)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	users, err := h.userService.SearchUsers(query, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
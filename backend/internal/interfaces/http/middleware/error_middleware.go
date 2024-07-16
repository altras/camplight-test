package middleware

import (
	"backend/internal/logging/logging"
	"encoding/json"
	"net/http"
)

func ErrorMiddleware(logger *logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.ErrorLog.Printf("Panic: %v", err)
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func JSONError(w http.ResponseWriter, err error, logger *logging.Logger) {
	var appErr *errors.AppError
	if e, ok := err.(*errors.AppError); ok {
		appErr = e
	} else {
		appErr = errors.NewAppError(http.StatusInternalServerError, "Internal server error", err)
	}

	logger.ErrorLog.Printf("Error: %v", appErr)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Code)
	json.NewEncoder(w).Encode(map[string]string{"error": appErr.Message})
}
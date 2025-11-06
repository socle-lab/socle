package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelope{Error: message})
}

// func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(data)
// }

// func errorResponse(message string) string {
// 	return fmt.Sprintf(`{"success": false, "message": "%s"}`, message)
// }

// func successResponse(data interface{}) map[string]interface{} {
// 	return map[string]interface{}{
// 		"success": true,
// 		"message": "Request successful",
// 		"data":    data,
// 	}
// }

func successResponse(data interface{}) gin.H {
	return gin.H{"success": true, "message": "OK", "data": data}
}
func errorResponse(message string) gin.H {
	return gin.H{"success": false, "message": message, "data": nil}
}

// FailureResponsePayload defines the structure of the response payload.
type ResponsePayload struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// Payload error.
	Payload = "validation.payload"

	// Unauthorized error.
	Unauthorized = "unauthorized"

	// Expired error.
	Expired = "session_expired"

	// Addition error.
	Addition = "numbers-service.addition"

	// Application error.
	Application = "numbers-service.application"

	// NotImplemented error.
	NotImplemented = "not implemented"
	//MSGErrorList ..
	MSGErrorList = "something happened while retrieving the list"
	//MSGErrorOne ..
	MSGErrorOne = "something happened while retrieving the one"
	//MSGErrorCreate ..
	MSGErrorCreate = "something happened while reserving the number"
	//MSGErrorDelete ..
	MSGErrorDelete = "something happened while removing"
	//MSGSuccess ..
	MSGSuccess = "success"
	//MSGErrorUpdate ..
	MSGErrorUpdate = "something happened while updating"
)

func statusMap() map[string]int {
	m := map[string]int{
		NotImplemented: http.StatusNotImplemented,
		Payload:        http.StatusBadRequest,
		Unauthorized:   http.StatusUnauthorized,
		Expired:        http.StatusForbidden,
		Addition:       http.StatusServiceUnavailable,
		Application:    http.StatusInternalServerError,
	}
	return m
}

// Error REST type.
type Error struct {
	Code    string `json:"type"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// Fail convenience for REST error responses.
func Fail(c *gin.Context, code, message string, err error) {
	sc, ok := statusMap()[code]
	log.Printf("Error: %v ok: %v sc : %v msg %v", code, ok, sc, message)
	if !ok {
		sc = http.StatusInternalServerError
	}
	res := &Error{
		Code:    code,
		Reason:  err.Error(),
		Message: message,
	}
	c.JSON(sc, res)
}

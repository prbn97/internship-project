package utils

import (
	"api/main.go/models"
	"encoding/json"
	"net/http"
)

// 400
func BadRequest(res http.ResponseWriter, req *http.Request, msg string) {
	res.WriteHeader(http.StatusBadRequest)
	errorJson := models.TodoError{
		Error:   "bad request",
		Message: msg,
	}
	jsonBytes, err := json.Marshal(errorJson)
	if err != nil {
		return
	}
	res.Write(jsonBytes)
}

// 401
func Unauthorized(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusUnauthorized)
	res.Write([]byte(`{"error": "unauthorized"}`))
}

// 403
func Forbidden(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusForbidden)
	res.Write([]byte(`{"error": "forbidden"}`))
}

// 404
func NotFound(res http.ResponseWriter, req *http.Request, msg string) {
	res.WriteHeader(http.StatusNotFound)
	errorJson := models.TodoError{
		Error:   "Not Found",
		Message: msg,
	}
	jsonBytes, err := json.Marshal(errorJson)
	if err != nil {
		return
	}
	res.Write(jsonBytes)
}

// 405
func MethodNotAllowed(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusMethodNotAllowed)
	res.Write([]byte(`{"error": "method not allowed"}`))
}

// 409
func Conflict(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusConflict)
	res.Write([]byte(`{"error": "conflict"}`))
}

// 410
func Gone(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusGone)
	res.Write([]byte(`{"error": "gone"}`))
}

// 500
func InternalServerError(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(`{"error": "internal server error"}`))
}

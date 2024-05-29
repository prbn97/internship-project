package utils

import (
	"api/main.go/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func HandleJSONDecodeError(res http.ResponseWriter, req *http.Request, err error) {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		BadRequest(res, req, msg)
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := "Request body contains badly-formed JSON"
		BadRequest(res, req, msg)
	case errors.As(err, &unmarshalTypeError):
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		BadRequest(res, req, msg)
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
		BadRequest(res, req, msg)
	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		BadRequest(res, req, msg)
	default:
		BadRequest(res, req, "Invalid JSON")
	}
}

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

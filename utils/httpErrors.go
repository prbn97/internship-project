package utils

import "net/http"

// 400
func BadRequest(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(`{"error": "bad request"}`))
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
func NotFound(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`{"error": "not found"}`))
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

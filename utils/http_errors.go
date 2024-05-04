package utils

import "net/http"

func BadRequest(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(`{"error": "bad request"}`))
}

func NotFound(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`{"error": "not found"}`))
}

func InternalServerError(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`{"error": "internal server error"}`))
}

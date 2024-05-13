package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoHandler_List(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(NewTodoHandler().ServeHTTP)

	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	// Aqui você pode adicionar verificações adicionais no corpo da resposta, se necessário.
}

func TestTodoHandler_Create(t *testing.T) {
	// Criar uma solicitação simulada com um corpo JSON para criar um item de Todo
	todo := []byte(`{"title":"Test Todo","description":"Test Description","completed":false}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(todo))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(NewTodoHandler().ServeHTTP)

	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Aqui você pode adicionar verificações adicionais no corpo da resposta, se necessário.
}

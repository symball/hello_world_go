package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	r := setupRouter()

	// Create a test request
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message": "hello world"}`, w.Body.String())
}

func TestTestEndpoint(t *testing.T) {
	r := setupRouter()

	// Create a test request
	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"secret": "AIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGewQe"}`, w.Body.String())
}

func TestTestCredentialLeakEndpoint(t *testing.T) {
	r := setupRouter()

	// Create a test request
	req, _ := http.NewRequest("GET", "/test-credential-leak", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"secret": "BIzaSyDaGmWKa4JsXZ-HjGw7ISLn_3namBGfuQe"}`, w.Body.String())
}

package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func Request(method, path string, body interface{}) *httptest.ResponseRecorder {
	value, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, path, bytes.NewReader(value))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w
}

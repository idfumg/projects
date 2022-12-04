package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	tests := []struct {
		name string
		body string
		Code int
	}{
		{"valid user", `{"email":"admin@example.com","password":"secret"}`, http.StatusOK},
		{"not json", `I'm not a json`, http.StatusBadRequest},
		{"empty body", ``, http.StatusBadRequest},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email":""}`, http.StatusUnauthorized},
		{"empty password", `{"email":"admin@example.com"}`, http.StatusUnauthorized},
		{"wrong email", `{"email":"admin_wrong@example.com"}`, http.StatusUnauthorized},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.body)
			req := httptest.NewRequest("POST", "/auth", reader)
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(app.authenticate)

			handler.ServeHTTP(res, req)

			if test.Code != res.Code {
				t.Errorf("Wrong status code, want %d, got %d", test.Code, res.Code)
			}
		})
	}
}

package main

import (
	"context"
	"myapp/pkg/data"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_addIpToContext(t *testing.T) {
	tests := []struct {
		name        string
		headerName  string
		headerValue string
		addr        string
		isEmptyAddr bool
		want        string
	}{
		{"WillBeDefaultIP", "", "", "", false, ""},
		{"IPWillNotBeFound", "", "", "", true, "unknown"},
		{"IPFromHeader", "X-Forwarded-For", "192.3.2.1", "", false, "192.3.2.1"},
		{"NormalIPInRemoteAddr", "", "", "127.0.0.1:8080", false, "127.0.0.1"},
		{"WrongIP", "", "", "127.0.0.1?g", false, "unknown"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				val := r.Context().Value(contextUserKey)
				if val == nil {
					t.Error(contextUserKey, "value is not present")
				}

				ip, ok := val.(string)
				if !ok {
					t.Error("context value is not a string")
				}

				if len(ip) == 0 {
					t.Error("ip is empty")
				}

				if test.want != "" && !strings.EqualFold(ip, test.want) {
					t.Errorf("wrong ip. got %v(%d), want %v(%d)", ip, len(ip), test.want, len(test.want))
				}
			})

			handler := app.addIpToContext(next)
			req := httptest.NewRequest("GET", "http://example.com", nil)

			if test.isEmptyAddr {
				req.RemoteAddr = ""
			}

			if len(test.headerName) > 0 {
				req.Header.Add(test.headerName, test.headerValue)
			}

			if len(test.addr) > 0 {
				req.RemoteAddr = test.addr
			}

			handler.ServeHTTP(httptest.NewRecorder(), req)
		})
	}
}

func Test_app_ipFromContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), contextUserKey, "192.3.2.1")
	ip := app.ipFromContext(ctx)
	if len(ip) == 0 {
		t.Errorf("ip is empty: %v", ip)
	}
}

func Test_app_auth(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	tests := []struct {
		name   string
		isAuth bool
	}{
		{"logged in", true},
		{"not logged in", false},
	}

	for _, test := range tests {
		handlerToTest := app.auth(next)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if test.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		res := httptest.NewRecorder()
		handlerToTest.ServeHTTP(res, req)

		if test.isAuth && res.Code != http.StatusOK {
			t.Errorf("Wrong code; want %v, got %v", http.StatusOK, res.Code)
		}

		if !test.isAuth && res.Code != http.StatusTemporaryRedirect {
			t.Errorf("Wrong status; want %v, got %v", http.StatusTemporaryRedirect, res.Code)
		}
	}
}

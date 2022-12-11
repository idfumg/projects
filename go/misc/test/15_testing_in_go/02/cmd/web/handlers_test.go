package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_app_handlers(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		wantStatusCode int
	}{
		{"Home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	routes := app.routes()

	server := httptest.NewServer(routes)
	defer server.Close()

	pathToTemplates = "./../../templates/"

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := server.Client().Get(server.URL + test.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != test.wantStatusCode {
				t.Errorf("%s: error status code. got %v, want %v", test.name, resp.StatusCode, test.wantStatusCode)
			}
		})
	}
}

func TestAppHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Home)
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("TestAppHome returned wrong status code; want: 200, got: %v", res.Code)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))
	return req.WithContext(ctx)
}

func Test_app_Login(t *testing.T) {
	tests := []struct {
		name               string
		postedData         url.Values
		expectedStatusCode int
		expectedLoc        string
	}{
		{
			name: "Valid Login",
			postedData: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secret"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/user/profile",
		},
		{
			name: "Missing form data",
			postedData: url.Values{
				"email":    {""},
				"password": {""},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
		{
			name: "Invalid credentials",
			postedData: url.Values{
				"email":    {"you@example.com"},
				"password": {"secret"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/login", strings.NewReader(test.postedData.Encode()))
			req = addContextAndSessionToRequest(req, app)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(app.Login)
			handler.ServeHTTP(res, req)

			if res.Code != test.expectedStatusCode {
				t.Errorf("Wrong status; want: %v, got: %v", test.expectedStatusCode, res.Code)
			}

			actualLoc, err := res.Result().Location()
			if err != nil {
				t.Errorf("No location header set")
			}
			if actualLoc.String() != test.expectedLoc {
				t.Errorf("Wrong location; want: %v, got: %v", test.expectedLoc, actualLoc)
			}
		})
	}
}

package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_app_routes(t *testing.T) {
	tests := []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/user/profile", "GET"},
		{"/static/*", "GET"},
	}

	mux := app.routes()
	chiRoutes := mux.(chi.Routes)
	for _, test := range tests {
		if !isRouteExist(test.route, test.method, chiRoutes) {
			t.Errorf("route %s is not registered", test.route)
		}
	}
}

func isRouteExist(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}

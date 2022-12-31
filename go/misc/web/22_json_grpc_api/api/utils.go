package api

import (
	"context"
	"encoding/json"
	"math/rand"
	"myapp/types"
	"net/http"
)

type ApiFunc func(context.Context, http.ResponseWriter, *http.Request) error

func toHandlerFunc(fn ApiFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, types.RequestIdKey("requestID"), rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(ctx, w, r); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

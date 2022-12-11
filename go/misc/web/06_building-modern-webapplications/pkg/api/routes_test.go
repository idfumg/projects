package api

import (
	"myapp/pkg/config"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	appConfig, err := config.New(config.Test)
	if err != nil {
		t.Error(err)
	}

	mux := Routes(appConfig)

	switch v := mux.(type) {
	case *chi.Mux:
		break
	default:
		t.Errorf("Expected type chi.Mux, but got %T", v)
	}
}

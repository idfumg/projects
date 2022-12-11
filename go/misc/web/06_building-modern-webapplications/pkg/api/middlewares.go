package api

import (
	"myapp/pkg/config"
	"net/http"

	"github.com/justinas/nosurf"
)

func LoadCSRF(appConfig *config.AppConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   appConfig.InProduction,
			SameSite: http.SameSiteLaxMode,
		})
		return csrfHandler
	}
}

func LoadSession(appConfig *config.AppConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return appConfig.Session.LoadAndSave(next)
	}
}

package rest

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func (s *Server) noSurfMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   !s.config.GetIsTest(),
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

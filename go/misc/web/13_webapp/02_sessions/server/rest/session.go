package rest

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func (s *Server) sessionLoadMiddleware(next http.Handler) http.Handler {
	return s.session.LoadAndSave(next)
}

func createSession(config Config) *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = !config.GetIsTest()
	return s
}

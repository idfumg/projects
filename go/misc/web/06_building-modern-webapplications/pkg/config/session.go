package config

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func NewSession(inProduction bool) (*scs.SessionManager, error) {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = inProduction
	return session, nil
}

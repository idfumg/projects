package rest

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func addCORS(s *Server) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(s)
}

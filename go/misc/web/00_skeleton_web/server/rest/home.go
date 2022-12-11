package rest

import "net/http"

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	err := s.cache.Render(w, "home.page.gohtml", &TemplateData{})
	if err != nil {
		makeTextResponse(w, http.StatusNotFound, "Page wasn't found")
		return
	}
}

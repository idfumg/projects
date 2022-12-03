package rest

import (
	"fmt"
	"net/http"
)

func (s *Server) about(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	err := s.cache.Render(w, "about.page.gohtml", &TemplateData{
		StringMap: stringMap,
	})

	if err != nil {
		makeTextResponse(w, http.StatusNotFound, fmt.Sprintf("Page wasn't found: %s", err))
		return
	}
}

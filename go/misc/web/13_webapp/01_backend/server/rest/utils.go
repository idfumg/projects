package rest

import (
	"encoding/json"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

func makeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func makeTextResponse(w http.ResponseWriter, status int, text string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)
	w.Write([]byte(text))
}


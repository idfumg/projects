package main

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type helloResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`          // do not output this
	Date    string `json:",omitempty"` // do not output if empty
	Id      int    `json:"id,string"`  // convert to a string and rename
}

type helloRequest struct {
	Name string `json:"name"`
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

type validationContextKey string

func (v validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)
	v.next.ServeHTTP(w, r)
}

type helloHandler struct {
}

func newHelloHandler() http.Handler {
	return helloHandler{}
}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloResponse{Message: "Hello " + name}
	json.NewEncoder(w).Encode(response)
}

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}

type GzipHandler struct {
	next http.Handler
}

func newGzipHandler(next http.Handler) *GzipHandler {
	return &GzipHandler{next: next}
}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")
	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate is not implemented")
	} else {
		h.servePlain(w, r)
	}
}

func (h *GzipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(GzipResponseWriter{gzw, w}, r)
}

func (h *GzipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}

func main() {
	port := 8080

	http.Handle("/hello", newValidationHandler(newGzipHandler(newHelloHandler())))
	catHandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", catHandler))

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

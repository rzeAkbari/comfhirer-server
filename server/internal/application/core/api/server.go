package api

import (
	"io"
	"net/http"
)

type Compiler interface {
	Compile(map[string]any) []byte
}

type Scraper interface {
	Scrape(file []byte) map[string]any
}

type Server struct {
	Compiler
	Scraper
}

func NewServer(c Compiler, s Scraper) *Server {
	return &Server{
		c,
		s,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, _ := r.FormFile("prescription")

	b, _ := io.ReadAll(file)
	flatFhir := s.Scraper.Scrape(b)
	fhirMarshal := s.Compiler.Compile(flatFhir)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(fhirMarshal)
}

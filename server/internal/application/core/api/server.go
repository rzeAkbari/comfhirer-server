package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Compiler interface {
	Compile(map[string]any) ([]byte, []error)
}

type Scraper interface {
	Scrape(file []byte) (map[string]any, error)
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
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(400)
		e, _ := json.Marshal(err)
		w.Write(e)
	}
	file, _, _ := r.FormFile("prescription")

	b, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(400)
		e, _ := json.Marshal(err)
		w.Write(e)
	}
	flatFhir, err := s.Scraper.Scrape(b)
	if err != nil {
		w.WriteHeader(400)
		e, _ := json.Marshal(err)
		w.Write(e)
	}
	fhirMarshal, errs := s.Compiler.Compile(flatFhir)
	if len(errs) > 0 {
		w.WriteHeader(400)
		e, _ := json.Marshal(errs)
		w.Write(e)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(fhirMarshal)
}

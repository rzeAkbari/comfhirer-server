package comfhirer_server

import (
	"encoding/json"
	"io"
	"net/http"
)

type Compiler interface {
	Compile([]byte) (Bundle, error)
}

type Scraper interface {
	Scrape(file []byte) ([]byte, error)
}

type Server struct {
	Compiler Compiler
	Scraper  Scraper
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, _ := io.ReadAll(r.Body)
	comfhirerKeys, _ := s.Scraper.Scrape(file)
	bundle, _ := s.Compiler.Compile(comfhirerKeys)

	result, _ := json.Marshal(bundle)

	w.Write(result)
}

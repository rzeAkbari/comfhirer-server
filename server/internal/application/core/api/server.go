package api

import (
	"io"
	"net/http"
)

type Server struct {
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseMultipartForm(32 << 20)
	file, _, _ := r.FormFile("prescription")

	b, _ := io.ReadAll(file)

	return b
}

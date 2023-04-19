package cmd

import (
	"encoding/json"
	"github.com/rzeAkbari/comfhirer-server/comfhirer"
	"github.com/rzeAkbari/comfhirer-server/internal/application/core/api"
	"log"
	"net/http"
)

type wiring struct{}

func (w wiring) Compile(m map[string]any) []byte {
	fhir := comfhirer.Run(m)

	return fhir
}

func (w wiring) Scrape(file []byte) map[string]any {
	s := api.Scrape(file)
	fhirFlat := map[string]any{}

	json.Unmarshal([]byte(s), &fhirFlat)

	return fhirFlat
}

func Run() {
	w := wiring{}
	s := api.NewServer(w, w)

	log.Fatal(http.ListenAndServe(":5500", s))
}

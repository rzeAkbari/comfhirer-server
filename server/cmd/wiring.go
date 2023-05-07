package cmd

import (
	"encoding/json"
	"github.com/rzeAkbari/comfhirer-server/comfhirer"
	"github.com/rzeAkbari/comfhirer-server/internal/application/core/api"
	"log"
	"net/http"
	"os"
)

type wiring struct{}

func (w wiring) Compile(m map[string]any) []byte {
	fhir := comfhirer.Run(m)

	return fhir
}

func (w wiring) Scrape(file []byte) map[string]any {
	s, _ := api.Scrape(file)
	fhirFlat := map[string]any{}

	json.Unmarshal([]byte(s), &fhirFlat)

	return fhirFlat
}

func Run() {
	w := wiring{}
	s := api.NewServer(w, w)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "5500"
	}

	log.Fatal(http.ListenAndServe(":"+port, s))
}

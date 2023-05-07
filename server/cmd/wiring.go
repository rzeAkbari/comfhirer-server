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

func (w wiring) Compile(m map[string]any) ([]byte, []error) {
	fhir, errs := comfhirer.Run(m)

	return fhir, errs
}

func (w wiring) Scrape(file []byte) (map[string]any, error) {
	s, err := api.Scrape(file)
	if err != nil {
		return nil, err
	}
	fhirFlat := map[string]any{}

	err = json.Unmarshal([]byte(s), &fhirFlat)
	if err != nil {
		return nil, err
	}

	return fhirFlat, nil
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

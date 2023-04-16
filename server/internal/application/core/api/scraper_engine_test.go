package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/internal/application/core/api"
	"os"
	"testing"
)

func TestScraper(t *testing.T) {
	t.Run("scrape file", func(t *testing.T) {
		file, _ := os.ReadFile("./prescription.pdf")
		got := api.Scrape(file)

		want := "{\"nome\": \"AKBARI RAZIEH\", \"aic_code\": \"034930014\", \"farmaco\": \"KESTINE\", \"principio_attivo\": \"EBASTINA\", \"Confezione_di_riferimento\": \" 10MG 30 UNITA' USO ORALE\"}"
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/internal/application/core/api"
	"os"
	"testing"
)

func TestScraper(t *testing.T) {
	t.Run("scrape file", func(t *testing.T) {
		file, _ := os.ReadFile("./prescription.pdf")
		got, _ := api.Scrape(file)

		want := "{\"Patient.name.[0].given.{0}\": \"AKBARI RAZIEH\", \"Medication.code.coding.[0].code\": \"034930014\", \"Medication.code.coding.[0].display\": \"KESTINE\", \"Medication.ingredient.[0].item.concept.coding.[0].display\": \"EBASTINA\", \"Medication.doseForm.coding.[0].display\": \" 10MG 30 UNITA' USO ORALE\"}"
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

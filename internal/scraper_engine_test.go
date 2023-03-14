package internal_test

import (
	"github.com/rzeAkbari/comfhirer-server/internal"
	"os"
	"testing"
)

func TestScraper(t *testing.T) {
	t.Run("scrape file", func(t *testing.T) {
		file, _ := os.ReadFile("./fixture/prescription.jpg")
		got := internal.Do(file)

		want := "Medication.code.coding.[0].system=\"http://hl7.org/fhir/sid/ndc\"\nMedication.code.coding.[0].code=\"0409-6531-02\"\nMedication.code.coding.[0].display=\"Vancomycin Hydrochloride (VANCOMYCIN HYDROCHLORIDE)\"\nMedication.status=\"active\"\nMedication.manufacturer.reference=\"#org4\"\nMedication.form.coding.[0].system=\"http://snomed.info/sct\"\nMedication.form.coding.[0].code=\"385219001\"\nMedication.form.coding.[0].display=\"Injection Solution (qualifier value)\"\nMedication.ingredient.[0].itemCodeableConcept.coding.[0].system=\"http://www.nlm.nih.gov/research/umls/rxnorm\"\nMedication.ingredient.[0].itemCodeableConcept.coding.[0].code=\"66955\"\nMedication.ingredient.[0].itemCodeableConcept.coding.[0].display=\"Vancomycin Hydrochloride\"\nMedication.ingredient.[0].isActive=true\nMedication.ingredient.[0].strength.numerator.value=500\nMedication.ingredient.[0].strength.numerator.system=\"http://unitsofmeasure.org\"\nMedication.ingredient.[0].strength.numerator.code=\"mg\"\nMedication.ingredient.[0].strength.denominator.value=10\nMedication.ingredient.[0].strength.denominator.system=\"http://unitsofmeasure.org\"\nMedication.ingredient.[0].strength.denominator.code=\"mL\"\nMedication.batch.lotNumber=\"9494788\"\nMedication.batch.expirationDate=\"2017-05-22"

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

package cmd

import (
	"fmt"
	"github.com/rzeAkbari/comfhirer-server/comfhirer"
)

func main() {
	m := map[string]any{
		"Patient.birthDate":               "20-12-1988",
		"Medication.code.coding.[0].code": "A09",
	}
	bbundle := comfhirer.Run(m)

	fmt.Print(string(bbundle))
}

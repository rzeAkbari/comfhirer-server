package main

import (
	"fmt"
	"github.com/rzeAkbari/comfhirer-server/comfhirer"
)

func main() {
	m := map[string]any{
		"Patient.birthDate": "20-12-1988",
	}
	bbundle := comfhirer.Run(m)

	fmt.Print(string(bbundle))
}

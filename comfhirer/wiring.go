package comfhirer

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/adapter"
	"log"
)

func Run(input map[string]any) ([]byte, []error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	w := adapter.NewComfhirer()

	return w.Comfhire(input)
}

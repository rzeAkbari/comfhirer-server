package comfhirer

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/adapter"
)

func Run(input map[string]any) ([]byte, []error) {
	w := adapter.NewComfhirer()

	return w.Comfhire(input)
}

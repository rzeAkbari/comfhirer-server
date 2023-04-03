package comfhirer

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/adapter"
)

func Run(input map[string]any) []byte {
	w := adapter.NewComfhirer()

	return w.Comfhire(input)
}

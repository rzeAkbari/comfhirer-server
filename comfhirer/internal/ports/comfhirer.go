package ports

type Comfhirer interface {
	Comfhire(map[string]any) []byte
}

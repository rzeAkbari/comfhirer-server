package adapter

import (
	"encoding/json"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
)

type Comfhire struct {
	api.Tokenizer
	api.Parser
	api.Traverser
}

func (c Comfhire) Comfhire(m map[string]any) []byte {
	var ast []domain.ASTNode

	for key, value := range m {
		lexemes := c.Tokenize(key, value)
		astNode := c.Parse(lexemes)
		ast = append(ast, astNode)
	}

	bundle := c.Travers(ast)

	r, _ := json.Marshal(bundle)

	return r
}

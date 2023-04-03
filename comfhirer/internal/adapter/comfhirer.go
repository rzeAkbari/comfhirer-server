package adapter

import (
	"encoding/json"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
)

type Comfhir struct {
	tokenizer api.Tokenizer
	parser    api.Parser
	traverser api.Traverser
}

func (c Comfhir) Comfhire(m map[string]any) []byte {
	var ast []domain.ASTNode

	for key, value := range m {
		lexemes := c.tokenizer.Tokenize(key, value)
		astNode := c.parser.Parse(lexemes)
		ast = append(ast, astNode)
	}

	bundle := c.traverser.Travers(ast)

	r, _ := json.Marshal(bundle)

	return r
}

func NewComfhirer() Comfhir {
	return Comfhir{
		tokenizer: api.Tokenizer{},
		parser:    api.Parser{},
		traverser: api.Traverser{},
	}
}

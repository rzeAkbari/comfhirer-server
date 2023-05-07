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

func (c Comfhir) Comfhire(m map[string]any) ([]byte, []error) {
	var ast []domain.ASTNode

	for key, value := range m {
		lexemes := c.tokenizer.Tokenize(key, value)
		astNode := c.parser.Parse(lexemes)
		ast = append(ast, astNode)
	}

	bundle, err := c.traverser.Travers(ast)
	if len(err) > 0 {
		return nil, err
	}

	r, _ := json.Marshal(bundle)

	return r, nil
}

func NewComfhirer() Comfhir {
	return Comfhir{
		tokenizer: api.Tokenizer{},
		parser:    api.Parser{},
		traverser: api.Traverser{},
	}
}

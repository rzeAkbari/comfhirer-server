package api

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"regexp"
	"strings"
)

func Tokenize(key string, value any) domain.Lexemes {
	var tokens []domain.FieldToken
	fields := strings.Split(key, ".")
	exp := domain.Lexemes{
		ResourceToken: fields[0],
		ValueToken:    value,
	}
	for i := 1; i < len(fields); i++ {
		token := domain.NewToken(tokenType(fields[i]), normalize(fields[i]))
		tokens = append(tokens, token)
	}
	exp.FieldToken = tokens

	return exp
}

func tokenType(field string) domain.FieldTokenType {
	if strings.Contains(field, "[") {
		return domain.ArrayObject
	}
	if strings.Contains(field, "{") {
		return domain.ArrayValue
	}
	return domain.SimpleField
}

func normalize(field string) string {
	r, _ := regexp.Compile("([^\\[\\]}{]+)")
	return r.FindString(field)
}

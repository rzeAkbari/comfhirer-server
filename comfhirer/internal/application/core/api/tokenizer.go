package api

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"regexp"
	"strings"
)

type Tokenizer struct {
}

func (t Tokenizer) Tokenize(key string, value any) domain.Lexemes {
	var tokens []domain.FieldToken
	fields := strings.Split(key, ".")

	exp := domain.Lexemes{
		ResourceToken: fields[0],
		Index:         getResourceIndex(fields),
		ValueToken:    value,
	}

	for i := 1; i < len(fields); i++ {
		if isIndex(fields[i]) {
			continue
		}
		token := domain.NewToken(tokenType(fields[i]), normalize(fields[i]))
		tokens = append(tokens, token)
	}
	exp.FieldToken = tokens

	return exp
}

func isIndex(field string) bool {
	return strings.Contains(field, "(")
}

func getResourceIndex(fields []string) string {
	if len(fields) > 1 && isIndex(fields[1]) {
		left := strings.Replace(fields[1], "(", "", 1)
		index := strings.Replace(left, ")", "", 1)

		return index
	}

	return "0"
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

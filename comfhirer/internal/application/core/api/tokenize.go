package api

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"regexp"
	"strings"
)

func Do(key string, value any) []domain.Token {
	var tokens []domain.Token
	fields := strings.Split(key, ".")
	tokens = append(tokens, domain.New(domain.Resource, fields[0]))

	for i := 1; i < len(fields); i++ {
		token := domain.New(tokenType(fields[i]), normalize(fields[i]))
		tokens = append(tokens, token)
	}

	tokens = append(tokens, domain.New(domain.Data, value))

	return tokens
}

func tokenType(field string) domain.TokenType {
	if strings.Contains(field, "[") {
		return domain.ArrayObject
	}
	if strings.Contains(field, "{") {
		return domain.ArrayValue
	}
	return domain.Field
}

func normalize(field string) string {
	r, _ := regexp.Compile("([^\\[\\]}{]+)")
	return r.FindString(field)
}

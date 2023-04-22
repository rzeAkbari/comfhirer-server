package api

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
)

type Parser struct {
}

func (p Parser) Parse(lexemes domain.Lexemes) domain.ASTNode {

	if (len(lexemes.FieldToken)) == 0 {
		return domain.NewASTNode(lexemes.ResourceToken, lexemes.ValueToken, domain.FhirField{}, lexemes.Index)
	}

	field := *fhirFields(lexemes.FieldToken)

	return domain.NewASTNode(lexemes.ResourceToken, lexemes.ValueToken, field, lexemes.Index)
}

func fhirFields(tokens []domain.FieldToken) *domain.FhirField {
	if len(tokens) == 0 {
		return nil
	}
	field := domain.FhirField{
		Name:            tokens[0].FieldTokenName,
		FieldParsedType: getFieldType(tokens[0].FieldTokenType),
		FhirField:       fhirFields(tokens[1:]),
	}

	return &field
}

func getFieldType(tokenType domain.FieldTokenType) domain.FieldParsedType {
	switch tokenType {
	case domain.SimpleField:
		return domain.SingleField
	case domain.ArrayObject:
		return domain.MultipleNestedField
	case domain.ArrayValue:
		return domain.MultipleValueField
	}

	return -1
}

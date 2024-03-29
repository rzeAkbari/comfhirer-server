package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserBehaviour(t *testing.T) {
	parser := api.Parser{}
	t.Run("parse token with resource", func(t *testing.T) {
		lexemes := domain.Lexemes{
			ResourceToken: "Patient",
			Index:         "0",
			ValueToken:    "",
		}

		got := parser.Parse(lexemes)
		want := domain.NewASTNode("Patient", "", domain.FhirField{}, "0")

		assert.Equal(t, want, got)
	})

	t.Run("parse token with one field", func(t *testing.T) {
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "birthDate"),
		}
		lexemes := domain.Lexemes{
			ResourceToken: "Patient",
			Index:         "0",
			FieldToken:    tokens,
			ValueToken:    "20-12-1988",
		}

		got := parser.Parse(lexemes)

		field := domain.FhirField{
			Name:            "birthDate",
			FieldParsedType: domain.SingleField,
		}
		want := domain.NewASTNode("Patient", "20-12-1988", field, "0")

		assert.Equal(t, want, got)
	})

	t.Run("parse token with comfhirer position", func(t *testing.T) {
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "telecom"),
			domain.NewToken(domain.ArrayObject, "0"),
			domain.NewToken(domain.SimpleField, "rank"),
		}
		lexemes := domain.Lexemes{
			ResourceToken: "Patient",
			Index:         "1",
			FieldToken:    tokens,
			ValueToken:    1,
		}

		got := parser.Parse(lexemes)

		field := domain.FhirField{
			Name:            "telecom",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "0",
				FieldParsedType: domain.MultipleNestedField,
				FhirField: &domain.FhirField{
					Name:            "rank",
					FieldParsedType: domain.SingleField,
				},
			},
		}
		want := domain.NewASTNode("Patient", 1, field, "1")

		assert.Equal(t, want, got)
	})

	t.Run("parse token with comfhirer value position field", func(t *testing.T) {
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "name"),
			domain.NewToken(domain.ArrayObject, "0"),
			domain.NewToken(domain.SimpleField, "given"),
			domain.NewToken(domain.ArrayValue, "0"),
		}
		lexemes := domain.Lexemes{
			ResourceToken: "Patient",
			Index:         "1",
			FieldToken:    tokens,
			ValueToken:    "Jane",
		}

		got := parser.Parse(lexemes)

		field := domain.FhirField{
			Name:            "name",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "0",
				FieldParsedType: domain.MultipleNestedField,
				FhirField: &domain.FhirField{
					Name:            "given",
					FieldParsedType: domain.SingleField,
					FhirField: &domain.FhirField{
						Name:            "0",
						FieldParsedType: domain.MultipleValueField,
					},
				},
			},
		}
		want := domain.NewASTNode("Patient", "Jane", field, "1")

		assert.Equal(t, want, got)

	})
}

package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizeBehaviour(t *testing.T) {

	tokenizer := api.Tokenizer{}

	t.Run("tokenize simple resource", func(t *testing.T) {
		got := tokenizer.Tokenize("Patient", "")
		want := domain.Lexemes{
			ResourceToken: "Patient",
			ValueToken:    "",
		}
		assert.Equal(t, want, got)
	})

	t.Run("tokenize string data", func(t *testing.T) {

		got := tokenizer.Tokenize("Patient.birthDate", "20-12-1988")
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "birthDate"),
		}
		want := domain.Lexemes{
			ResourceToken: "Patient",
			FieldToken:    tokens,
			ValueToken:    "20-12-1988",
		}

		assert.Equal(t, want, got)
	})

	t.Run("tokenize boolean data", func(t *testing.T) {
		got := tokenizer.Tokenize("Person.active", true)
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "active"),
		}
		want := domain.Lexemes{
			ResourceToken: "Person",
			FieldToken:    tokens,
			ValueToken:    true,
		}
		assert.Equal(t, want, got)
	})

	t.Run("tokenize comfhirer object position field", func(t *testing.T) {
		got := tokenizer.Tokenize("Patient.telecom.[0].rank", 1)
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "telecom"),
			domain.NewToken(domain.ArrayObject, "0"),
			domain.NewToken(domain.SimpleField, "rank"),
		}

		want := domain.Lexemes{
			ResourceToken: "Patient",
			FieldToken:    tokens,
			ValueToken:    1,
		}

		assert.Equal(t, want, got)
	})

	t.Run("tokenize single comfhirer value position field", func(t *testing.T) {
		got := tokenizer.Tokenize("Patient.name.[0].given.{1}", "Jane")
		tokens := []domain.FieldToken{
			domain.NewToken(domain.SimpleField, "name"),
			domain.NewToken(domain.ArrayObject, "0"),
			domain.NewToken(domain.SimpleField, "given"),
			domain.NewToken(domain.ArrayValue, "1"),
		}
		want := domain.Lexemes{
			ResourceToken: "Patient",
			FieldToken:    tokens,
			ValueToken:    "Jane",
		}

		assert.Equal(t, want, got)
	})
}

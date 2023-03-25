package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizeBehaviour(t *testing.T) {

	t.Run("tokenize string data", func(t *testing.T) {

		got := api.Do("Patient.birthDate", "20-12-1988")
		want := []domain.Token{
			domain.New(domain.Resource, "Patient"),
			domain.New(domain.Field, "birthDate"),
			domain.New(domain.Data, "20-12-1988"),
		}

		assert.Equal(t, want, got)
	})

	t.Run("tokenize boolean data", func(t *testing.T) {
		got := api.Do("Person.active", true)
		want := []domain.Token{
			domain.New(domain.Resource, "Person"),
			domain.New(domain.Field, "active"),
			domain.New(domain.Data, true),
		}

		assert.Equal(t, want, got)
	})

	t.Run("tokenize comfhirer object position field", func(t *testing.T) {
		got := api.Do("Patient.telecom.[0].rank", 1)
		want := []domain.Token{
			domain.New(domain.Resource, "Patient"),
			domain.New(domain.Field, "telecom"),
			domain.New(domain.ArrayObject, "0"),
			domain.New(domain.Field, "rank"),
			domain.New(domain.Data, 1),
		}

		assert.Equal(t, want, got)
	})

	t.Run("tokenize single comfhirer value position field", func(t *testing.T) {
		got := api.Do("Patient.name.[0].given.{1}", "Jane")
		want := []domain.Token{
			domain.New(domain.Resource, "Patient"),
			domain.New(domain.Field, "name"),
			domain.New(domain.ArrayObject, "0"),
			domain.New(domain.Field, "given"),
			domain.New(domain.ArrayValue, "1"),
			domain.New(domain.Data, "Jane"),
		}

		assert.Equal(t, want, got)
	})
}

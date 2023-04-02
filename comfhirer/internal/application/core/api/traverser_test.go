package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	fhir_r4 "github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/fhir/r4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraverserBehaviour(t *testing.T) {
	t.Run("traverse ast with a single field", func(t *testing.T) {
		field := domain.FhirField{
			Name:            "birthDate",
			FieldParsedType: domain.SingleField,
		}
		ast := []domain.ASTNode{domain.NewASTNode("Patient", "20-12-1988", field)}

		got := api.Travers(ast)

		want := fhir_r4.Bundle{
			ResourceType: "Bundle",
			Entry: []fhir_r4.BundleEntry{
				{
					Resource: fhir_r4.Patient{ResourceType: "Patient", BirthDate: "20-12-1988"},
				},
			},
		}

		assert.Equal(t, want, got)

	})

	t.Run("traverse ast with comfhirer position", func(t *testing.T) {
		field := domain.FhirField{
			Name:            "maritalStatus",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "coding",
				FieldParsedType: domain.SingleField,
				FhirField: &domain.FhirField{
					Name:            "0",
					FieldParsedType: domain.MultipleNestedField,
					FhirField: &domain.FhirField{
						Name:            "code",
						FieldParsedType: domain.SingleField,
					},
				},
			},
		}
		ast := []domain.ASTNode{domain.NewASTNode("Patient", "M", field)}

		got := api.Travers(ast)

		want := fhir_r4.Bundle{
			ResourceType: "Bundle",
			Entry: []fhir_r4.BundleEntry{
				{
					Resource: fhir_r4.Patient{
						ResourceType: "Patient",
						MaritalStatus: fhir_r4.CodeableConcept{
							Coding: []fhir_r4.Coding{
								{
									Code: "M",
								},
							},
						}},
				},
			},
		}

		assert.Equal(t, want, got)

	})

	t.Run("traverse ast with two nodes", func(t *testing.T) {
		birthDayField := domain.FhirField{
			Name:            "birthDate",
			FieldParsedType: domain.SingleField,
		}

		maritalStatusField := domain.FhirField{
			Name:            "maritalStatus",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "coding",
				FieldParsedType: domain.SingleField,
				FhirField: &domain.FhirField{
					Name:            "0",
					FieldParsedType: domain.MultipleNestedField,
					FhirField: &domain.FhirField{
						Name:            "code",
						FieldParsedType: domain.SingleField,
					},
				},
			},
		}
		ast := []domain.ASTNode{
			domain.NewASTNode("Patient", "20-12-1988", birthDayField),
			domain.NewASTNode("Patient", "M", maritalStatusField)}

		got := api.Travers(ast)

		want := fhir_r4.Bundle{
			ResourceType: "Bundle",
			Entry: []fhir_r4.BundleEntry{
				{
					Resource: fhir_r4.Patient{
						ResourceType: "Patient",
						BirthDate:    "20-12-1988",
						MaritalStatus: fhir_r4.CodeableConcept{
							Coding: []fhir_r4.Coding{
								{
									Code: "M",
								},
							},
						}},
				},
			},
		}

		assert.Equal(t, want, got)

	})
}

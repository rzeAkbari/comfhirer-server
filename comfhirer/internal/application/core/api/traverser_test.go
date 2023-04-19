package api_test

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/api"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraverserBehaviour(t *testing.T) {
	traveser := api.Traverser{}

	t.Run("traverse ast with a single field", func(t *testing.T) {
		field := domain.FhirField{
			Name:            "birthDate",
			FieldParsedType: domain.SingleField,
		}
		ast := []domain.ASTNode{domain.NewASTNode("Patient", "20-12-1988", field)}

		got := traveser.Travers(ast)

		want := domain.Bundle{
			ResourceType: "Bundle",
			Entry: []domain.BundleEntry{
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
						BirthDate:    "20-12-1988",
					},
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

		got := traveser.Travers(ast)

		want := domain.Bundle{
			ResourceType: "Bundle",
			Entry: []domain.BundleEntry{
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
						MaritalStatus: &domain.CodeableConcept{
							Coding: []domain.Coding{
								{
									Code: "M",
								},
							},
						},
					},
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

		got := traveser.Travers(ast)

		want := domain.Bundle{
			ResourceType: "Bundle",
			Entry: []domain.BundleEntry{
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
						BirthDate:    "20-12-1988",
						MaritalStatus: &domain.CodeableConcept{
							Coding: []domain.Coding{
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

	t.Run("traverse ast with comfhirer value position field", func(t *testing.T) {
		name := domain.FhirField{
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

		middleName := domain.FhirField{
			Name:            "name",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "0",
				FieldParsedType: domain.MultipleNestedField,
				FhirField: &domain.FhirField{
					Name:            "given",
					FieldParsedType: domain.SingleField,
					FhirField: &domain.FhirField{
						Name:            "1",
						FieldParsedType: domain.MultipleValueField,
					},
				},
			},
		}
		ast := []domain.ASTNode{
			domain.NewASTNode("Patient", "Jane", name),
			domain.NewASTNode("Patient", "Mary", middleName)}

		got := traveser.Travers(ast)

		want := domain.Bundle{
			ResourceType: "Bundle",
			Entry: []domain.BundleEntry{
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
						Name: []domain.HumanName{
							{
								Given: []string{"Jane", "Mary"},
							},
						},
					},
				},
			},
		}

		assert.Equal(t, want, got)

	})

	t.Run("traverse ast with different node name", func(t *testing.T) {
		name := domain.FhirField{
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

		med := domain.FhirField{
			Name:            "code",
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
			domain.NewASTNode("Patient", "Jane", name),
			domain.NewASTNode("Medication", "A09", med)}

		got := traveser.Travers(ast)

		want := domain.Bundle{
			ResourceType: "Bundle",
			Entry: []domain.BundleEntry{
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
						Name: []domain.HumanName{
							{
								Given: []string{"Jane"},
							},
						},
					},
				},
				{
					Resource: domain.Medication{
						ResourceType: "Medication",
						Code: &domain.CodeableConcept{
							Coding: []domain.Coding{
								{
									Code: "A09",
								},
							},
						},
					},
				},
			},
		}

		assert.Equal(t, want.Entry, got.Entry)
	})
}

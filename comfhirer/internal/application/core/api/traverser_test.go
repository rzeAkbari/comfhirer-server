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
		ast := []domain.ASTNode{domain.NewASTNode("Patient", "20-12-1988", field, "0")}

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
		ast := []domain.ASTNode{domain.NewASTNode("Patient", "M", field, "0")}

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

	t.Run("traverse ast with multiple comfhirer position ", func(t *testing.T) {
		identifierTwo := domain.FhirField{
			Name:            "identifier",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "1",
				FieldParsedType: domain.MultipleNestedField,
				FhirField: &domain.FhirField{
					Name:            "use",
					FieldParsedType: domain.SingleField,
				},
			},
		}

		identifierThree := domain.FhirField{
			Name:            "identifier",
			FieldParsedType: domain.SingleField,
			FhirField: &domain.FhirField{
				Name:            "2",
				FieldParsedType: domain.MultipleNestedField,
				FhirField: &domain.FhirField{
					Name:            "use",
					FieldParsedType: domain.SingleField,
				},
			},
		}

		ast := []domain.ASTNode{
			domain.NewASTNode("Patient", "usual", identifierTwo, "0"),
			domain.NewASTNode("Patient", "official", identifierThree, "0")}

		got := traveser.Travers(ast)

		want := domain.Bundle{
			ResourceType: "Bundle",
			Entry: []domain.BundleEntry{
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
						Identifier: []domain.Identifier{
							{},
							{
								Use: "usual",
							},
							{
								Use: "official",
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
			domain.NewASTNode("Patient", "20-12-1988", birthDayField, "0"),
			domain.NewASTNode("Patient", "M", maritalStatusField, "1")}

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
				{
					Resource: domain.Patient{
						ResourceType: "Patient",
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
			domain.NewASTNode("Patient", "Jane", name, "0"),
			domain.NewASTNode("Patient", "Mary", middleName, "0")}

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
			domain.NewASTNode("Patient", "Jane", name, "0"),
			domain.NewASTNode("Medication", "A09", med, "0")}

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

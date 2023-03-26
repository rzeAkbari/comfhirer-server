package domain

type NodeType int

const (
	FhirResource NodeType = iota
)

type ASTNode struct {
	NodeType  NodeType
	NodeName  string
	NodeValue any
	FhirField
}

type FieldParsedType int

const (
	SingleField FieldParsedType = iota
	MultipleNestedField
	MultipleValueField
)

type FhirField struct {
	Name string
	FieldParsedType
	*FhirField
}

func NewASTNode(name string, value any, field FhirField) ASTNode {
	return ASTNode{
		NodeType:  FhirResource,
		NodeValue: value,
		NodeName:  name,
		FhirField: field,
	}
}

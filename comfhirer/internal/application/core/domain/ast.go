package domain

type NodeType int

const (
	FhirResource NodeType = iota
)

type ASTNode struct {
	NodeType  NodeType
	NodeIndex string
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

func NewASTNode(name string, value any, field FhirField, index string) ASTNode {
	return ASTNode{
		NodeType:  FhirResource,
		NodeIndex: index,
		NodeValue: value,
		NodeName:  name,
		FhirField: field,
	}
}

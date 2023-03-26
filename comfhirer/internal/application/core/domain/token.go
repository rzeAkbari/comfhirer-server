package domain

type Lexemes struct {
	ResourceToken string
	FieldToken    []FieldToken
	ValueToken    any
}

type FieldToken struct {
	FieldTokenType FieldTokenType
	FieldTokenName string
}

type FieldTokenType int

const (
	SimpleField FieldTokenType = iota
	ArrayObject
	ArrayValue
)

func NewToken(t FieldTokenType, v string) FieldToken {
	return FieldToken{
		FieldTokenType: t,
		FieldTokenName: v,
	}
}

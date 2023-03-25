package domain

type Expression string

type Token struct {
	TokenType  TokenType
	TokenValue any
}

type TokenType int

const (
	Resource TokenType = iota
	Field
	Data
	ArrayObject
	ArrayValue
)

func New(t TokenType, v any) Token {
	return Token{
		TokenType:  t,
		TokenValue: v,
	}
}

package token

type Token struct {
	Typ tokenType
	Val any
}

func NewToken(t tokenType, input string, line int, pos int) Token {
	return Token{Typ: t, Val: input}
}

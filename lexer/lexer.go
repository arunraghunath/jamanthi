package lexer

import (
	"errors"
	"log"
	"strconv"

	"github.com/arunraghunath/jamanthi/token"
)

type Lexer struct {
	input   string
	line    int
	curpos  int
	curchar byte
	nextpos int
	eof     bool
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readNextChar()
	return lex
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipSpace()
	switch {
	case l.curchar == '+':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '+' {
				l.readNextChar()
				tok = token.Token{Typ: token.INCREMENT, Val: "++"}
			} else if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.PLUSASSIGN, Val: "+="}
			} else {
				tok = token.Token{Typ: token.PLUS, Val: "+"}
			}
		} else {
			tok = token.Token{Typ: token.PLUS, Val: "+"}
		}
		l.readNextChar()

	case l.curchar == '-':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '-' {
				l.readNextChar()
				tok = token.Token{Typ: token.DECREMENT, Val: "--"}
			} else if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.MINUSASSIGN, Val: "-="}
			} else {
				tok = token.Token{Typ: token.MINUS, Val: "+"}
			}
		} else {
			tok = token.Token{Typ: token.MINUS, Val: "+"}
		}
		l.readNextChar()

	case l.curchar == '*':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.ASTERIKASSIGN, Val: "*="}
			} else {
				tok = token.Token{Typ: token.ASTERIK, Val: "*"}
			}
		} else {
			tok = token.Token{Typ: token.ASTERIK, Val: "*"}
		}
		l.readNextChar()
	case l.curchar == '/':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.SLASHASSIGN, Val: "/="}
			} else {
				tok = token.Token{Typ: token.SLASH, Val: "/"}
			}
		} else {
			tok = token.Token{Typ: token.SLASH, Val: "/"}
		}
		l.readNextChar()
	case l.curchar == '%':
		tok = token.Token{Typ: token.SLASH, Val: "%"}
		l.readNextChar()
	case l.curchar == '(':
		tok = token.Token{Typ: token.LPAREN, Val: "("}
		l.readNextChar()
	case l.curchar == ')':
		tok = token.Token{Typ: token.RPAREN, Val: ")"}
		l.readNextChar()
	case l.curchar == '{':
		tok = token.Token{Typ: token.LBRACE, Val: "{"}
		l.readNextChar()
	case l.curchar == '}':
		tok = token.Token{Typ: token.RBRACE, Val: "}"}
		l.readNextChar()
	case l.curchar == '>':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.GTET, Val: ">="}
			} else {
				tok = token.Token{Typ: token.GT, Val: ">"}
			}
		} else {
			tok = token.Token{Typ: token.GT, Val: ">"}
		}
		l.readNextChar()
	case l.curchar == '<':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.LTET, Val: "<="}
			} else {
				tok = token.Token{Typ: token.LT, Val: "<"}
			}
		} else {
			tok = token.Token{Typ: token.LT, Val: "<"}
		}
		l.readNextChar()
	case l.curchar == '=':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.EQUALS, Val: "=="}
			} else {
				tok = token.Token{Typ: token.ASSIGN, Val: "="}
			}
		} else {
			tok = token.Token{Typ: token.ASSIGN, Val: "="}
		}
		l.readNextChar()
	case l.curchar == '!':
		ch, err := l.peekChar()
		if err == nil {
			if ch == '=' {
				l.readNextChar()
				tok = token.Token{Typ: token.NOTEQUALS, Val: "!="}
			} else {
				tok = token.Token{Typ: token.NOT, Val: "!"}
			}
		} else {
			tok = token.Token{Typ: token.ASSIGN, Val: "!"}
		}
		l.readNextChar()
	case isLetter(l.curchar):
		tok = l.getString()
	case isDigit(l.curchar):
		tok = l.getNumber()
	default:
		tok = token.Token{Typ: token.EOF, Val: ""}
	}
	//l.readNextChar()
	return tok
}

func (l *Lexer) readNextChar() error {
	if l.isEOF() {
		l.curchar = byte(token.EOF)
		l.curpos = l.nextpos
		return errors.New("EOF")
	} else {
		l.curchar = l.input[l.nextpos]
		l.curpos = l.nextpos
		l.nextpos++
		return nil
	}
}

func (l *Lexer) skipSpace() {
	if !l.isEOF() {
		for l.curchar == ' ' || l.curchar == '\t' || l.curchar == '\r' || l.curchar == '\n' {
			l.readNextChar()
		}

	}
}

func (l *Lexer) getString() token.Token {
	startpos := l.curpos
	endpos := startpos + 1

	for !isSpace(l.curchar) && isLetter(l.curchar) && l.curchar != byte(token.EOF) {
		l.readNextChar()
		endpos = l.curpos
	}
	return token.Token{Typ: token.STRING, Val: l.input[startpos:endpos]}
}

func (l *Lexer) getNumber() token.Token {
	startpos := l.curpos
	endpos := startpos + 1
	for !isSpace(l.curchar) && isDigit(l.curchar) && l.curchar != byte(token.EOF) {
		l.readNextChar()
		endpos = l.curpos
	}
	res, err := strconv.Atoi(l.input[startpos:endpos])
	if err != nil {
		log.Println(err)
	}
	return token.Token{Typ: token.INT, Val: res}
}

func (l *Lexer) peekChar() (byte, error) {
	if l.isEOF() {
		return '0', errors.New("EOF")
	}
	return l.input[l.nextpos], nil
}

func isLetter(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func isSpace(ch byte) bool {
	if ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' {
		return true
	}
	return false
}

func (l *Lexer) isEOF() bool {
	if l.nextpos >= len(l.input) {
		l.eof = true
		return true
	}
	return false
}

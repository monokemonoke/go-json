package myjson

import "log"

type Lexer struct {
	source string
	curpos int
}

func NewLexer(source string) *Lexer {
	var lexer Lexer
	lexer.source = source
	lexer.curpos = 0
	return &lexer
}

func (l *Lexer) curChar() (string, bool) {
	if l.curpos >= len(l.source) {
		return "", false
	}

	res := l.source[l.curpos]
	l.curpos++
	return string(res), true
}

func (l *Lexer) Lex() []TOKEN_TYPE {
	var tokens []TOKEN_TYPE

	char, ok := l.curChar()
	for ok {
		log.Println(char)
		char, ok = l.curChar()
	}

	tokens = []TOKEN_TYPE{
		TOKEN_LEFT_BRACKET,
		TOKEN_KEY,
		TOKEN_STRING,
		TOKEN_RIGHT_BRACKET,
	}

	return tokens
}

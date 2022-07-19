package myjson

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


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

func (*Lexer) Lex() []TOKEN_TYPE {
	return []TOKEN_TYPE{
		TOKEN_LEFT_BRACKET,
		TOKEN_KEY,
		TOKEN_STRING,
		TOKEN_RIGHT_BRACKET,
	}
}

package myjson

func Lexer(text string) []TOKEN_TYPE {
	return []TOKEN_TYPE{
		TOKEN_LEFT_BRACKET,
		TOKEN_KEY,
		TOKEN_STRING,
		TOKEN_RIGHT_BRACKET,
	}
}
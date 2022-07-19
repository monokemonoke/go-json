package myjson

import (
	"testing"
)

func TestLexer(t *testing.T) {
	testInput := `{ "key" : "value" }`
	expect := []TOKEN_TYPE{
		TOKEN_LEFT_BRACKET,
		TOKEN_KEY,
		TOKEN_STRING,
		TOKEN_RIGHT_BRACKET,
	}
	actual := Lexer(testInput)

	if len(actual) != len(expect) {
		t.Fatalf("Expect %d tokens but got %d tokens", len(expect), len(actual))
	}
	for i := 0; i < len(expect); i++ {
		if actual[i] != expect[i] {
			t.Errorf("Expect tokentype %d but got %d", expect[i], actual[i])
		}
	}
}

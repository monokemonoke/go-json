package main

import (
	"github.com/monokemonoke/go-json/myjson"
)

func main() {
	testInput := `{ "key" : "value" }`
	lexer := myjson.NewLexer(testInput)
	lexer.Lex()
}
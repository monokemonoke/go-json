package main

import (
	"fmt"

	"github.com/monokemonoke/go-json/myjson"
)

func main() {
	fmt.Println("Hello go")
	fmt.Println(myjson.Hoge())

	testInput := `{ "key" : "value" }`
	lexer := myjson.NewLexer(testInput)
	lexer.Lex()
}
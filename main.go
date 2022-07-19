package main

import (
	"fmt"

	"github.com/monokemonoke/go-json/myjson"
)

func main() {
	testInput := `{ "price" : "100" }`
	lexer := myjson.NewParser(testInput)
	json := lexer.Parse()
	for k, v := range json {
		fmt.Printf("%s : %s \n", k, v)
	}

	testInput2 := `{ "price" : "100", "apple" :"banana", "hello world": "hogehoge"}`
	lexer2 := myjson.NewParser(testInput2)
	json2 := lexer2.Parse()
	for k, v := range json2 {
		fmt.Printf("%s : %s \n", k, v)
	}
}

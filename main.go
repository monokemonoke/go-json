package main

import (
	"fmt"

	"github.com/monokemonoke/go-json/myjson"
)

func main() {
	testInput := `{ "price" : 100 }`
	parser := myjson.NewParser(testInput)
	json := parser.ParseJson()
	// lexer := myjson.NewParser(testInput)
	// json := lexer.Parse()
	for k, v := range json {
		fmt.Println(k, v)
	}

	testInput2 := `{ "price" : 100, "apple" :"banana", "hello world": true , "hoge": false}`
	parser2 := myjson.NewParser(testInput2)
	json2 := parser2.ParseJson()
	for k, v := range json2 {
		fmt.Printf("%s : %s \n", k, v)
	}
}

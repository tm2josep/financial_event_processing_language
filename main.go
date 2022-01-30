package main

import (
	"fepl/lexer"
	"fmt"
)

func main() {
	var lex = new(lexer.Lexer)
	source := "alloc @'field_name' 500 @'field_name_2';\n assess @'field_name'"
	for token := range lex.Stream(source) {
		fmt.Println(token.Name, token.Location)
		fmt.Printf("%#v\n", token.Content)
	}
}

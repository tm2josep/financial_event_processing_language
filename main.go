package main

import (
	"fepl/lexer"
	"fmt"
)

func main() {
	lex := new(lexer.Lexer)
	source := "(@'field' + 10)"
	tokens := make(chan lexer.Token)
	go lex.Stream(source, tokens)
	for token := range tokens {
		fmt.Println(token)
	}
}

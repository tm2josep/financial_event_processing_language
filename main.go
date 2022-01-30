package main

import (
	"fepl/lexer"
	"fepl/parser"
	"fmt"
)

func main() {
	lex := new(lexer.Lexer)
	source := "alloc @'name_one' (500 + 3) @'name_two'"
	tokens := make(chan lexer.Token)
	go lex.Stream(source, tokens)
	pars := new(parser.Parser)
	fmt.Println(pars.GetAst(tokens))
}

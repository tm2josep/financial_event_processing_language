package main

import (
	"fepl/behavior"
	"fepl/lexer"
	"fepl/parser"
	"fmt"
)

func main() {
	lex := new(lexer.Lexer)
	source := "(@'field' + 10)"
	tokens := make(chan lexer.Token)
	go lex.Stream(source, tokens)
	pars := new(parser.Parser)
	ast := pars.GetAst(tokens)
	fmt.Println(behavior.Build(ast))
}

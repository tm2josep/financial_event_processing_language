package main

import (
	"fepl/lexer"
	"fepl/parser"
	"fmt"
)

func main() {
	lex := new(lexer.Lexer)
	source := "alloc @'claim' 500 @'retained';"
	tokens := make(chan lexer.Token)
	go lex.Stream(source, tokens)
	rootNode := parser.Parse(tokens)
	nodes := make(chan parser.Node)
	go rootNode.Walk(nodes)
	for node := range nodes {
		fmt.Println(node.Kind)
	}
}

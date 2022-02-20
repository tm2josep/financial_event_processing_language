package main

import (
	"fepl/behavior"
	"fepl/lexer"
	"fepl/parser"
	"fmt"
	"math/rand"
)

func generateLosses(out chan map[string]interface{}) {
	defer close(out)
	for i := 0; i < 1000; i++ {
		loss := make(map[string]interface{})

		loss["claim"] = rand.Float64() * 1000
		loss["retained"] = 0.00

		out <- loss
	}
}

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

	allocation := behavior.Allocation{
		Source: behavior.Field{Name: "claim"},
		Value:  &behavior.TermValue{CurrentValue: 500},
		Target: behavior.Field{Name: "retained"},
	}

	losses := make(chan map[string]interface{})
	go generateLosses(losses)
	for loss := range losses {
		fmt.Println("Before -", "Claim:", loss["claim"], "Retained:", loss["retained"])
		newLoss, err := allocation.Apply(loss)
		if err != nil {
			panic(err)
		}
		fmt.Println("After  -", "Claim:", newLoss["claim"], "Retained:", newLoss["retained"])
	}
}

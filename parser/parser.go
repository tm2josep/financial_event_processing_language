package parser

import (
	"fepl/lexer"
	"fmt"
	"reflect"
)

type Parser struct {
	root   *Node
	tokens []lexer.Token
	rules  []rule
}

type rule struct {
	name        string
	expressions []string
}

type Node struct {
	Kind       string
	Token      lexer.Token
	ChildNodes []*Node
}

func setup() Parser {
	parser := new(Parser)
	parser.addRule(rule{"block", []string{"block", "statement"}})
	parser.addRule(rule{"statement", []string{"expression", "EOL"}})
	parser.addRule(rule{"expression", []string{"ALLOC", "field", "value", "field"}})
	parser.addRule(rule{"field", []string{"FIELD_START", "NAME", "FIELD_END"}})
	parser.addRule(rule{"value", []string{"NUMBER"}})
	return *parser
}

func (parser *Parser) addRule(rule rule) {
	parser.rules = append(parser.rules, rule)
}

func Parse(tokens chan lexer.Token) Node {
	parser := setup()
	for token := range tokens {
		parser.tokens = append(parser.tokens, token)
	}
	return parser.resolve()
}

func (rule *rule) matches(nodes []*Node) bool {
	nodeKinds := []string{}
	for _, node := range nodes {
		nodeKinds = append(nodeKinds, node.Kind)
	}
	// zip
	matches := reflect.DeepEqual(rule.expressions, nodeKinds)
	return matches
}

func printMatch(name string, nodeList []*Node) {
	fmt.Print(len(nodeList), " ")
	for _, n := range nodeList {
		fmt.Print(n.Kind + " ")
	}
	fmt.Print("---> " + name)
	fmt.Print("\n")
}

func (n *Node) descend(out chan Node) {
	for i := 0; i < len(n.ChildNodes); i++ {
		out <- *n.ChildNodes[i]
		n.ChildNodes[i].descend(out)
	}
}

func (n *Node) Walk(out chan Node) {
	n.descend(out)
	close(out)
}

func (parser *Parser) resolve() Node {
	nodes := []*Node{}
	for _, token := range parser.tokens {
		nodes = append(nodes, &Node{token.Kind, token, []*Node{}})
	}

mainloop:
	for len(nodes) > 1 {
		for i := 0; i < len(nodes); i++ {
			for _, rule := range parser.rules {
				for p := 0; p <= len(rule.expressions); p++ {
					j := i + p
					match := rule.matches(nodes[i:j])
					if !match {
						continue
					}

					children := nodes[i:j]

					printMatch(rule.name, children)

					node := &Node{rule.name, lexer.Token{}, children}
					temp := []*Node{}

					temp = append(temp, nodes[:i]...)
					temp = append(temp, node)
					temp = append(temp, nodes[j:]...)

					nodes = temp

					continue mainloop
				}
			}
		}
	}

	return *nodes[0]
}

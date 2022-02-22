package parser

import (
	"errors"
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

type NodeWalk struct {
	Node  Node
	Depth int
}

func setup() Parser {
	parser := new(Parser)
	parser.addRule(rule{"block", []string{"block", "statement"}})
	parser.addRule(rule{"block", []string{"statement"}})
	parser.addRule(rule{"statement", []string{"AGGREGATE", "field", "agg_fields", "EOL"}})
	parser.addRule(rule{"statement", []string{"ALLOC", "field", "value", "field", "EOL"}})
	parser.addRule(rule{"field", []string{"FIELD_START", "NAME", "FIELD_END"}})
	parser.addRule(rule{"agg_field", []string{"field", "AGG_MODE"}})
	parser.addRule(rule{"agg_fields", []string{"agg_fields", "agg_field"}})
	parser.addRule(rule{"agg_fields", []string{"agg_field"}})
	parser.addRule(rule{"value", []string{"NUMBER"}})
	return *parser
}

func (parser *Parser) addRule(rule rule) {
	parser.rules = append(parser.rules, rule)
}

func Parse(tokens chan lexer.Token) (Node, error) {
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

func (n *Node) descend(out chan NodeWalk, depth int) {
	for i := 0; i < len(n.ChildNodes); i++ {
		result := NodeWalk{*n.ChildNodes[i], depth}
		out <- result
		n.ChildNodes[i].descend(out, depth+1)
	}
}

func (n *Node) Walk(out chan NodeWalk) {
	root := NodeWalk{*n, 0}
	out <- root
	n.descend(out, 1)
	close(out)
}

func (parser *Parser) collapse(nodes []*Node) ([]*Node, error) {
	for i := 0; i < len(nodes); i++ {
		for _, rule := range parser.rules {
			for p := 0; p <= len(rule.expressions); p++ {
				j := i + p
				if j > len(nodes) {
					continue
				}

				match := rule.matches(nodes[i:j])
				if !match {
					continue
				}

				children := nodes[i:j]

				printMatch("", nodes)
				printMatch(rule.name, children)

				node := &Node{rule.name, lexer.Token{}, children}
				temp := []*Node{}

				temp = append(temp, nodes[:i]...)
				temp = append(temp, node)
				temp = append(temp, nodes[j:]...)

				return temp, nil
			}
		}
	}

	return []*Node{}, errors.New("parser could not resolve")
}

func (parser *Parser) resolve() (Node, error) {
	nodes := []*Node{}
	for _, token := range parser.tokens {
		nodes = append(nodes, &Node{token.Kind, token, []*Node{}})
	}

	for len(nodes) > 1 {
		newNodes, err := parser.collapse(nodes)
		if err != nil {
			return Node{}, err
		}
		nodes = newNodes
	}

	return *nodes[0], nil
}

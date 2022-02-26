package parser

import (
	"errors"
	"fepl/lexer"
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

func setup(tokens chan lexer.Token) Parser {
	parser := new(Parser)
	for _, rule := range rules {
		parser.addRule(rule)
	}
	for token := range tokens {
		parser.tokens = append(parser.tokens, token)
	}
	return *parser
}

func (parser *Parser) addRule(rule rule) {
	parser.rules = append(parser.rules, rule)
}

func Parse(tokens chan lexer.Token) (Node, error) {
	parser := setup(tokens)
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

func (rule *rule) collapse(children []*Node) (*Node, error) {
	match := rule.matches(children)
	if !match {
		err := errors.New("rule doesn't match")
		return &Node{}, err
	}

	// printMatch(rule.name, children)

	node := &Node{rule.name, lexer.Token{}, children}

	return node, nil
}

func (parser *Parser) collapse(nodes []*Node) ([]*Node, error) {
	for i := 0; i < len(nodes); i++ {
		for _, rule := range parser.rules {
			for p := 0; p <= len(rule.expressions); p++ {
				j := i + p
				if j > len(nodes) {
					continue
				}

				node, err := rule.collapse(nodes[i:j])

				if err != nil {
					continue
				}

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

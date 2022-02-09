package parser

import (
	"fepl/lexer"
	"reflect"
)

type Parser struct {
	position int
	tokens   []lexer.Token
	rules    []rule
}

type rule struct {
	name        string
	expressions []string
}

type Node struct {
	Kind       string
	token      lexer.Token
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

func (rule *rule) matches(nodes []Node) bool {
	nodeKinds := []string{}
	for _, node := range nodes {
		nodeKinds = append(nodeKinds, node.Kind)
	}
	// zip
	matches := reflect.DeepEqual(rule.expressions, nodeKinds)
	return matches
}

func (parser *Parser) resolve() Node {
	// Convert token list into node list
	nodeList := []Node{}
	for _, token := range parser.tokens {
		nodeList = append(nodeList, Node{token.Kind, token, []*Node{}})
	}

	// Scan over nodes to match rules, resolving until one root node
mainloop:
	for len(nodeList) > 1 {
		for i := 0; i < len(nodeList); i++ {
			for j := i + 1; j <= len(nodeList); j++ {
				for _, rule := range parser.rules {

					match := rule.matches(nodeList[i:j])
					if match {
						node := Node{rule.name, lexer.Token{}, []*Node{}}
						for _, n := range nodeList {
							node.ChildNodes = append(node.ChildNodes, &n)
						}
						temp := append(nodeList[:i], node)
						temp = append(temp, nodeList[j:]...)
						nodeList = temp

						continue mainloop
					}
				}
			}
		}
	}

	return nodeList[0]
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

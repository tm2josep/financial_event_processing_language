package parser

import (
	"fepl/lexer"
	"strconv"
)

type Parser struct {
	pc int
	pt []lexer.Token
}

func (parser *Parser) current() lexer.Token {
	if parser.pc > len(parser.pt) || parser.pc < 0 {
		panic("parser index out of range")
	}

	return parser.pt[parser.pc]
}

func (parser *Parser) advance(steps ...int) {
	if len(steps) == 0 {
		parser.pc++
	} else if len(steps) == 1 {
		parser.pc += steps[0]
	} else {
		panic("parser.advance only accepts 0 or 1 arguments")
	}

	if parser.pc > len(parser.pt) {
		panic("parser advanced out of range")
	}
}

func (parser *Parser) GetAst(tokens chan lexer.Token) Ast {
	parser.pc = 0
	parser.pt = []lexer.Token{}
	// Consume all tokens into slice
	for token := range tokens {
		parser.pt = append(parser.pt, token)
	}

	ast := Ast{
		Kind: "Program",
		Body: []Node{},
	}

	var n Node
	for parser.pc < len(parser.pt) {
		n = parser.walk()
		if n.Kind == "" {
			continue
		}
		ast.Body = append(ast.Body, n)
	}

	return ast
}

func (parser *Parser) makeBinop(a Node) Node {
	token := parser.current()
	binopKind := ""

	switch token.Content {
	case "+":
		binopKind = "Add"
	case "-":
		binopKind = "Subtract"
	case "*":
		binopKind = "Multiply"
	case "/":
		binopKind = "Divide"
	case "^":
		binopKind = "Exponentiate"
	default:
		panic("unknown binary operator")
	}

	parser.advance()
	// Recursively build the sides if needed
	return Node{Kind: binopKind, Body: []Node{a, parser.walk()}}
}

func (parser *Parser) makeExpression() Node {
	// Starts with "("
	parser.advance()
	n := Node{
		Kind: "Expression",
		Body: []Node{},
	}

	token := parser.current()
	// construct left hand side of operator
	a := parser.walk()

	token = parser.current()
	if token.Kind == "BINOP" {
		// If there is an operator, get the right hand side
		n.Body = append(n.Body, parser.makeBinop(a))
	} else {
		n.Body = []Node{a}
	}

	return n
}

func (parser *Parser) allocationStatement() Node {
	// skip alloc keyword
	parser.advance()

	n := Node{
		Kind: "Allocation",
		Body: []Node{},
	}

	// get source field
	rhs := parser.walk()
	if rhs.Kind != "Field" {
		panic("expected allocation field source")
	}

	n.Body = append(n.Body, rhs)

	// get source value
	valueNode := parser.walk()

	if valueNode.Kind != "Expression" && valueNode.Kind != "Field" {
		panic("expected allocation value expression")
	}

	if valueNode.Kind == "Field" {
		valueNode = Node{
			Kind: "Expression",
			Body: []Node{valueNode},
		}
	}

	n.Body = append(n.Body, valueNode)

	// get source field
	lhs := parser.walk()

	if lhs.Kind != "Field" {
		panic("expected allocation field target")
	}

	n.Body = append(n.Body, lhs)

	return n
}

func (parser *Parser) makeField() Node {
	// skip field start
	parser.advance()

	// take field name
	n := Node{
		Kind:  "Field",
		Value: parser.current().Content,
	}

	// skip field end
	parser.advance(2)
	return n
}

func (parser *Parser) walk() Node {
	token := parser.current()

	if token.Kind == "RPAREN" {
		parser.advance()
		return Node{Kind: ""}
	}

	if token.Kind == "LPAREN" {
		return parser.makeExpression()
	}

	if token.Kind == "NUMBER" {
		parser.advance()
		fixedValue, _ := strconv.ParseFloat(token.Content, 64)
		return Node{
			Kind: "Expression",
			Body: []Node{
				{
					Kind:       "NumberLiteral",
					FixedValue: fixedValue,
				},
			},
		}
	}

	if token.Kind == "FIELD_START" {
		return parser.makeField()
	}

	if token.Kind == "ALLOC" {
		return parser.allocationStatement()
	}

	panic("unknown parsing error")
}

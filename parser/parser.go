package parser

import (
	"fepl/lexer"
	"fmt"
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

func (parser *Parser) advance() {
	parser.pc++
	if parser.pc > len(parser.pt) {
		panic("parser advanced out of range")
	}
}

func (parser *Parser) GetAst(tokens chan lexer.Token) ast {
	parser.pc = 0
	parser.pt = []lexer.Token{}
	// Consume all tokens into slice
	for token := range tokens {
		parser.pt = append(parser.pt, token)
	}

	ast := ast{
		kind: "Program",
		body: []node{},
	}

	for parser.pc < len(parser.pt) {
		ast.body = append(ast.body, parser.walk())
	}

	return ast
}

func (parser *Parser) makeBinop(a node) node {
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
	return node{kind: binopKind, body: []node{a, parser.walk()}}
}

func (parser *Parser) makeExpression() node {
	// Starts with "("
	parser.advance()
	n := node{
		kind: "Expression",
		body: []node{},
	}

	token := parser.current()
	// construct left hand side of operator
	a := parser.walk()

	token = parser.current()
	if token.Kind == "BINOP" {
		// If there is an operator, get the right hand side
		n.body = append(n.body, parser.makeBinop(a))
	} else {
		n.body = []node{a}
	}

	return n
}

func (parser *Parser) allocationStatement() node {
	// skip alloc keyword
	parser.advance()

	n := node{
		kind: "Allocation",
		body: []node{},
	}

	// get source field
	rhs := parser.walk()
	if rhs.kind != "Field" {
		panic("expected allocation field source")
	}

	n.body = append(n.body, rhs)

	// get source value
	valueNode := parser.walk()

	if valueNode.kind != "Expression" && valueNode.kind != "Field" {
		panic("expected allocation value expression")
	}

	n.body = append(n.body, valueNode)

	// get source field
	lhs := parser.walk()

	if lhs.kind != "Field" {
		panic("expected allocation field target")
	}

	n.body = append(n.body, lhs)

	return n
}

func (parser *Parser) makeField() node {
	// skip field start
	parser.advance()

	// take field name
	n := node{
		kind:  "Field",
		value: parser.current().Content,
	}

	// skip field end
	parser.pc += 2
	fmt.Println(n)
	return n
}

func (parser *Parser) walk() node {
	token := parser.current()
	fmt.Println(token)

	if token.Kind == "RPAREN" {
		parser.advance()
		return parser.walk()
	}

	if token.Kind == "LPAREN" {
		return parser.makeExpression()
	}

	if token.Kind == "NUMBER" {
		parser.advance()
		return node{
			kind: "Expression",
			body: []node{
				{
					kind:  "NumberLiteral",
					value: token.Content,
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

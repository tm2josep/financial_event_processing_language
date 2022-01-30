package behavior

import (
	"fepl/parser"
	"fmt"
	"math"
	"strconv"
)

type Expression struct {
	Body []Expression
}

func isBinOp(node parser.Node) bool {
	return (node.Kind == "Add" ||
		node.Kind == "Subtract" ||
		node.Kind == "Multiply" ||
		node.Kind == "Divide" ||
		node.Kind == "Exponentiate")
}

func doBinop(kind string, lhs float64, rhs float64) float64 {
	switch kind {
	case "Add":
		return lhs + rhs
	case "Subtract":
		return lhs - rhs
	case "Multiply":
		return lhs * rhs
	case "Divide":
		if rhs == 0 {
			return math.NaN()
		}
		return lhs / rhs
	case "Exponentiate":
		return math.Pow(rhs, lhs)
	}

	panic("unsupported operation")
}

func resolveExpression(node parser.Node, context map[string]string) float64 {
	if node.Kind == "NumberLiteral" {
		return node.FixedValue
	}

	if node.Kind == "Field" {
		num, err := strconv.ParseFloat(context[node.Value], 64)
		if err != nil {
			panic("float parser error")
		}
		return num
	}

	if node.Kind == "Expression" {
		return resolveExpression(node.Body[0], context)
	}

	if isBinOp(node) {
		node_a := node.Body[0]
		a := resolveExpression(node_a, context)
		node_b := node.Body[1]
		b := resolveExpression(node_b, context)
		return doBinop(node.Kind, a, b)
	}

	panic("unresolvable expression")
}

func Build(n parser.Ast) Expression {
	context := make(map[string]string)
	context["field"] = "500.00"
	for _, node := range n.Body {
		if node.Kind == "Expression" {
			value := resolveExpression(node, context)
			fmt.Println(value)
		}
	}

	expr := new(Expression)
	return *expr
}

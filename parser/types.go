package parser

type Node struct {
	Kind       string
	Value      string
	FixedValue float64
	Body       []Node
}

type Ast Node

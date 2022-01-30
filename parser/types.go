package parser

type node struct {
	kind  string
	value string
	body  []node
}

type ast node

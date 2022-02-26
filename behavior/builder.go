package behavior

import (
	"fepl/parser"
	"fmt"
)

type rule struct {
	parentNode string
}

type Builder struct {
	stack []parser.NodeWalk
}

func (builder *Builder) Build(nodeWalk chan parser.NodeWalk) {
	for node := range nodeWalk {
		fmt.Println(node.Node.Kind, node.Node.Token)
	}
}

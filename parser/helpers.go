package parser

import (
	"fmt"
	"unicode"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func printMatch(name string, nodeList []*Node) {
	fmt.Print(len(nodeList), " ")
	for _, n := range nodeList {
		fmt.Print(n.Kind + " ")
	}
	fmt.Print("---> " + name)
	fmt.Print("\n")
}

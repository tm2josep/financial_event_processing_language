package lexer

import (
	"regexp"
)

type rule struct {
	name string
	expr *regexp.Regexp
}

type Lexer struct {
	Location int
}

type Token struct {
	Name     string
	Content  string
	Location []int
}

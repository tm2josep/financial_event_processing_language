package lexer

import (
	"regexp"
)

type rule struct {
	name string
	expr *regexp.Regexp
}

type Token struct {
	Kind     string
	Content  string
	Location []int
}

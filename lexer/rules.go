package lexer

import "regexp"

var rules = []rule{
	{"WHITESPACE", regexp.MustCompile(`\s`)},
	{"LPAREN", regexp.MustCompile(`\(`)},
	{"RPAREN", regexp.MustCompile(`\)`)},
	{"ALLOC", regexp.MustCompile(`alloc`)},
	{"SCOPE", regexp.MustCompile(`scope`)},
	{"DISCARD", regexp.MustCompile(`discard`)},
	{"ASSESS", regexp.MustCompile(`assess`)},
	{"AGGREGATE", regexp.MustCompile(`aggregate`)},
	{"SET_VALUE", regexp.MustCompile(`set\_value`)},
	{"KEY", regexp.MustCompile(`key`)},
	{"EOL", regexp.MustCompile(`;`)},
	{"PERCENT", regexp.MustCompile(`\%`)},
	{"NUMBER", regexp.MustCompile(`(\-)?(([\d,]+\.\d+)|(\.\d+)|([\d,]+))`)},
	{"FIELD_START", regexp.MustCompile(`@[']`)},
	{"STRING", regexp.MustCompile(`\"(\\.|[^"\\])*\"`)},
	{"FIELD_END", regexp.MustCompile(`[']`)},
	{"AGG_MODE", regexp.MustCompile(`\:((sum)|(mean)|(median)|(mode)|(max)|(min)|(count))`)},
	{"VAR_START", regexp.MustCompile(`\$`)},
	{"NAME", regexp.MustCompile(`[A-Za-z]+[0-9A-Za-z\_]+`)},
	{"EXP", regexp.MustCompile(`\^`)},
	{"BINOP", regexp.MustCompile(`[\*\/\+\-\^]`)},
	{"COMPARATOR", regexp.MustCompile(`(==)|(>=)|(<=)|(>)|(<)}`)},
}

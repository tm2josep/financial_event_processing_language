package lexer

type Lexer struct {
	Location int
}

func (lex Lexer) peek(source string) Token {
	for _, rule := range rules {
		loc := rule.expr.FindIndex([]byte(source[lex.Location:]))
		if loc != nil && loc[0] == 0 {
			return Token{rule.name, source[(lex.Location + loc[0]):(lex.Location + loc[1])], loc}
		}
	}

	panic("No Match")
}

func (lex *Lexer) eat(source string) Token {
	tokenMatch := lex.peek(source)
	lex.Location += tokenMatch.Location[1]
	return tokenMatch
}

func (lex Lexer) Stream(source string, out chan Token) {
	for len(source) > lex.Location {
		token := lex.eat(source)
		if token.Kind == "WHITESPACE" {
			continue
		}
		out <- token
	}
	close(out)
}

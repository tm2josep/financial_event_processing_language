package parser

var rules = []rule{
	{"block", []string{"block", "statement"}},
	{"block", []string{"statement"}},
	{"statement", []string{"AGGREGATE", "field", "agg_fields", "EOL"}},
	{"statement", []string{"ALLOC", "field", "value", "field", "EOL"}},
	{"field", []string{"FIELD_START", "NAME", "FIELD_END"}},
	{"agg_field", []string{"field", "AGG_MODE"}},
	{"agg_fields", []string{"agg_fields", "agg_field"}},
	{"agg_fields", []string{"agg_field"}},
	{"value", []string{"LPAREN", "value", "RPAREN"}},
	{"value", []string{"NUMBER"}},
	{"value", []string{"value", "BINOP", "value"}},
	{"value", []string{"LPAREN", "field", "RPAREN"}},
}

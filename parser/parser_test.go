package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func tok(k token.Kind, v string) token.Token {
	return token.New(k, v, 0, 0)
}

func push(in chan token.Token, tokens ...token.Token) {
	go func() {
		for _, tk := range tokens {
			in <- tk
		}
		close(in)
	}()
}

func doTestParse(t *testing.T, exp Expr, tokens ...token.Token) {

	in := make(chan token.Token, len(tokens))
	p := New(in)

	push(in, tokens...)

	act := p.Parse()
	require.Equal(t, exp, act)
}

// Parse an assignment statement
// Parse a string literal assignment
func TestParser_parse_1(t *testing.T) {

	tokens := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.STR_LITERAL, "xyz"),
		tok(token.TERMINATOR, "\n"),
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // opener
		tokens[4],          // closer
		[]Stat{
			assignStat{
				tokenExpr{tokens[1]},
				[]token.Token{ // ids
					tokens[0],
				},
				[]Expr{ // srcs
					valueExpr{
						tokenExpr{tokens[2]},
						Value{STR, tokens[2].Value}, // v
					},
				},
			},
		},
	}

	doTestParse(t, exp, tokens...)
}

// Parse several assignment statements
// Parse a bool literal assignment
// Parse a real literal assignment
// Parse a string template assignment
func TestParser_parse_2(t *testing.T) {

	tokens := []token.Token{
		// Bool
		tok(token.ID, "a"),
		tok(token.ASSIGN, ":="),
		tok(token.BOOL_LITERAL, "TRUE"),
		tok(token.TERMINATOR, "\n"),
		// Number
		tok(token.ID, "b"),
		tok(token.ASSIGN, ":="),
		tok(token.REAL_LITERAL, "123.456"),
		tok(token.TERMINATOR, "\n"),
		// String template
		tok(token.ID, "c"),
		tok(token.ASSIGN, ":="),
		tok(token.STR_TEMPLATE, `"Caribbean"`),
		tok(token.TERMINATOR, "\n"),
		// EOF
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // opener
		tokens[12],         // closer
		[]Stat{
			assignStat{
				tokenExpr{tokens[1]},
				[]token.Token{tokens[0]}, // ids
				[]Expr{ // srcs
					valueExpr{tokenExpr{tokens[2]}, Value{BOOL, true}}, // v
				},
			},
			assignStat{
				tokenExpr{tokens[5]},
				[]token.Token{tokens[4]}, // ids
				[]Expr{ // srcs
					valueExpr{tokenExpr{tokens[6]}, Value{REAL, 123.456}}, // v
				},
			},
			assignStat{
				tokenExpr{tokens[9]},
				[]token.Token{tokens[8]}, // ids
				[]Expr{ // srcs
					valueExpr{tokenExpr{tokens[10]}, Value{STR, `"Caribbean"`}}, // v
				},
			},
		},
	}

	doTestParse(t, exp, tokens...)
}

// Parse multiple assignment statement
// Parse a bool literal assignment
// Parse a int literal assignment
// Parse a string template assignment
func TestParser_parse_3(t *testing.T) {

	tokens := []token.Token{
		// ids
		tok(token.ID, "a"),
		tok(token.DELIM, ","),
		tok(token.ID, "b"),
		tok(token.DELIM, ","),
		tok(token.ID, "c"),
		tok(token.ASSIGN, ":="),
		// srcs
		tok(token.BOOL_LITERAL, "TRUE"),
		tok(token.DELIM, ","),
		tok(token.REAL_LITERAL, "123.456"),
		tok(token.DELIM, ","),
		tok(token.STR_TEMPLATE, `"Caribbean"`),
		tok(token.TERMINATOR, "\n"),
		// EOF
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // opener
		tokens[12],         // closer
		[]Stat{
			assignStat{
				tokenExpr{tokens[5]},
				[]token.Token{ // ids
					tokens[0],
					tokens[2],
					tokens[4],
				},
				[]Expr{ // srcs
					valueExpr{tokenExpr{tokens[6]}, Value{BOOL, true}},
					valueExpr{tokenExpr{tokens[8]}, Value{REAL, 123.456}},
					valueExpr{tokenExpr{tokens[10]}, Value{STR, `"Caribbean"`}},
				},
			},
		},
	}

	doTestParse(t, exp, tokens...)
}

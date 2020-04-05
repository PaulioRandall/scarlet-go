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

	tks := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.STR, "xyz"),
		tok(token.TERMINATOR, "\n"),
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // opener
		tks[4],             // closer
		[]Stat{
			assignStat{
				tks[1],
				[]token.Token{ // ids
					tks[0],
				},
				[]Expr{ // srcs
					valueExpr{
						tks[2],
						Value{STR, tks[2].Value}, // v
					},
				},
			},
		},
	}

	doTestParse(t, exp, tks...)
}

// Parse several assignment statements
// Parse a bool literal assignment
// Parse a ID to ID assignment
func TestParser_parse_2(t *testing.T) {

	tks := []token.Token{
		// Bool
		tok(token.ID, "a"),
		tok(token.ASSIGN, ":="),
		tok(token.BOOL, "TRUE"),
		tok(token.TERMINATOR, "\n"), // 3
		// Number
		tok(token.ID, "b"),
		tok(token.ASSIGN, ":="),
		tok(token.REAL, "123.456"),
		tok(token.TERMINATOR, "\n"), // 7
		// String template
		tok(token.ID, "c"),
		tok(token.ASSIGN, ":="),
		tok(token.ID, "b"),
		tok(token.TERMINATOR, "\n"), // 11
		// EOF
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // opener
		tks[12],            // closer
		[]Stat{
			assignStat{
				tks[1],
				[]token.Token{tks[0]}, // ids
				[]Expr{ // srcs
					valueExpr{tks[2], Value{BOOL, true}}, // v
				},
			},
			assignStat{
				tks[5],
				[]token.Token{tks[4]}, // ids
				[]Expr{ // srcs
					valueExpr{tks[6], Value{REAL, float64(123.456)}}, // v
				},
			},
			assignStat{
				tks[9],
				[]token.Token{tks[8]}, // ids
				[]Expr{ // srcs
					idExpr{tks[10], "b"}, // v
				},
			},
		},
	}

	doTestParse(t, exp, tks...)
}

// Parse multiple assignment statement
// Parse a bool literal assignment
// Parse a int literal assignment
// Parse a string template assignment
func TestParser_parse_3(t *testing.T) {

	tks := []token.Token{
		// ids
		tok(token.ID, "a"),
		tok(token.DELIM, ","),
		tok(token.ID, "b"),
		tok(token.DELIM, ","),
		tok(token.ID, "c"),
		tok(token.ASSIGN, ":="),
		// srcs
		tok(token.BOOL, "TRUE"),
		tok(token.DELIM, ","),
		tok(token.INT, "123"),
		tok(token.DELIM, ","),
		tok(token.TEMPLATE, `"Caribbean"`),
		tok(token.TERMINATOR, "\n"),
		// EOF
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // opener
		tks[12],            // closer
		[]Stat{
			assignStat{
				tks[5],
				[]token.Token{ // ids
					tks[0],
					tks[2],
					tks[4],
				},
				[]Expr{ // srcs
					valueExpr{tks[6], Value{BOOL, true}},
					valueExpr{tks[8], Value{INT, int64(123)}},
					valueExpr{tks[10], Value{STR, `"Caribbean"`}},
				},
			},
		},
	}

	doTestParse(t, exp, tks...)
}

// Parse list assignment statement
// Parse a list within a list
// Parse a list with a comma after the last item
func TestParser_parse_4(t *testing.T) {

	tks := []token.Token{
		// Line 1
		tok(token.ID, "list"),
		tok(token.ASSIGN, ":="),
		tok(token.OPEN_LIST, "{"),
		tok(token.TERMINATOR, "\n"), // index: 3
		// Line 2
		tok(token.STR, "abc"),
		tok(token.DELIM, ","),
		tok(token.REAL, "123.456"),
		tok(token.DELIM, ","),
		tok(token.TERMINATOR, "\n"), // 8
		// Line 3
		tok(token.OPEN_LIST, "{"),
		tok(token.TEMPLATE, "xyz"),
		tok(token.DELIM, ","),
		tok(token.BOOL, "TRUE"), // 12
		tok(token.CLOSE_LIST, "}"),
		tok(token.DELIM, ","),
		tok(token.TERMINATOR, "\n"), // 15
		// Line 4
		tok(token.CLOSE_LIST, "}"),
		tok(token.TERMINATOR, "\n"), // 17
		// EOF
		tok(token.EOF, ""),
	}

	exp := blockStat{
		tok(token.SOF, ""), // field: opener
		tks[18],            // closer
		[]Stat{
			assignStat{
				tks[1],
				[]token.Token{tks[0]}, // ids
				[]Expr{ // srcs
					listExpr{
						tks[2],  // start
						tks[16], // end
						[]Expr{ // items
							valueExpr{tks[4], Value{STR, "abc"}},
							valueExpr{tks[6], Value{REAL, float64(123.456)}},
							listExpr{
								tks[9],  // start
								tks[13], // end
								[]Expr{ // items
									valueExpr{tks[10], Value{STR, "xyz"}},
									valueExpr{tks[12], Value{BOOL, true}},
								},
							},
						},
					},
				},
			},
		},
	}

	doTestParse(t, exp, tks...)
}

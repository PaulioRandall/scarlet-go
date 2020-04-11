package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/stretchr/testify/require"
)

func tok(l lexeme.Lexeme, v string) lexeme.Token {
	return lexeme.Token{l, v, 0, 0}
}

func push(in chan lexeme.Token, tokens ...lexeme.Token) {
	go func() {
		for _, tk := range tokens {
			in <- tk
		}
		close(in)
	}()
}

func doTestParse(t *testing.T, exp Expr, tokens ...lexeme.Token) {

	in := make(chan lexeme.Token, len(tokens))
	p := New(in)

	push(in, tokens...)

	act := p.Parse()
	require.Equal(t, exp, act)
}

// Parse an assignment statement
// Parse a string literal assignment
func TestParser_parse_1(t *testing.T) {

	tks := []lexeme.Token{
		tok(lexeme.LEXEME_ID, "abc"),
		tok(lexeme.LEXEME_ASSIGN, ":="),
		tok(lexeme.LEXEME_STRING, "xyz"),
		tok(lexeme.LEXEME_TERMINATOR, "\n"),
		tok(lexeme.LEXEME_EOF, ""),
	}

	exp := blockStat{
		tok(lexeme.LEXEME_SOF, ""), // opener
		tks[4],                     // closer
		[]Stat{
			assignStat{
				tks[1],
				[]lexeme.Token{ // ids
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

	tks := []lexeme.Token{
		// Bool
		tok(lexeme.LEXEME_ID, "a"),
		tok(lexeme.LEXEME_ASSIGN, ":="),
		tok(lexeme.LEXEME_BOOL, "TRUE"),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // 3
		// Number
		tok(lexeme.LEXEME_ID, "b"),
		tok(lexeme.LEXEME_ASSIGN, ":="),
		tok(lexeme.LEXEME_FLOAT, "123.456"),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // 7
		// String template
		tok(lexeme.LEXEME_ID, "c"),
		tok(lexeme.LEXEME_ASSIGN, ":="),
		tok(lexeme.LEXEME_ID, "b"),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // 11
		// EOF
		tok(lexeme.LEXEME_EOF, ""),
	}

	exp := blockStat{
		tok(lexeme.LEXEME_SOF, ""), // opener
		tks[12],                    // closer
		[]Stat{
			assignStat{
				tks[1],
				[]lexeme.Token{tks[0]}, // ids
				[]Expr{ // srcs
					valueExpr{tks[2], Value{BOOL, true}}, // v
				},
			},
			assignStat{
				tks[5],
				[]lexeme.Token{tks[4]}, // ids
				[]Expr{ // srcs
					valueExpr{tks[6], Value{REAL, float64(123.456)}}, // v
				},
			},
			assignStat{
				tks[9],
				[]lexeme.Token{tks[8]}, // ids
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

	tks := []lexeme.Token{
		// ids
		tok(lexeme.LEXEME_ID, "a"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_ID, "b"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_ID, "c"),
		tok(lexeme.LEXEME_ASSIGN, ":="),
		// srcs
		tok(lexeme.LEXEME_BOOL, "TRUE"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_INT, "123"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_TEMPLATE, `"Caribbean"`),
		tok(lexeme.LEXEME_TERMINATOR, "\n"),
		// EOF
		tok(lexeme.LEXEME_EOF, ""),
	}

	exp := blockStat{
		tok(lexeme.LEXEME_SOF, ""), // opener
		tks[12],                    // closer
		[]Stat{
			assignStat{
				tks[5],
				[]lexeme.Token{ // ids
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

	tks := []lexeme.Token{
		// Line 1
		tok(lexeme.LEXEME_ID, "list"),
		tok(lexeme.LEXEME_ASSIGN, ":="),
		tok(lexeme.LEXEME_OPEN_LIST, "{"),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // index: 3
		// Line 2
		tok(lexeme.LEXEME_STRING, "abc"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_FLOAT, "123.456"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // 8
		// Line 3
		tok(lexeme.LEXEME_OPEN_LIST, "{"),
		tok(lexeme.LEXEME_TEMPLATE, "xyz"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_BOOL, "TRUE"), // 12
		tok(lexeme.LEXEME_CLOSE_LIST, "}"),
		tok(lexeme.LEXEME_DELIM, ","),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // 15
		// Line 4
		tok(lexeme.LEXEME_CLOSE_LIST, "}"),
		tok(lexeme.LEXEME_TERMINATOR, "\n"), // 17
		// EOF
		tok(lexeme.LEXEME_EOF, ""),
	}

	exp := blockStat{
		tok(lexeme.LEXEME_SOF, ""), // field: opener
		tks[18],                    // closer
		[]Stat{
			assignStat{
				tks[1],
				[]lexeme.Token{tks[0]}, // ids
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

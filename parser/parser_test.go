package parser

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/stretchr/testify/require"
)

func push(in chan Token, tokens ...Token) {
	go func() {
		for _, tk := range tokens {
			in <- tk
		}
		close(in)
	}()
}

func doTestParse(t *testing.T, exp Expr, tokens ...Token) {

	in := make(chan Token, len(tokens))
	p := New(in)

	push(in, tokens...)

	act := p.Parse()
	require.Equal(t, exp, act)
}

// Parse an assignment statement
// Parse a string literal assignment
func TestParser_parse_1(t *testing.T) {

	tks := []Token{
		Token{LEXEME_ID, "abc", 0, 0},
		Token{LEXEME_ASSIGN, ":=", 0, 0},
		Token{LEXEME_STRING, "xyz", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0},
		Token{LEXEME_EOF, "", 0, 0},
	}

	exp := blockStat{
		Token{LEXEME_SOF, "", 0, 0}, // opener
		tks[4],                      // closer
		[]Stat{
			assignStat{
				tks[1],
				[]Token{ // ids
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

	tks := []Token{
		// Bool
		Token{LEXEME_ID, "a", 0, 0},
		Token{LEXEME_ASSIGN, ":=", 0, 0},
		Token{LEXEME_BOOL, "TRUE", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // 3
		// Number
		Token{LEXEME_ID, "b", 0, 0},
		Token{LEXEME_ASSIGN, ":=", 0, 0},
		Token{LEXEME_FLOAT, "123.456", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // 7
		// String template
		Token{LEXEME_ID, "c", 0, 0},
		Token{LEXEME_ASSIGN, ":=", 0, 0},
		Token{LEXEME_ID, "b", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // 11
		// EOF
		Token{LEXEME_EOF, "", 0, 0},
	}

	exp := blockStat{
		Token{LEXEME_SOF, "", 0, 0}, // opener
		tks[12],                     // closer
		[]Stat{
			assignStat{
				tks[1],
				[]Token{tks[0]}, // ids
				[]Expr{ // srcs
					valueExpr{tks[2], Value{BOOL, true}}, // v
				},
			},
			assignStat{
				tks[5],
				[]Token{tks[4]}, // ids
				[]Expr{ // srcs
					valueExpr{tks[6], Value{REAL, float64(123.456)}}, // v
				},
			},
			assignStat{
				tks[9],
				[]Token{tks[8]}, // ids
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

	tks := []Token{
		// ids
		Token{LEXEME_ID, "a", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_ID, "b", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_ID, "c", 0, 0},
		Token{LEXEME_ASSIGN, ":=", 0, 0},
		// srcs
		Token{LEXEME_BOOL, "TRUE", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_INT, "123", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_TEMPLATE, `"Caribbean"`, 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0},
		// EOF
		Token{LEXEME_EOF, "", 0, 0},
	}

	exp := blockStat{
		Token{LEXEME_SOF, "", 0, 0}, // opener
		tks[12],                     // closer
		[]Stat{
			assignStat{
				tks[5],
				[]Token{ // ids
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

	tks := []Token{
		// Line 1
		Token{LEXEME_ID, "list", 0, 0},
		Token{LEXEME_ASSIGN, ":=", 0, 0},
		Token{LEXEME_LIST_OPEN, "[", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // index: 3
		// Line 2
		Token{LEXEME_STRING, "abc", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_FLOAT, "123.456", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // 8
		// Line 3
		Token{LEXEME_LIST_OPEN, "[", 0, 0},
		Token{LEXEME_TEMPLATE, "xyz", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_BOOL, "TRUE", 0, 0}, // 12
		Token{LEXEME_LIST_CLOSE, "]", 0, 0},
		Token{LEXEME_DELIM, ",", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // 15
		// Line 4
		Token{LEXEME_LIST_CLOSE, "]", 0, 0},
		Token{LEXEME_TERMINATOR, "\n", 0, 0}, // 17
		// EOF
		Token{LEXEME_EOF, "", 0, 0},
	}

	exp := blockStat{
		Token{LEXEME_SOF, "", 0, 0}, // field: opener
		tks[18],                     // closer
		[]Stat{
			assignStat{
				tks[1],
				[]Token{tks[0]}, // ids
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

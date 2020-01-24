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

func doTestParseAssign(t *testing.T, exp Expr, tokens ...token.Token) {

	in := make(chan token.Token, len(tokens))
	p := New(in)

	push(in, tokens...)
	act := p.parseAssign(<-in)
	require.Equal(t, exp, act)
}

func TestParser_parseAssign_1(t *testing.T) {
	// Parse string literal

	tokens := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.STR_LITERAL, "She said yes!"),
		tok(token.TERMINATOR, "\n"),
	}

	exp := assignStat{
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
	}

	doTestParseAssign(t, exp, tokens...)
}

func TestParser_parseAssign_2(t *testing.T) {
	// Parse string template

	tokens := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.BOOL_LITERAL, "TRUE"),
		tok(token.TERMINATOR, "\n"),
	}

	exp := assignStat{
		tokenExpr{tokens[1]},
		[]token.Token{ // ids
			tokens[0],
		},
		[]Expr{ // srcs
			valueExpr{
				tokenExpr{tokens[2]},
				Value{BOOL, true}, // v
			},
		},
	}

	doTestParseAssign(t, exp, tokens...)
}

func TestParser_parseAssign_3(t *testing.T) {
	// Parse number

	tokens := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.REAL_LITERAL, "123.456"),
		tok(token.TERMINATOR, "\n"),
	}

	exp := assignStat{
		tokenExpr{tokens[1]},
		[]token.Token{ // ids
			tokens[0],
		},
		[]Expr{ // srcs
			valueExpr{
				tokenExpr{tokens[2]},
				Value{REAL, 123.456}, // v
			},
		},
	}

	doTestParseAssign(t, exp, tokens...)
}

func TestParser_parseAssign_4(t *testing.T) {
	// Parse string templates

	tokens := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.STR_TEMPLATE, `"Caribbean"`),
		tok(token.TERMINATOR, "\n"),
	}

	exp := assignStat{
		tokenExpr{tokens[1]},
		[]token.Token{ // ids
			tokens[0],
		},
		[]Expr{ // srcs
			valueExpr{
				tokenExpr{tokens[2]},
				Value{STR, `"Caribbean"`}, // v
			},
		},
	}

	doTestParseAssign(t, exp, tokens...)
}

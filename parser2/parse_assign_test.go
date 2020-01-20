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
	}()
}

func doTest(t *testing.T, exp Expr, tokens ...token.Token) {

	in := make(chan token.Token, len(tokens))
	out := make(chan Expr)
	p := New(in, out)

	push(in, tokens...)
	act := p.parseAssign()
	require.Equal(t, exp, act)
}

func TestParser_parseAssign(t *testing.T) {

	tokens := []token.Token{
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.STR_LITERAL, "123"),
	}

	exp := assignExpr{
		tokenExpr{tokens[1]},
		tokens[0], // id
		valueExpr{ // src
			tokenExpr{tokens[2]},
			Value{STR, tokens[2].Value}, // v
		},
	}

	println(exp.String())

	doTest(t, exp, tokens...)
}

package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func doTestParseStats(t *testing.T, exp Expr, tokens ...token.Token) {

	in := make(chan token.Token, len(tokens))
	p := New(in)

	push(in, tokens...)
	act := p.parseStats(<-in)

	require.Equal(t, exp, act)
}

func TestParser_parseStats(t *testing.T) {

	tokens := []token.Token{
		tok(token.DO, "DO"),
		tok(token.ID, "abc"),
		tok(token.ASSIGN, ":="),
		tok(token.STR_LITERAL, "123"),
		tok(token.TERMINATOR, "\n"),
		tok(token.END, "END"),
	}

	exp := blockStat{
		tokens[0], // opener
		tokens[5], // closer
		[]Stat{
			assignStat{
				tokenExpr{tokens[2]},
				[]token.Token{ // assignStat.id
					tokens[1],
				},
				[]Expr{ // assignStat.srcs
					valueExpr{
						tokenExpr{tokens[3]},
						Value{STR, tokens[3].Value}, // assignStat.v
					},
				},
			},
		},
	}

	doTestParseStats(t, exp, tokens...)
}

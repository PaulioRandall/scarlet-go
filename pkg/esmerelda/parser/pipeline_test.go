package parser

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"

	"github.com/stretchr/testify/require"
)

func Test_P2(t *testing.T) {

	// p.hasMore(), p.match(), p.accept()

	tok := func(ty TokenType) Token {
		return NewToken(ty, "", 0, 0)
	}

	p := newPipeline(&tkStream{[]Token{
		tok(TK_NUMBER),
		tok(TK_PLUS),
		tok(TK_NUMBER),
	}})

	require.Equal(t, true, p.hasMore())
	require.Equal(t, false, p.empty())

	require.Equal(t, true, p.match(TK_NUMBER))
	require.Equal(t, false, p.accept(TK_PLUS))
	require.Equal(t, true, p.accept(TK_NUMBER))

	require.Equal(t, true, p.hasMore())
	require.Equal(t, false, p.empty())

	require.Equal(t, true, p.match(TK_PLUS))
	require.Equal(t, false, p.accept(TK_NUMBER))
	require.Equal(t, true, p.accept(TK_PLUS))

	require.Equal(t, true, p.hasMore())
	require.Equal(t, false, p.empty())

	require.Equal(t, true, p.match(TK_NUMBER))
	require.Equal(t, false, p.accept(TK_PLUS))
	require.Equal(t, true, p.accept(TK_NUMBER))

	require.Equal(t, false, p.hasMore())
	require.Equal(t, true, p.empty())

	require.Equal(t, false, p.match(TK_NUMBER))
	require.Equal(t, false, p.accept(TK_NUMBER))

	p.backup()

	require.Equal(t, true, p.match(TK_NUMBER))
	require.Equal(t, false, p.accept(TK_PLUS))
	require.Equal(t, true, p.accept(TK_NUMBER))

	require.Equal(t, false, p.hasMore())
	require.Equal(t, true, p.empty())

	require.Equal(t, false, p.match(TK_NUMBER))
	require.Equal(t, false, p.accept(TK_NUMBER))
}

func Test_P3(t *testing.T) {

	// p.expect()

	tok := func(ty TokenType) Token {
		return NewToken(ty, "", 0, 0)
	}

	checkOk := func(t *testing.T, p *pipeline, exp Token, ty TokenType) {
		act, e := p.expect(ty)
		require.Nil(t, nil, e)
		require.Equal(t, exp, act)
	}

	checkErr := func(t *testing.T, p *pipeline, ty TokenType) {
		act, e := p.expect(ty)
		require.NotNil(t, e)
		require.Nil(t, nil, act)
	}

	p := newPipeline(&tkStream{[]Token{
		tok(TK_NUMBER),
		tok(TK_PLUS),
		tok(TK_NUMBER),
	}})

	checkErr(t, p, TK_PLUS)
	checkOk(t, p, tok(TK_NUMBER), TK_NUMBER)

	checkErr(t, p, TK_NUMBER)
	checkOk(t, p, tok(TK_PLUS), TK_PLUS)

	checkErr(t, p, TK_PLUS)
	checkOk(t, p, tok(TK_NUMBER), TK_NUMBER)

	checkErr(t, p, TK_NUMBER)
}

func Test_P4(t *testing.T) {

	// p.expectAnyOf()

	tok := func(ty TokenType) Token {
		return NewToken(ty, "", 0, 0)
	}

	checkOk := func(t *testing.T, p *pipeline, exp Token, tys ...TokenType) {
		act, e := p.expectAnyOf(tys...)
		require.Nil(t, nil, e)
		require.Equal(t, exp, act)
	}

	checkErr := func(t *testing.T, p *pipeline, tys ...TokenType) {
		act, e := p.expectAnyOf(tys...)
		require.NotNil(t, e)
		require.Nil(t, nil, act)
	}

	p := newPipeline(&tkStream{[]Token{
		tok(TK_NUMBER),
		tok(TK_PLUS),
		tok(TK_NUMBER),
	}})

	checkErr(t, p, TK_PLUS)
	checkOk(t, p, tok(TK_NUMBER), TK_PLUS, TK_NUMBER)

	checkErr(t, p, TK_NUMBER)
	checkOk(t, p, tok(TK_PLUS), TK_PLUS, TK_NUMBER)

	checkErr(t, p, TK_PLUS)
	checkOk(t, p, tok(TK_NUMBER), TK_PLUS, TK_NUMBER)

	checkErr(t, p, TK_PLUS, TK_NUMBER)
}

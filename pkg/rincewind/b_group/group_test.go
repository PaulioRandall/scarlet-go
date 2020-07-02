package group

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/stat"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"
	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"
	"github.com/stretchr/testify/require"
)

type dummyStream struct {
	tks  []Token
	size int
	idx  int
}

func (d *dummyStream) Next() Token {

	if d.idx >= d.size {
		return nil
	}

	tk := d.tks[d.idx]
	d.idx++
	return tk
}

func doTest(t *testing.T, in []Token, exps []grp) {

	require.NotNil(t, exps, "SANITY CHECK! Expected grp missing")

	stream := &dummyStream{
		tks:  in,
		size: len(in),
	}
	acts := []grp{}

	var (
		g grp
		f GroupFunc
		e error
	)

	for f = New(stream); f != nil; {
		g, f, e = f()
		pet.RequireNil(t, e)
		acts = append(acts, g)
	}

	requireGrp(t, exps, acts)
}

func doErrorTest(t *testing.T, in []Token) {

	s := &dummyStream{
		tks:  in,
		size: len(in),
	}

	var e error
	for f := New(s); f != nil; {
		if _, f, e = f(); e != nil {
			return
		}
	}

	require.Fail(t, "Expected error")
}

func requireGrp(t *testing.T, exps, acts []grp) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+exps[i].String()+")\n...but no actual grps remain")

		require.True(t, i < expSize,
			"Did not expect any more grps\n...but got ("+acts[i].String()+")")

		require.Equal(t, exps[i].st, acts[i].st,
			"Want: %s, given: %s", exps[i].st.String(), acts[i].st.String())

		tkt.RequireSlice(t, exps[i].tks, acts[i].tks)
	}
}

func Test1_1(t *testing.T) {

	// @Set(x, 1)
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Set"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.HalfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	exp := []grp{
		grp{
			st:  ST_SPELL_CALL,
			tks: in,
		},
	}

	doTest(t, in, exp)
}

package scan

import (
	"testing"

	tkn "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"
	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"

	"github.com/stretchr/testify/require"
)

type dummyItr struct {
	symbols []rune
	size    int
	i       int
}

func (d *dummyItr) Next() (rune, bool) {

	if d.i >= d.size {
		return rune(0), false
	}

	ru := d.symbols[d.i]
	d.i++
	return ru, true
}

func doTest(t *testing.T, in string, exps []tkn.Token) {

	require.NotNil(t, exps, "SANITY CHECK! Expected tokens missing")

	itr := &dummyItr{
		symbols: []rune(in),
		size:    len(in),
	}
	acts := []tkn.Token{}

	var (
		tk tok
		f  ScanFunc
		e  error
	)

	for f = New(itr); f != nil; {
		tk, f, e = f()
		pet.RequireNil(t, e)
		acts = append(acts, tk)
	}

	tkt.RequireSlice(t, exps, acts)
}

func doErrorTest(t *testing.T, in string) {

	itr := &dummyItr{
		symbols: []rune(in),
		size:    len(in),
	}

	var e error
	for f := New(itr); f != nil; {
		if _, f, e = f(); e != nil {
			return
		}
	}

	require.Fail(t, "Expected error")
}

func Test_S1(t *testing.T) {

	in := "@Set(x, 1)"

	exp := []tkn.Token{
		tkt.Tok(tkn.GE_SPELL, tkn.SU_UNDEFINED, "@Set", 0, 0, 4),
		tkt.Tok(tkn.GE_PARENTHESIS, tkn.SU_PAREN_OPEN, "(", 0, 4, 5),
		tkt.Tok(tkn.GE_IDENTIFIER, tkn.SU_IDENTIFIER, "x", 0, 5, 6),
		tkt.Tok(tkn.GE_DELIMITER, tkn.SU_VALUE_DELIM, ",", 0, 6, 7),
		tkt.Tok(tkn.GE_WHITESPACE, tkn.SU_UNDEFINED, " ", 0, 7, 8),
		tkt.Tok(tkn.GE_LITERAL, tkn.SU_NUMBER, "1", 0, 8, 9),
		tkt.Tok(tkn.GE_PARENTHESIS, tkn.SU_PAREN_CLOSE, ")", 0, 9, 10),
	}

	doTest(t, in, exp)
}

func Test_T0_1(t *testing.T) {
	doErrorTest(t, "~")
}

func Test_T1_1(t *testing.T) {
	doTest(t, " \t\v\f", []tkn.Token{
		tkt.HalfTok(tkn.GE_WHITESPACE, tkn.SU_UNDEFINED, " \t\v\f"),
	})
}

func Test_T2_1(t *testing.T) {
	doTest(t, ";", []tkn.Token{
		tkt.HalfTok(tkn.GE_TERMINATOR, tkn.SU_TERMINATOR, ";"),
	})
}

func Test_T2_2(t *testing.T) {
	doTest(t, "\n", []tkn.Token{
		tkt.HalfTok(tkn.GE_TERMINATOR, tkn.SU_NEWLINE, "\n"),
	})
}

func Test_T2_3(t *testing.T) {
	doTest(t, "\r\n", []tkn.Token{
		tkt.HalfTok(tkn.GE_TERMINATOR, tkn.SU_NEWLINE, "\r\n"),
	})
}

func Test_T2_4(t *testing.T) {
	doErrorTest(t, "\r")
}

func Test_T3_1(t *testing.T) {
	doTest(t, "false", []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_BOOL, "false"),
	})
}

func Test_T3_2(t *testing.T) {
	doTest(t, "true", []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_BOOL, "true"),
	})
}

func Test_T4_1(t *testing.T) {
	doTest(t, "1", []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_NUMBER, "1"),
	})
}

func Test_T4_2(t *testing.T) {
	doTest(t, "123", []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_NUMBER, "123"),
	})
}

func Test_T4_3(t *testing.T) {
	doTest(t, "1.0", []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_NUMBER, "1.0"),
	})
}

func Test_T4_4(t *testing.T) {
	doTest(t, "123.456", []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_NUMBER, "123.456"),
	})
}

func Test_T4_5(t *testing.T) {
	doErrorTest(t, "123.")
}

func Test_T5_1(t *testing.T) {
	doTest(t, `""`, []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_STRING, `""`),
	})
}

func Test_T5_2(t *testing.T) {
	doTest(t, `"abc"`, []tkn.Token{
		tkt.HalfTok(tkn.GE_LITERAL, tkn.SU_STRING, `"abc"`),
	})
}

func Test_T5_3(t *testing.T) {
	doErrorTest(t, `"`)
}

func Test_T5_4(t *testing.T) {
	doErrorTest(t, `"abc`)
}

func Test_T5_5(t *testing.T) {
	doErrorTest(t, `"\"`)
}

func Test_T5_6(t *testing.T) {
	doErrorTest(t, `"\"\"abc\"\"`)
}

func Test_T6_1(t *testing.T) {
	doTest(t, "a", []tkn.Token{
		tkt.HalfTok(tkn.GE_IDENTIFIER, tkn.SU_IDENTIFIER, "a"),
	})
}

func Test_T6_2(t *testing.T) {
	doTest(t, "abc", []tkn.Token{
		tkt.HalfTok(tkn.GE_IDENTIFIER, tkn.SU_IDENTIFIER, "abc"),
	})
}

func Test_T6_3(t *testing.T) {
	doTest(t, "a_b", []tkn.Token{
		tkt.HalfTok(tkn.GE_IDENTIFIER, tkn.SU_IDENTIFIER, "a_b"),
	})
}

func Test_T6_4(t *testing.T) {
	doTest(t, "ab_", []tkn.Token{
		tkt.HalfTok(tkn.GE_IDENTIFIER, tkn.SU_IDENTIFIER, "ab_"),
	})
}

func Test_T6_5(t *testing.T) {
	doTest(t, "_", []tkn.Token{
		tkt.HalfTok(tkn.GE_IDENTIFIER, tkn.SU_VOID, "_"),
	})
}

func Test_T7_1(t *testing.T) {
	doTest(t, "(", []tkn.Token{
		tkt.HalfTok(tkn.GE_PARENTHESIS, tkn.SU_PAREN_OPEN, "("),
	})
}

func Test_T7_2(t *testing.T) {
	doTest(t, ")", []tkn.Token{
		tkt.HalfTok(tkn.GE_PARENTHESIS, tkn.SU_PAREN_CLOSE, ")"),
	})
}

func Test_T8_1(t *testing.T) {
	doTest(t, "@abc", []tkn.Token{
		tkt.HalfTok(tkn.GE_SPELL, tkn.SU_UNDEFINED, "@abc"),
	})
}

func Test_T8_2(t *testing.T) {
	doTest(t, "@abc.xyz", []tkn.Token{
		tkt.HalfTok(tkn.GE_SPELL, tkn.SU_UNDEFINED, "@abc.xyz"),
	})
}

func Test_T8_3(t *testing.T) {
	doTest(t, "@a.b.c.d", []tkn.Token{
		tkt.HalfTok(tkn.GE_SPELL, tkn.SU_UNDEFINED, "@a.b.c.d"),
	})
}

func Test_T8_4(t *testing.T) {
	doErrorTest(t, "@")
}

func Test_T8_5(t *testing.T) {
	doErrorTest(t, "@.")
}

func Test_T8_6(t *testing.T) {
	doErrorTest(t, "@a.")
}

func Test_T8_7(t *testing.T) {
	doErrorTest(t, "@a..a")
}

func Test_T9_1(t *testing.T) {
	doTest(t, ",", []tkn.Token{
		tkt.HalfTok(tkn.GE_DELIMITER, tkn.SU_VALUE_DELIM, ","),
	})
}

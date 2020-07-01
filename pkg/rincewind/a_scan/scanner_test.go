package scan

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

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

func doTest(t *testing.T, in string, exps []tok) {

	require.NotNil(t, exps, "SANITY CHECK! Expected tokens missing")

	itr := &dummyItr{
		symbols: []rune(in),
		size:    len(in),
	}
	acts := []tok{}

	var (
		tk tok
		f  ScanFunc
		e  error
	)

	for f = New(itr); f != nil; {
		tk, f, e = f()
		requireNilError(t, e)
		acts = append(acts, tk)
	}

	requireTokSlice(t, exps, acts)
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

func requireNilError(t *testing.T, e error) {

	if e == nil {
		return
	}

	s := fmt.Sprintf("%+v", e)
	require.Fail(t, s)
}

func requireTokSlice(t *testing.T, exps, acts []tok) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected ("+exps[i].String()+")\nBut no actual tokens remain")

		require.True(t, i < expSize,
			"Did not expect any more tokens\nBut got ("+acts[i].String()+")")

		requireTok(t, exps[i], acts[i])
	}
}

func requireTok(t *testing.T, exp, act tok) {

	require.NotNil(t, act, "Expected token ("+exp.String()+")\nBut got nil")
	m := "Expected (" + exp.String() + ")\nActual   (" + act.String() + ")"

	require.Equal(t, exp.ge, act.ge, m)
	require.Equal(t, exp.su, act.su, m)
	require.Equal(t, exp.raw, act.raw, m)

	require.Equal(t, exp.line, act.line, m)
	require.Equal(t, exp.colBegin, act.colBegin, m)
	require.Equal(t, exp.colEnd, act.colEnd, m)
}

func fullTok(ge GenType, su SubType, raw string, line, colBegin, colEnd int) tok {
	return tok{
		ge:       ge,
		su:       su,
		raw:      raw,
		line:     line,
		colBegin: colBegin,
		colEnd:   colEnd,
	}
}

func halfTok(ge GenType, su SubType, raw string) tok {
	return tok{
		ge:     ge,
		su:     su,
		raw:    raw,
		colEnd: len(raw),
	}
}

func Test_S1(t *testing.T) {

	in := "@Set(x, 1)"

	exp := []tok{
		fullTok(GE_SPELL, SU_UNDEFINED, "@Set", 0, 0, 4),
		fullTok(GE_PARENTHESIS, SU_PAREN_OPEN, "(", 0, 4, 5),
		fullTok(GE_IDENTIFIER, SU_IDENTIFIER, "x", 0, 5, 6),
		fullTok(GE_DELIMITER, SU_VALUE_DELIM, ",", 0, 6, 7),
		fullTok(GE_WHITESPACE, SU_UNDEFINED, " ", 0, 7, 8),
		fullTok(GE_LITERAL, SU_NUMBER, "1", 0, 8, 9),
		fullTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")", 0, 9, 10),
	}

	doTest(t, in, exp)
}

func Test_T1_1(t *testing.T) {
	doTest(t, " \t\v\f", []tok{halfTok(GE_WHITESPACE, SU_UNDEFINED, " \t\v\f")})
}

func Test_T2_1(t *testing.T) {
	doTest(t, ";", []tok{halfTok(GE_TERMINATOR, SU_TERMINATOR, ";")})
}

func Test_T2_2(t *testing.T) {
	doTest(t, "\n", []tok{halfTok(GE_TERMINATOR, SU_NEWLINE, "\n")})
}

func Test_T2_3(t *testing.T) {
	doTest(t, "\r\n", []tok{halfTok(GE_TERMINATOR, SU_NEWLINE, "\r\n")})
}

func Test_T3_1(t *testing.T) {
	doTest(t, "false", []tok{halfTok(GE_LITERAL, SU_BOOL, "false")})
}

func Test_T3_2(t *testing.T) {
	doTest(t, "true", []tok{halfTok(GE_LITERAL, SU_BOOL, "true")})
}

func Test_T4_1(t *testing.T) {
	doTest(t, "1", []tok{halfTok(GE_LITERAL, SU_NUMBER, "1")})
}

func Test_T4_2(t *testing.T) {
	doTest(t, "123", []tok{halfTok(GE_LITERAL, SU_NUMBER, "123")})
}

func Test_T4_3(t *testing.T) {
	doTest(t, "1.0", []tok{halfTok(GE_LITERAL, SU_NUMBER, "1.0")})
}

func Test_T4_4(t *testing.T) {
	doTest(t, "123.456", []tok{halfTok(GE_LITERAL, SU_NUMBER, "123.456")})
}

func Test_T4_5(t *testing.T) {
	doErrorTest(t, "123.")
}

func Test_T5_1(t *testing.T) {
	doTest(t, `""`, []tok{halfTok(GE_LITERAL, SU_STRING, `""`)})
}

func Test_T5_2(t *testing.T) {
	doTest(t, `"abc"`, []tok{halfTok(GE_LITERAL, SU_STRING, `"abc"`)})
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
	doTest(t, "a", []tok{halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "a")})
}

func Test_T6_2(t *testing.T) {
	doTest(t, "abc", []tok{halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "abc")})
}

func Test_T6_3(t *testing.T) {
	doTest(t, "a_b", []tok{halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "a_b")})
}

func Test_T6_4(t *testing.T) {
	doTest(t, "ab_", []tok{halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "ab_")})
}

func Test_T6_5(t *testing.T) {
	doTest(t, "_", []tok{halfTok(GE_IDENTIFIER, SU_VOID, "_")})
}

func Test_T7_1(t *testing.T) {
	doTest(t, "(", []tok{halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "(")})
}

func Test_T7_2(t *testing.T) {
	doTest(t, ")", []tok{halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")")})
}

func Test_T8_1(t *testing.T) {
	doTest(t, "@abc", []tok{halfTok(GE_SPELL, SU_UNDEFINED, "@abc")})
}

func Test_T8_2(t *testing.T) {
	doTest(t, "@abc.xyz", []tok{halfTok(GE_SPELL, SU_UNDEFINED, "@abc.xyz")})
}

func Test_T8_3(t *testing.T) {
	doTest(t, "@a.b.c.d", []tok{halfTok(GE_SPELL, SU_UNDEFINED, "@a.b.c.d")})
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
	doTest(t, ",", []tok{halfTok(GE_DELIMITER, SU_VALUE_DELIM, ",")})
}

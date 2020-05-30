package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func(Factory, []Token) ([]Statement, error) = ParseStatements

func Test_S1(t *testing.T) {

	// a

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(TERMINATOR, ""),
	}

	exp := Identifier{tok(IDENTIFIER, "a")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_S2(t *testing.T) {

	// _

	given := []Token{
		tok(VOID, "_"),
		tok(TERMINATOR, ""),
	}

	exp := Identifier{tok(VOID, "_")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_S3(t *testing.T) {

	// TRUE

	given := []Token{
		tok(BOOL, "TRUE"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(BOOL, "TRUE")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_S4(t *testing.T) {

	// 1

	given := []Token{
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(NUMBER, "1")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_S5(t *testing.T) {

	// "abc"

	given := []Token{
		tok(STRING, `"abc"`),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(STRING, `"abc"`)}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_F1(t *testing.T) {

	// :
	// Because an assignment token is never at the start of a statement

	given := []Token{
		tok(ASSIGN, ":"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

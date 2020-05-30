package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var testFunc func(Factory, []Token) ([]Statement, error) = ParseStatements

func Test_E1(t *testing.T) {

	// a

	given := []Token{
		tok(IDENTIFIER, "a"),
		tok(TERMINATOR, ""),
	}

	exp := Identifier{tok(IDENTIFIER, "a")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_E2(t *testing.T) {

	// _

	given := []Token{
		tok(VOID, "_"),
		tok(TERMINATOR, ""),
	}

	exp := Identifier{tok(VOID, "_")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_E3(t *testing.T) {

	// 1

	given := []Token{
		tok(BOOL, "TRUE"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(BOOL, "TRUE")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_E4(t *testing.T) {

	// 1

	given := []Token{
		tok(NUMBER, "1"),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(NUMBER, "1")}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_E5(t *testing.T) {

	// "abc"

	given := []Token{
		tok(STRING, `"abc"`),
		tok(TERMINATOR, ""),
	}

	exp := Literal{tok(STRING, `"abc"`)}

	act, e := testFunc(NewFactory(), given)
	expectOneStat(t, exp, act, e)
}

func Test_EF1(t *testing.T) {

	// :

	given := []Token{
		tok(ASSIGN, ":"),
		tok(TERMINATOR, ""),
	}

	act, e := testFunc(NewFactory(), given)
	expectError(t, act, e)
}

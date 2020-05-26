package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func EF1_FuncDef(t *testing.T, f ParseFunc) {

	// f: E(a, b) a + b

	given := []Token{
		tok(IDENTIFIER, "f"),
		tok(ASSIGN, ":"),
		tok(EXPR_FUNC, "E"),
		tok(PAREN_OPEN, "("),
		tok(IDENTIFIER, "a"),
		tok(DELIMITER, ","),
		tok(IDENTIFIER, "b"),
		tok(PAREN_CLOSE, ")"),
		tok(IDENTIFIER, "a"),
		tok(ADD, "+"),
		tok(IDENTIFIER, "b"),
		tok(TERMINATOR, ""),
	}

	targets := []AssignTarget{
		AssignTarget{tok(IDENTIFIER, "f"), nil},
	}

	exprFunc := ExprFuncDef{
		Key: tok(EXPR_FUNC, "E"),
		Inputs: []Token{
			tok(IDENTIFIER, "a"),
			tok(IDENTIFIER, "b"),
		},
	}

	exprFunc.Expr = Operation{
		Value{tok(IDENTIFIER, "a")},
		tok(ADD, "+"),
		Value{tok(IDENTIFIER, "b")},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":"),
		[]Expression{exprFunc},
	}

	expectOneStat(t, exp, f(given))
}

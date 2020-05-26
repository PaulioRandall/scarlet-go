package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func EF1_ExprFuncDef(t *testing.T, f ParseFunc) {

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
		Identifier{tok(IDENTIFIER, "a")},
		tok(ADD, "+"),
		Identifier{tok(IDENTIFIER, "b")},
	}

	exp := Assignment{
		false,
		targets,
		tok(ASSIGN, ":"),
		[]Expression{exprFunc},
	}

	expectOneStat(t, exp, f(given))
}

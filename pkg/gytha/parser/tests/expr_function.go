package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func EF1_ExprFuncDef(t *testing.T, f ParseFunc) {

	// f: E(a, b) a + b

	given := []Token{
		tok(TK_IDENTIFIER, "f"),
		tok(TK_ASSIGNMENT, ":"),
		tok(TK_EXPR_FUNC, "E"),
		tok(TK_PAREN_OPEN, "("),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_DELIMITER, ","),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_PAREN_CLOSE, ")"),
		tok(TK_IDENTIFIER, "a"),
		tok(TK_PLUS, "+"),
		tok(TK_IDENTIFIER, "b"),
		tok(TK_TERMINATOR, "\n"),
	}

	targets := []AssignTarget{
		AssignTarget{tok(TK_IDENTIFIER, "f"), nil},
	}

	exprFunc := ExprFuncDef{
		Key: tok(TK_EXPR_FUNC, "E"),
		Inputs: []Token{
			tok(TK_IDENTIFIER, "a"),
			tok(TK_IDENTIFIER, "b"),
		},
	}

	exprFunc.Expr = Operation{
		Identifier{tok(TK_IDENTIFIER, "a")},
		tok(TK_PLUS, "+"),
		Identifier{tok(TK_IDENTIFIER, "b")},
	}

	exp := Assignment{
		false,
		targets,
		tok(TK_ASSIGNMENT, ":"),
		[]Expression{exprFunc},
	}

	expectOneStat(t, exp, f(given))
}

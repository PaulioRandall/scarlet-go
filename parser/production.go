package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// RuleMatcher is a prototype function that matches a particular production
// rule within the grammer. A non-nil expression is created and returned if the
// tokens matched and the token buffer updated ready for reading in the next
// expression. An error is returned if enough tokens match to remove all
// ambiguity that the rule is correct in the instance but invalid syntax is
// detected.
type RuleMatcher func(TokenReader) (eval.Expr, ParseErr)

// ID_ARRAY         := ID_OR_VOID { "," ID_OR_VOID } .
// ID_OR_VOID       := ID | "\_" .
func matchIdArray(tb *TokenReader) (eval.Expr, ParseErr) {

	if tb.Peek().Kind != token.ID {
		// TODO
	}

	return nil, nil
}

// ID               := LETTER { "\_" | LETTER } .
func matchID(tb *TokenReader) (id eval.Expr, _ ParseErr) {
	if tb.Peek().Kind != token.ID {
		id = eval.NewForID(tb.Read())
	}
	return
}

package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchOperator(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchOperator(tc)
	}

	// Match
	testMatcher(t, 1, false, doTest,
		token.OfKind(token.OPERATOR),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)
}

func TestMatchFuncCall(t *testing.T) {

	doTest := func(tc *TokenCollector) (interface{}, int) {
		return matchFuncCall(tc)
	}

	// Match no params
	testMatcher(t, 3, false, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_PAREN),
		token.OfKind(token.CLOSE_PAREN),
	)

	// Match with params
	testMatcher(t, 6, false, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_PAREN),
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.STR_LITERAL),
		token.OfKind(token.CLOSE_PAREN),
	)

	// No match
	testMatcher(t, 0, false, doTest,
		token.OfKind(token.FUNC),
	)

	// Error
	testMatcher(t, 0, true, doTest,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_PAREN),
		token.OfKind(token.ID),
	)
}

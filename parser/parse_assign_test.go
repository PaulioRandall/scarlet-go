package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestMatchAssignIds_1(t *testing.T) {

	// Match single
	testMatcher(t, 1, false, matchAssignIds,
		token.OfKind(token.ID),
	)

	// Match multiple
	testMatcher(t, 5, false, matchAssignIds,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.ID),
	)

	// No match
	testMatcher(t, 0, false, matchAssignIds,
		token.OfKind(token.FUNC),
	)

	// Invalid syntax
	testMatcher(t, 0, true, matchAssignIds,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.FUNC),
	)
}

func TestMatchAssignStart_1(t *testing.T) {

	// Match
	testMatcher(t, 2, false, matchAssignStart,
		token.OfKind(token.ID),
		token.OfKind(token.ASSIGN),
	)

	// No match
	testMatcher(t, 0, false, matchAssignStart,
		token.OfKind(token.UNDEFINED),
	)

	// Error
	testMatcher(t, 0, true, matchAssignStart,
		token.OfKind(token.ID),
		token.OfKind(token.UNDEFINED),
	)
}

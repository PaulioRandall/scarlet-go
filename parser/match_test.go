package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func TestMatchSeq_1(t *testing.T) {

	tc := dummyTC(
		token.OfKind(token.ID),
		token.OfKind(token.ASSIGN),
	)

	// Match
	require.Equal(t, 2, matchSeq(tc,
		token.ID,
		token.ASSIGN,
	))

	// No match, non-matching token
	require.Equal(t, 0, matchSeq(tc,
		token.FUNC,
		token.ASSIGN,
	))

	// No match, not enough tokens in stream
	require.Equal(t, 0, matchSeq(tc,
		token.ID,
		token.ASSIGN,
		token.FUNC,
	))
}

func TestMatchAny_1(t *testing.T) {

	tc := dummyTC(
		token.OfKind(token.ID),
	)

	// Match
	require.Equal(t, 1, matchAny(tc,
		token.ASSIGN,
		token.ID,
	))

	// No match
	require.Equal(t, 0, matchAny(tc,
		token.ASSIGN,
		token.FUNC,
	))
}

func TestMatchEither_1(t *testing.T) {

	tc := dummyTC(
		token.OfKind(token.ID),
	)

	// Match
	require.Equal(t, 1, matchEither(tc,
		func(tc *TokenCollector) int { return 0 },
		func(tc *TokenCollector) int { return 1 },
	))

	// No match
	require.Equal(t, 0, matchEither(tc,
		func(tc *TokenCollector) int { return 0 },
		func(tc *TokenCollector) int { return 0 },
	))
}

func TestMatchIdOrItem_1(t *testing.T) {

	// Match ID
	testMatcher(t, 1, false, matchIdOrItem,
		token.OfKind(token.ID),
	)

	// Match item access
	testMatcher(t, 4, false, matchIdOrItem,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	)

	// No match
	testMatcher(t, 0, false, matchIdOrItem,
		token.OfKind(token.FUNC),
	)
}

func TestMatchItemAccess_1(t *testing.T) {

	// Match
	testMatcher(t, 3, false, matchItemAccess,
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	)

	// No match
	testMatcher(t, 0, false, matchItemAccess,
		token.OfKind(token.OPEN_GUARD),
		token.OfKind(token.FUNC),
	)
}

func TestMatchParam_1(t *testing.T) {

	// Match id
	testMatcher(t, 1, false, matchParam,
		token.OfKind(token.ID),
	)

	// Match literal
	testMatcher(t, 1, false, matchParam,
		token.OfKind(token.STR_LITERAL),
	)

	// Match void
	testMatcher(t, 1, false, matchParam,
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher(t, 0, false, matchParam,
		token.OfKind(token.FUNC),
	)
}

func TestMatchParamList_1(t *testing.T) {

	// Match single
	testMatcher(t, 1, false, matchParamList,
		token.OfKind(token.ID),
	)

	// Match multiple
	testMatcher(t, 5, false, matchParamList,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.STR_LITERAL),
		token.OfKind(token.DELIM),
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher(t, 0, false, matchParamList,
		token.OfKind(token.OPERATOR),
	)

	// Error
	testMatcher(t, 0, true, matchParamList,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
	)
}

func TestMatchCallStart_1(t *testing.T) {

	// Match
	testMatcher(t, 2, false, matchCallStart,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_PAREN),
	)

	// No match
	testMatcher(t, 0, false, matchCallStart,
		token.OfKind(token.FUNC),
	)
}

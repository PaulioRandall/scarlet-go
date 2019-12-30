package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func TestMatchSeq_1(t *testing.T) {

	tc := dummyTC_2(
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

	tc := dummyTC_2(
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

	tc := dummyTC_2(
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

func TestMatchIdArray_2_1(t *testing.T) {

	// Match single
	testMatcher_2(t, 1, false, matchIdArray_2,
		token.OfKind(token.ID),
	)

	// Match multiple
	testMatcher_2(t, 5, false, matchIdArray_2,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.ID),
	)

	// No match
	testMatcher_2(t, 0, false, matchIdArray_2,
		token.OfKind(token.FUNC),
	)

	// Invalid syntax
	testMatcher_2(t, 0, true, matchIdArray_2,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.FUNC),
	)
}

func TestMatchIdOrItem_2_1(t *testing.T) {

	// Match ID
	testMatcher_2(t, 1, false, matchIdOrItem_2,
		token.OfKind(token.ID),
	)

	// Match item access
	testMatcher_2(t, 4, false, matchIdOrItem_2,
		token.OfKind(token.ID),
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	)

	// No match
	testMatcher_2(t, 0, false, matchIdOrItem_2,
		token.OfKind(token.FUNC),
	)
}

func TestMatchItemAccess_2_1(t *testing.T) {

	// Match
	testMatcher_2(t, 3, false, matchItemAccess_2,
		token.OfKind(token.OPEN_GUARD),
		token.OfValue(token.INT_LITERAL, "123"),
		token.OfKind(token.CLOSE_GUARD),
	)

	// No match
	testMatcher_2(t, 0, false, matchItemAccess_2,
		token.OfKind(token.OPEN_GUARD),
		token.OfKind(token.FUNC),
	)
}

func TestMatchParam_2_1(t *testing.T) {

	// Match id
	testMatcher_2(t, 1, false, matchParam_2,
		token.OfKind(token.ID),
	)

	// Match literal
	testMatcher_2(t, 1, false, matchParam_2,
		token.OfKind(token.STR_LITERAL),
	)

	// Match void
	testMatcher_2(t, 1, false, matchParam_2,
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher_2(t, 0, false, matchParam_2,
		token.OfKind(token.FUNC),
	)
}

func TestMatchParamList_2_1(t *testing.T) {

	// Match single
	testMatcher_2(t, 1, false, matchParamList_2,
		token.OfKind(token.ID),
	)

	// Match multiple
	testMatcher_2(t, 5, false, matchParamList_2,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.STR_LITERAL),
		token.OfKind(token.DELIM),
		token.OfKind(token.VOID),
	)

	// No match
	testMatcher_2(t, 0, false, matchParamList_2,
		token.OfKind(token.OPERATOR),
	)

	// Error
	testMatcher_2(t, 0, true, matchParamList_2,
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
	)
}

func TestMatchCall_2_1(t *testing.T) {

	// Match no params
	testMatcher_2(t, 2, false, matchCall_2,
		token.OfKind(token.OPEN_PAREN),
		token.OfKind(token.CLOSE_PAREN),
	)

	// Match with params
	testMatcher_2(t, 5, false, matchCall_2,
		token.OfKind(token.OPEN_PAREN),
		token.OfKind(token.ID),
		token.OfKind(token.DELIM),
		token.OfKind(token.STR_LITERAL),
		token.OfKind(token.CLOSE_PAREN),
	)

	// No match
	testMatcher_2(t, 0, false, matchCall_2,
		token.OfKind(token.FUNC),
	)

	// Error
	testMatcher_2(t, 0, true, matchCall_2,
		token.OfKind(token.OPEN_PAREN),
		token.OfKind(token.ID),
	)
}

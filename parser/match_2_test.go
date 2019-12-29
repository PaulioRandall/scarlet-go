package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func TestMatch_1(t *testing.T) {

	tc := dummyTC_2(
		token.OfKind(token.ID),
		token.OfKind(token.ASSIGN),
		token.OfKind(token.BOOL_LITERAL),
	)

	act := match(tc,
		token.ID,
		token.ASSIGN,
	)

	require.Equal(t, 2, act)
}

func TestMatch_2(t *testing.T) {

	tc := dummyTC_2(
		token.OfKind(token.ID),
		token.OfKind(token.ASSIGN),
	)

	act := match(tc,
		token.FUNC,
		token.ASSIGN,
	)

	require.Equal(t, 0, act)
}

func TestMatch_3(t *testing.T) {

	tc := dummyTC_2(
		token.OfKind(token.ID),
		token.OfKind(token.ASSIGN),
	)

	act := match(tc,
		token.ID,
		token.ASSIGN,
		token.BOOL_LITERAL,
	)

	require.Equal(t, 0, act)
}

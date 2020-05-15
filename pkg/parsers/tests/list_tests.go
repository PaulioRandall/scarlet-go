package tests

import (
	"testing"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func L1_ListDef(t *testing.T, f ParseFunc) {

	// LIST {1,2,3}

	given := []Token{
		Token{LIST, "LIST", 0, 0},
		Token{BLOCK_OPEN, "{", 0, 0},
		Token{NUMBER, "1", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{NUMBER, "2", 0, 0},
		Token{DELIM, ",", 0, 0},
		Token{NUMBER, "3", 0, 0},
		Token{BLOCK_CLOSE, "}", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{EOF, "", 0, 0},
	}

	exprs := []st.Expression{
		st.Value(Token{NUMBER, "1", 0, 0}),
		st.Value(Token{NUMBER, "2", 0, 0}),
		st.Value(Token{NUMBER, "3", 0, 0}),
	}

	list := st.List{
		Key:   Token{LIST, "LIST", 0, 0},
		Open:  Token{BLOCK_OPEN, "{", 0, 0},
		Exprs: exprs,
		Close: Token{BLOCK_CLOSE, "}", 0, 0},
	}

	expectOneStat(t, list, f(given))
}

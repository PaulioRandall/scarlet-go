package strimmer

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

func TestWrap_1(t *testing.T) {
	lexor.ScanTokenTest(t,
		New(lexor.DummyScanner([]token.Token{
			token.NewToken(token.ID, "abc", 0, 0, 3),
		})),
		token.NewToken(token.ID, "abc", 0, 0, 3),
	)
}

func TestWrap_2(t *testing.T) {
	lexor.ScanTokenTest(t,
		New(lexor.DummyScanner([]token.Token{
			token.NewToken(token.ID, "abc", 0, 0, 3),
			token.NewToken(token.NEWLINE, "\n", 0, 3, 4),
			token.NewToken(token.WHITESPACE, "   ", 1, 0, 3),
			token.NewToken(token.ID, "efg", 1, 3, 6),
			token.NewToken(token.COMMENT, "// xyz", 1, 6, 12),
			token.NewToken(token.NEWLINE, "\n", 1, 12, 13),
		})),
		token.NewToken(token.ID, "abc", 0, 0, 3),
		token.NewToken(token.NEWLINE, "\n", 0, 3, 4),
		token.NewToken(token.ID, "efg", 1, 3, 6),
		token.NewToken(token.NEWLINE, "\n", 1, 12, 13),
	)
}

func TestWrap_3(t *testing.T) {
	lexor.ScanTokenTest(t,
		New(lexor.DummyScanner([]token.Token{
			token.NewToken(token.ID, "\t\t\t", 0, 0, 3),
		})),
	)

	lexor.ScanTokenTest(t,
		New(lexor.DummyScanner([]token.Token{
			token.NewToken(token.COMMENT, "// abc", 0, 0, 6),
		})),
	)
}

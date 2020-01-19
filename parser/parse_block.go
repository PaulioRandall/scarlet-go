package parser

/*
import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseBlock parses a block of statements.
func parseBlock(tm *TokenMatcher) (_ []Expr) {

	var stats []Expr

	for 0 == tm.MatchAny(token.END, token.UNDEFINED) {
		if ex := parseStatement(tm); ex != nil {
			stats = append(stats, ex)
		}
	}

	tm.Skip()
	return stats
}

// parseStatement parses a statement.
func parseStatement(tm *TokenMatcher) (_ Expr) {

	switch {
	case 1 == tm.Match(token.NEWLINE):
	case 1 == tm.MatchSeq(token.ID, token.OPEN_PAREN):
	case 1 == tm.Match(token.ID):
		return parseAssign(tm)
	}

	// TODO: Other possibile expressions

	return
}
*/

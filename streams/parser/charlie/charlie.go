// charlie package parses the assignment tokens within BetaStatements to
// produce CharlieStatements.
package charlie

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// CharlieStatement represents an partially parsed statement where the
// assignment tokens have been parsed but the expressions tokens have not.
//
// E.g.
// Consider `x := 1 + 1`:
// - the whole thing is a statement
// - `:=` remains the assignment token (Assign)
// - `x` remains the identifier token (IDs)
// - `1 + 1` remains the expression tokens (Exprs)
//
// Now consider `a, b := 1 + 2, 4 - 3`:
// - the whole thing is a statement
// - `:=` remains the assignment token (Assign)
// - `a, b` will become []Token{`a`, `b`} (IDs)
// - `1 + 2, 4 - 3` remains the expression tokens (Exprs)
type CharlieStatement struct {
	Assign lexeme.Token
	IDs    []lexeme.Token
	Exprs  []lexeme.Token
	Subs   []CharlieStatement
}

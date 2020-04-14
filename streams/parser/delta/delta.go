// delta package parses the assignment tokens within CharlieStatements to
// produce DeltaStatements.
package delta

import (
	//"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"
	//	"github.com/PaulioRandall/scarlet-go/streams/parser/charlie"
)

// DeltaStatement represents an partially parsed statement where the
// assignment and expression tokens have been parsed but are yet to be placed
// into a parse tree.
type DeltaStatement struct {
	Assign lexeme.Token
	IDs    []lexeme.Token
	Exprs  []lexeme.Token
	Subs   []DeltaStatement
}

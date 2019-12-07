package stat

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

type Statement interface {

	// Tokens returns the tokens that make up the statement but not those of
	// sub-statements.
	Tokens() []token.Token

	// Statements returns any sub-statements but not those of any sub-statements.
	Statements() []Statement

	// Kind returns the type of the statement.
	Kind() StatKind
}

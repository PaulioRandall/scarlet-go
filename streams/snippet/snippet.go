package snippet

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// Snippet represents an unparsed statement.
type Snippet struct {
	Kind    Kind
	Tokens  []lexeme.Token
	Snippet []Snippet
}

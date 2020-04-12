package snippet

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// Snippet represents an unparsed statement.
type Snippet struct {
	Tokens   []lexeme.Token
	Snippets []Snippet
}

func (base *Snippet) appendTokens(tks ...lexeme.Token) {
	base.Tokens = append(base.Tokens, tks...)
}

func (base *Snippet) appendSnippets(snips ...Snippet) {
	base.Snippets = append(base.Snippets, snips...)
}

package parser

import (
	"github.com/PaulioRandall/scarlet-go/number"
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// Pattern: IDENT
func expectIdent(l lexeme.Lexeme) (id Ident, e error) {

	if l.Token != token.IDENT {
		e = errSnip(l.Snippet, "Expected identifier")
		return
	}

	id = Ident{
		Snippet: l.Snippet,
		Val:     l.Val,
	}
	return
}

// Pattern: BOOL
func boolLit(l lexeme.Lexeme) BoolLit {
	return BoolLit{
		Snippet: l.Snippet,
		Val:     l.Token == token.TRUE,
	}
}

// Pattern: NUMBER
func numLit(l lexeme.Lexeme) NumLit {
	return NumLit{
		Snippet: l.Snippet,
		Val:     number.New(l.Val),
	}
}

// Pattern: STRING
func strLit(l lexeme.Lexeme) StrLit {
	return StrLit{
		Snippet: l.Snippet,
		Val:     l.Val,
	}
}

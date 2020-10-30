package parser2

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"
)

// Pattern: IDENT
func expectIdent(l token.Lexeme) (id tree.Ident, e error) {

	if l.Token != token.IDENT {
		e = errSnip(l.Snippet, "Expected identifier")
		return
	}

	id = tree.Ident{
		Snippet: l.Snippet,
		Val:     l.Val,
	}
	return
}

// Pattern: BOOL
func boolLit(l token.Lexeme) tree.BoolLit {
	return tree.BoolLit{
		Snippet: l.Snippet,
		Val:     l.Token == token.TRUE,
	}
}

// Pattern: NUMBER
func numLit(l token.Lexeme) tree.NumLit {
	return tree.NumLit{
		Snippet: l.Snippet,
		Val:     number.New(l.Val),
	}
}

// Pattern: STRING
func strLit(l token.Lexeme) tree.StrLit {
	return tree.StrLit{
		Snippet: l.Snippet,
		Val:     l.Val,
	}
}

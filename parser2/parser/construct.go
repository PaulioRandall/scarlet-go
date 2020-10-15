package parser

import (
	"github.com/PaulioRandall/scarlet-go/number"
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
)

// Pattern: IDENT
func expectIdent(l lexeme.Lexeme) (id tree.Ident, e error) {

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
func boolLit(l lexeme.Lexeme) tree.BoolLit {
	return tree.BoolLit{
		Snippet: l.Snippet,
		Val:     l.Token == token.TRUE,
	}
}

// Pattern: NUMBER
func numLit(l lexeme.Lexeme) tree.NumLit {
	return tree.NumLit{
		Snippet: l.Snippet,
		Val:     number.New(l.Val),
	}
}

// Pattern: STRING
func strLit(l lexeme.Lexeme) tree.StrLit {
	return tree.StrLit{
		Snippet: l.Snippet,
		Val:     l.Val,
	}
}

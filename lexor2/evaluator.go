package lexor

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// ****************************************************************************
// * Package API
// ****************************************************************************

// Evaluator is a structure for parsed source code tokens. This involves
// removing quotes from string literals and discarding whitespace tokens etc.
// It requires a TokenStream as a source for tokens and implements the
// TokenStream interface so it may be wrapped.
type Evaluator struct {
	ts TokenStream
}

// NewEvaluator creates a new evaluator to evaluate tokens within a stream.
func NewEvaluator(delegate TokenStream) TokenStream {
	return &Evaluator{
		ts: delegate,
	}
}

// Next satisfies the TokenStream interface.
func (ev *Evaluator) Next() (_ token.Token) {

	var tk token.Token
	var k token.Kind

	for tk = ev.ts.Next(); tk != token.ZERO(); tk = ev.ts.Next() {
		k = tk.Kind

		if k == token.WHITESPACE || k == token.COMMENT {
			continue
		}

		if k == token.STR_LITERAL || k == token.STR_TEMPLATE {
			trimStrQuotes(&tk)
		}

		return tk
	}

	return
}

// ****************************************************************************
// * Helper functions
// ****************************************************************************

// trimStrQuotes removes the leading and trailing quotes from string literals
// and templates.
func trimStrQuotes(tk *token.Token) {
	s := tk.Value
	tk.Value = s[1 : len(s)-1]
}

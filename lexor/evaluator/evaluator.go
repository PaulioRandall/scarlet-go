// evaluator performs evaluation on tokens obtained via an underlying
// ScanToken function, i.e. removing insignificant tokens, quotes from string
// literals, ect.
package evaluator

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New makes a new evaluator function that wraps the specified ScanToken
// function.
func New(src lexor.ScanToken) lexor.ScanToken {
	return wrap(src)
}

// wrap wraps a ScanToken with one that performs evaluation. The underlying
// ScanToken function maybe called multiple times before a value is returned.
func wrap(f lexor.ScanToken) lexor.ScanToken {

	if f == nil {
		return nil
	}

	return func() (_ token.Token, _ lexor.ScanToken, e lexor.ScanErr) {

		var t token.Token
		var st lexor.ScanToken

		for st = f; st != nil; {

			t, st, e = f()

			switch {
			case e != nil:
				return
			case t == (token.Token{}):
				e = lexor.NewScanErr(
					"SANITY CHECK: Unexpected EMPTY token", nil, 0, 0)
				return
			case t.Kind == token.UNDEFINED:
				e = lexor.NewScanErr(
					"SANITY CHECK: Unexpected UNDEFINED token", nil, 0, 0)
				return
			case eval(&t):
				return t, wrap(st), nil
			}
		}

		return
	}
}

// eval evaluates the value of the token if needed and then returns it.
func eval(t *token.Token) bool {
	switch t.Kind {
	case token.WHITESPACE, token.COMMENT:
		return false
	case token.STR_LITERAL, token.STR_TEMPLATE:
		trimStrQuotes(t)
	}

	return true
}

// trimStrQuotes removes the leading and trailing quotes from string literals
// and templates.
func trimStrQuotes(t *token.Token) {
	s := t.Value
	t.Value = s[1 : len(s)-1]
}

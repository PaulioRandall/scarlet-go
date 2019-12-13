// evaluator parses the value of certain tokens, i.e. removing quotes from
// string literals. It wraps another ScanToken function.
package evaluator

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New makes a new evaluator function.
func New(src lexor.ScanToken) lexor.ScanToken {
	return wrap(src)
}

// wrap wraps a ScanToken with one that evaluates the token before being
// returned.
func wrap(f lexor.ScanToken) lexor.ScanToken {

	if f == nil {
		return nil
	}

	return func() (t token.Token, st lexor.ScanToken, e lexor.ScanErr) {

		t, st, e = f()

		if e == nil && t != nil {
			t = eval(t)
			st = wrap(st)
		}

		return
	}
}

// eval evaluates the value of the token if needed and then returns it.
func eval(t token.Token) token.Token {
	if t.Kind() == token.STR_LITERAL {
		t = evalStr(t)
	}

	return t
}

// evalStr evaluates the value of a string literal by removing the leading and
// trailing quotes.
func evalStr(t token.Token) token.Token {
	s := t.Value()
	s = s[1 : len(s)-1]
	return token.TokenBySnippet(t.Kind(), s, t.Where())
}

package evaluator

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/lexor/strimmer"
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) lexor.ScanToken {
	st := strimmer.New(src)
	return wrap(st)
}

// wrap wraps a ScanToken function with one that iterates through the strimmer
// to find and evaluate tokens that require evaluation.
func wrap(f lexor.ScanToken) lexor.ScanToken {

	if f == nil {
		return nil
	}

	return func() (t token.Token, st lexor.ScanToken, e perror.Perror) {

		t, st, e = f()

		if e == nil && t != nil {
			t = evaluate(t)
			st = wrap(st)
		}

		return
	}
}

// evaluate evaluates the value of the token if needed and then returns it.
func evaluate(t token.Token) token.Token {
	if t.Kind() == token.STR_LITERAL {
		t = evaluateStrLiteral(t)
	}

	return t
}

// evaluates the value of a string literal by removing the leading and trailing
// back tik.
func evaluateStrLiteral(t token.Token) token.Token {
	s := t.Value()
	s = s[1 : len(s)-1]
	return token.Newish(t.Kind(), s, t.Where())
}

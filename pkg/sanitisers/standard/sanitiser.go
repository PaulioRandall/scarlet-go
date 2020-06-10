package standard

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

// SanitiseAll removes redundant tokens, such as comment and whitespace, as well
// as applying formatting to values, e.g trimming off the quotes from string
// literals and templates.
func SanitiseAll(in []Token) (out []Token) {

	var prev Token
	size := len(in)

	for i := 0; i < size; i++ {

		tk := sanitise(prev, in[i])

		if tk != nil {
			out = append(out, tk)
			prev = tk
		}
	}

	return out
}

func sanitise(prev, next Token) Token {

	if isParsableToken(prev, next) {
		return formatToken(next)
	}

	return nil
}

func isParsableToken(prev, next Token) bool {

	if next.Type().Redundant() {
		return false
	}

	if next.Type() == TK_NEWLINE || next.Type() == TK_TERMINATOR {

		if prev == nil {
			return false
		}

		switch prev.Type() {
		// Sometimes the extra terminator or newline is redundant.
		// Removing them makes parsing easier.
		case TK_TERMINATOR,
			TK_DELIMITER,
			TK_BLOCK_OPEN,
			TK_BLOCK_CLOSE,
			TK_WHEN,
			TK_LIST,
			TK_UNDEFINED:

			return false
		}
	}

	return true
}

func formatToken(tk Token) Token {

	switch tk.Type() {
	case TK_NEWLINE:
		// Non-redundant newline tokens are expression and statement terminators
		// in disguise.
		return NewToken(TK_TERMINATOR, tk.Value(), tk.Line(), tk.Col())

	case TK_STRING:
		v := tk.Value()
		v = v[1 : len(v)-1] // Remove quotes
		return NewToken(tk.Type(), v, tk.Line(), tk.Col())
	}

	return tk
}

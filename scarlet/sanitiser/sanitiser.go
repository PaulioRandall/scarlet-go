// Sanitise package removes tokens redundant or inconvenient to parsing.
// Traditionally, this process is performed during the scanning process but
// has been decoupled for flexbility.
package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

// Sanitise removes Tokens redundant to parsing. In addition to whitespace
// and comments, other elements are removed such as empty statements and
// inconvenient linefeeds.
func Sanitise(tks []token.Lexeme) []token.Lexeme {

	itr := token.NewLexItr(tks)
	r := make([]token.Lexeme, 0, len(tks))

	for prevIdx := -1; itr.More(); prevIdx = len(r) - 1 {
		curr := itr.Next()

		if prevIdx < 0 {
			r = append(r, curr)
			continue
		}

		switch prev := r[prevIdx]; {
		case prev.Token == token.DELIM && curr.Token == token.R_PAREN:
			r[prevIdx] = curr // Remove delimiters ',' before right paren ')'

		case prev.IsTerminator() && curr.Token == token.R_CURLY:
			r[prevIdx] = curr // Remove terminators ',' before closing curly brace '}'

		case prev.IsOpener() && curr.Token == token.NEWLINE:
			// Remove newlines after openers '(,[,{'

		case prev.Token == token.DELIM && curr.Token == token.NEWLINE:
			// Remove newlines after delimiters ','

		default:
			r = append(r, curr)
		}
	}

	return r
}

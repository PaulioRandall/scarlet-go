// Sanitise package removes tokens redundant or inconvenient to parsing.
// Traditionally, this process is performed during the scanning process but by
// decoupling the sanitisation from scanning the scanner can be reused in
// source code formatting and analysis tools.
package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

// LexemeIterator specifies the iteration functionality required by inputs to
// the sanitisation process.
type LexemeIterator interface {
	JumpToStart()
	More() bool
	Next() token.Lexeme
	Prev() token.Lexeme     // Steps back one place in the iterator
	LookBack() token.Lexeme // Peeks at the previous Lexeme
	Remove() token.Lexeme
}

// SanitiseAll removes Tokens redundant to parsing. In addition to whitespace
// and comments, other elements are removed such as empty statements and
// inconvenient linefeeds.
func SanitiseAll(itr LexemeIterator) {

	zero := token.UNDEFINED

	itr.JumpToStart()
	for itr.More() {
		curr := itr.Next().Token
		prev := itr.LookBack().Token

		switch {
		case curr.IsRedundant():
			itr.Remove() // Always remove tokens redundant to the parsing process.

		case prev == zero && curr.IsTerminator():
			itr.Remove() // Remove leading terminators

		case prev == zero:
			// No action for the first token

		case removeAfterPrev(prev, curr):
			itr.Remove()

		case removeBeforeNext(prev, curr):
			itr.Prev()
			itr.Remove()
			itr.Next()
		}
	}
}

func removeAfterPrev(prev, curr token.Token) bool {
	switch {
	case prev.IsTerminator() && curr.IsTerminator():
		return true // Remove successive terminators

	case prev.IsOpener() && curr == token.NEWLINE:
		return true // Remove newlines after openers '(,[,{'

	case prev == token.DELIM && curr == token.NEWLINE:
		return true // Remove newlines after delimiters ','
	}

	return false
}

func removeBeforeNext(curr, next token.Token) bool {
	switch {
	case curr == token.DELIM && next == token.R_PAREN:
		return true // Remove delimiters ',' before right paren ')'

	case curr.IsTerminator() && next == token.R_CURLY:
		return true // Remove terminators ',' before closing curly brace '}'
	}

	return false
}

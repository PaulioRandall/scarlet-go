package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

type LexemeIterator interface {
	JumpToStart()
	More() bool
	Next() lexeme.Lexeme
	Prev() lexeme.Lexeme
	LookBack() lexeme.Lexeme
	Remove() lexeme.Lexeme
}

// SanitiseAll removes tokens redundant to parsing. Traditionally, this process
// is performed during the scanning process but by decoupling the sanitisation
// process the scanner become more flexible and can be reused in source code
// formatting or by analysis tools.
func SanitiseAll(itr LexemeIterator) {

	ZERO := token.UNDEFINED

	itr.JumpToStart()
	for itr.More() {
		curr := itr.Next().Token
		prev := itr.LookBack().Token

		switch {
		case curr.IsRedundant():
			itr.Remove() // Always remove tokens redundant to the parsing process.

		case prev == ZERO && curr.IsTerminator():
			itr.Remove() // Remove leading terminators

		case prev == ZERO:
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

package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type Iterator interface {
	More() bool
	IsFirst() bool
	Next() lexeme.Lexeme
	Prev() lexeme.Lexeme
	LookBack() lexeme.Lexeme
	Remove() lexeme.Lexeme
}

// SanitiseAll removes tokens redundant to parsing so it's easier and less
// error prone.
func SanitiseAll(itr Iterator) {
	for itr.More() {
		switch curr := itr.Next().Type(); {
		case curr.IsRedundant():
			itr.Remove() // Remove tokens always redundant to the compiler

		case itr.IsFirst() && curr.IsTerminator():
			itr.Remove() // Remove leading terminators

		case itr.IsFirst():
			// No action for the first token

		case removeAfterPrev(itr.LookBack().Type(), curr):
			itr.Remove()

		case removeBeforeNext(itr.LookBack().Type(), curr):
			itr.Prev()
			itr.Remove()
			itr.Next()
		}
	}
}

func removeAfterPrev(prev, curr lexeme.TokenType) bool {
	switch {
	case prev.IsTerminator() && curr.IsTerminator():
		return true // Remove successive terminators

	case prev.IsOpener() && curr == lexeme.NEWLINE:
		return true // Remove newlines after openers '(,[,{'

	case prev == lexeme.DELIM && curr == lexeme.NEWLINE:
		return true // Remove newlines after delimiters ','
	}

	return false
}

func removeBeforeNext(curr, next lexeme.TokenType) bool {
	switch {
	case curr == lexeme.DELIM && next == lexeme.R_PAREN:
		return true // Remove delimiters ',' before right paren ')'

	case curr.IsTerminator() && next == lexeme.R_CURLY:
		return true // Remove terminators ',' before closing curly brace '}'
	}

	return false
}

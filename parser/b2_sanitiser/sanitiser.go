package sanitiser2

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type Iterator interface {
	HasNext() bool
	HasPrev() bool
	Next() lexeme.Lexeme
	LookBehind() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
	Remove() lexeme.Lexeme
}

// SanitiseAll removes tokens redundant to parsing so it's easier and less
// error prone.
func SanitiseAll(itr Iterator) {
	for itr.HasNext() {
		switch curr := itr.Next().Type(); {
		case curr.IsRedundant():
			itr.Remove() // Remove tokens always redundant to the compiler

		case !itr.HasPrev() && curr.IsTerminator():
			itr.Remove() // Remove leading terminators

		case itr.HasPrev() && removeAfterPrev(itr.LookBehind().Type(), curr):
			itr.Remove()

		case itr.HasNext() && removeBeforeNext(curr, itr.LookAhead().Type()):
			itr.Remove()
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
		return true // Remove delimiters ',' before closing paren ')'

	case curr.IsTerminator() && next == lexeme.R_CURLY:
		return true // Remove terminators ',' before closing curly brace '}'
	}

	return false
}

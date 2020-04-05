package lexor

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Collection of terminal and non-terminal symbols for ease of reference and to
// speed up syntax experimentation.
const (
	keyword_function         string = `F`
	keyword_normalblockStart string = `DO`
	keyword_matchBlockStart  string = `MATCH`
	keyword_blockEnd         string = `END`
	keyword_true             string = `TRUE`
	keyword_false            string = `FALSE`
	terminal_carriageReturn  rune   = '\r'
	terminal_lineFeed        rune   = '\n'
	terminal_commentStart    rune   = '/'
	terminal_fractionalDelim rune   = '.'
	terminal_stringStart     rune   = '`'
	terminal_stringEnd       rune   = '`'
	terminal_templateStart   rune   = '"'
	terminal_templateEnd     rune   = '"'
	terminal_templateEscape  rune   = '\\'
	terminal_wordUnderscore  rune   = '_'
	nonTerminal_assignment   string = `:=`
)

// identifyKeyword identifies the kind of token the keyword (non-terminal
// symbol) represents.
func identifyKeyword(nonTerminal string) token.Kind {

	switch nonTerminal {
	case keyword_function:
		return token.FUNC
	case keyword_normalblockStart:
		return token.DO
	case keyword_matchBlockStart:
		return token.MATCH
	case keyword_blockEnd:
		return token.END
	case keyword_true, keyword_false:
		return token.BOOL
	}

	return token.UNDEFINED
}

// empty returns true if the scanners rune slice is empty.
func (s *scanner) empty() bool {
	return len(s.runes) == 0
}

// len returns the length of the scanners rune slice.
func (s *scanner) len() int {
	return len(s.runes)
}

// matches returns true if the scanners rune slice matches the specified
// sequence of runes.
func (s *scanner) matches(start int, terminals ...rune) bool {

	haystack := s.runes[start:]

	if len(terminals) > len(haystack) {
		return false
	}

	for i, ru := range terminals {
		if haystack[i] != ru {
			return false
		}
	}

	return true
}

// doesNotMatch returns false if the scanners rune slice matches the specified
// sequence of runes.
func (s *scanner) doesNotMatch(start int, terminals ...rune) bool {
	return !s.matches(start, terminals...)
}

// matchesNewline returns true if the scanners rune slice begins with a sequence
// of newline terminals.
func (s *scanner) matchesNewline(start int) bool {
	return s.countNewlineTerminals(start) > 0
}

// noMatchNewline returns false if the scanners rune slice begins with a
// sequence of newline terminals.
func (s *scanner) doesNotMatchNewline(start int) bool {
	return !s.matchesNewline(start)
}

// matchUntil iterates the scanners rune slice executing the function for
// on each rune. The number of runes counted before the function results in true
// is returned to the user. If the function never returns true then the length
// of the rune slice, from the start index, is returned.
func (s *scanner) matchUntil(start int, f func(int, rune) bool) int {

	var i int
	var ru rune

	for i, ru = range s.runes[start:] {
		if f(i, ru) {
			break
		}
	}

	return i
}

// matchUntilNewline returns the count of runes from the beginning of the
// scanners rune slice to the first newline (exclusive).
func (s *scanner) matchUntilNewline(start int) int {
	return s.matchUntil(start, func(i int, ru rune) bool {
		return s.matchesNewline(i)
	})
}

// countNewlineTerminals returns the number of newline terminal symbols that
// make up the next token in the scanners rune slice. Zero is returned if the
// next sequence of terminals don't represent a newline.
func (s *scanner) countNewlineTerminals(start int) int {

	const (
		NONE int = 0
		LF   int = 1
		CRLF int = 2
	)

	r := s.runes[start:]
	size := len(r)

	switch {
	case size < 1:
		return NONE
	case r[0] == terminal_lineFeed:
		return LF
	case size == 1:
		return NONE
	case r[0] == terminal_carriageReturn && r[1] == terminal_lineFeed:
		return CRLF
	}

	return NONE
}

package lexor

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Collection of terminal and non-terminal symbols for ease of reference and to
// speed up syntax experimentation.
const (
	terminal_carriageReturn         rune   = '\r'
	terminal_lineFeed               rune   = '\n'
	terminal_commentStart           rune   = '/'
	terminal_fractionalDelim        rune   = '.'
	terminal_stringStart            rune   = '`'
	terminal_stringEnd              rune   = '`'
	terminal_templateStart          rune   = '"'
	terminal_templateEnd            rune   = '"'
	terminal_templateEscape         rune   = '\\'
	terminal_wordUnderscore         rune   = '_'
	nonTerminal_assignment          string = `:=`
	nonTerminal_returnParams        string = `->`
	terminal_openParen              rune   = '('
	terminal_closeParen             rune   = ')'
	terminal_openGuard              rune   = '['
	terminal_closeGuard             rune   = ']'
	terminal_openList               rune   = '{'
	terminal_closeList              rune   = '}'
	terminal_delim                  rune   = ','
	terminal_void                   rune   = '_'
	terminal_terminator             rune   = ';'
	termianl_spellPrefix            rune   = '@'
	terminal_universalNegation      rune   = '~'
	terminal_teaDrinkingNegation    rune   = '¬'
	terminal_mathAddition           rune   = '+'
	terminal_mathSubtraction        rune   = '-'
	terminal_mathMultiplication     rune   = '*'
	terminal_mathDivision           rune   = '/'
	terminal_mathRemainder          rune   = '%'
	terminal_logicalAnd             rune   = '&'
	terminal_logicalOr              rune   = '|'
	terminal_equality               rune   = '='
	terminal_inEquality             rune   = '#'
	nonTerminal_lessThanOrEquals    string = "<="
	nonTerminal_greaterThanOrEquals string = "=>"
	terminal_lessThan               rune   = '<'
	terminal_moreThan               rune   = '>'
)

// identifyKeyword identifies the kind of token the keyword (non-terminal
// symbol) represents.
func identifyKeyword(nonTerminal string) token.Kind {

	switch nonTerminal {
	case token.NON_TERMINAL_FUNCTION:
		return token.FUNC
	case token.NON_TERMINAL_NORMAL_BLOCK_START:
		return token.DO
	case token.NON_TERMINAL_MATCH_BLOCK_START:
		return token.MATCH
	case token.NON_TERMINAL_BLOCK_END:
		return token.END
	case token.NON_TERMINAL_TRUE, token.NON_TERMINAL_FALSE:
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
func (s *scanner) matchesTerminal(i int, terminal rune) bool {

	if i >= len(s.runes) {
		panic("SANITY CHECK! Bad argument, start is bigger than the remaining source code")
	}

	return s.runes[i] == terminal
}

// matches returns true if the scanners rune slice matches the specified
// string.
func (s *scanner) matchesNonTerminal(start int, needle string) bool {

	haystack := s.runes[start:]

	if len(needle) > len(haystack) {
		panic("SANITY CHECK! Bad argument, the `needle` is bigger than the `haystack`")
	}

	for i, ru := range needle {
		if haystack[i] != ru {
			return false
		}
	}

	return true
}

// matches returns true if the scanners rune slice matches the specified
// sequence of runes.
// @Deprecated
func (s *scanner) matchesSequence(start int, terminals ...rune) bool {

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
	return !s.matchesSequence(start, terminals...)
}

// matchesNewline returns true if the scanners rune slice begins with a sequence
// of newline terminals.
func (s *scanner) matchesNewline(start int) bool {
	return s.howManyNewlineTerminals(start) > 0
}

// noMatchNewline returns false if the scanners rune slice begins with a
// sequence of newline terminals.
func (s *scanner) doesNotMatchNewline(start int) bool {
	return !s.matchesNewline(start)
}

// howManyRunesUntil iterates the scanners rune slice executing the function for
// on each rune. The number of runes counted before the function results in true
// is returned to the user. If the function never returns true then the length
// of the rune slice, from the start index, is returned.
func (s *scanner) howManyRunesUntil(start int, f func(int, rune) bool) int {

	var i int
	var ru rune

	for i, ru = range s.runes[start:] {
		if f(i, ru) {
			break
		}
	}

	return i
}

// countUntilNewline returns the count of runes from the beginning of the
// scanners rune slice to the first newline (exclusive).
func (s *scanner) howManyRunesUntilNewline(start int) int {
	return s.howManyRunesUntil(start, func(i int, ru rune) bool {
		return s.matchesNewline(i)
	})
}

// countNewlineTerminals returns the number of newline terminal symbols that
// make up the next token in the scanners rune slice. Zero is returned if the
// next sequence of terminals don't represent a newline.
func (s *scanner) howManyNewlineTerminals(start int) int {

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

// tokenize slices off the next token from the scanners rune array and updates
// the line and column numbers accordingly.
func (s *scanner) tokenize(n int, k token.Kind, newline bool) (tk token.Token) {

	if s.len() < n {
		panic("SANITY CHECK! Bad argument, n is bigger than the remaining source code")
	}

	tk = token.Token{
		Kind:  k,
		Value: string(s.runes[:n]),
		Line:  s.line,
		Col:   s.col,
	}

	s.runes = s.runes[n:]

	if newline {
		s.line++
		s.col = 0
	} else {
		s.col += n
	}

	return
}

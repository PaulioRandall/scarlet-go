package lexor

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// symbolStream provides access to an ordered stream of terminal symbols. The
// word 'symbol' and Go type 'rune' are used interchangable.
type symbolStream struct {
	runes []rune // Symbols representing a script
	line  int    // Current line index within the script
	col   int    // Current column index within the line
}

// empty returns true if the scanners rune slice is empty.
func (ss *symbolStream) empty() bool {
	return len(ss.runes) == 0
}

// len returns the length of the scanners rune slice.
func (ss *symbolStream) len() int {
	return len(ss.runes)
}

// matches returns true if the scanners rune slice matches the specified
// sequence of runes.
func (ss *symbolStream) matchesTerminal(i int, terminal rune) bool {

	if i >= len(ss.runes) {
		panic("SANITY CHECK! Bad argument, start is bigger than the remaining source code")
	}

	return ss.runes[i] == terminal
}

func (ss *symbolStream) matchesNonTerminal(start int, needle string) bool {

	haystack := ss.runes[start:]

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

func (ss *symbolStream) doesNotMatchNonTerminal(start int, needle string) bool {
	return !ss.matchesNonTerminal(start, needle)
}

// matchesNewline returns true if the scanners rune slice begins with a sequence
// of newline terminals.
func (ss *symbolStream) matchesNewline(start int) bool {
	return ss.countNewlineRunes(start) > 0
}

// noMatchNewline returns false if the scanners rune slice begins with a
// sequence of newline terminals.
func (ss *symbolStream) doesNotMatchNewline(start int) bool {
	return !ss.matchesNewline(start)
}

// howManyRunesUntil iterates the scanners rune slice executing the function for
// on each rune. The number of runes counted before the function results in true
// is returned to the user. If the function never returns true then the length
// of the rune slice, from the start index, is returned.
func (ss *symbolStream) howManyRunesUntil(start int, f func(int, rune) bool) (i int) {

	var ru rune

	for i, ru = range ss.runes[start:] {
		if f(i, ru) {
			break
		}
	}

	return i
}

func (ss *symbolStream) countNewlineRunes(start int) int {

	const (
		LF        = token.NEWLINE_LF
		CRLF      = token.NEWLINE_CRLF
		NOT_FOUND = 0
	)

	size := ss.len()

	if size > 0 && ss.matchesNonTerminal(start, LF) {
		return len(LF)
	}

	if size > 1 && ss.matchesNonTerminal(start, CRLF) {
		return len(CRLF)
	}

	return NOT_FOUND
}

func (ss *symbolStream) runesUntilNewline(start int) int {
	return ss.howManyRunesUntil(start, func(i int, ru rune) bool {
		return ss.matchesNewline(i)
	})
}

func (ss *symbolStream) countWordRunes(start int) int {
	return ss.howManyRunesUntil(start, func(i int, ru rune) bool {

		if i == 0 && ru == '_' {
			return true
		}

		return ru != '_' && !unicode.IsLetter(ru)
	})
}

func (ss *symbolStream) countDigitRunes(start int) int {
	return ss.howManyRunesUntil(start, func(_ int, ru rune) bool {
		return !unicode.IsDigit(ru)
	})
}

// tokenize slices off the next token from the scanners rune array and updates
// the line and column numbers accordingly.
func (ss *symbolStream) tokenize(n int, lex token.Lexeme, newline bool) (tk token.Token) {

	if ss.len() < n {
		panic("SANITY CHECK! Bad argument, n is bigger than the remaining source code")
	}

	tk = token.Token{
		Lexeme: lex,
		Value:  string(ss.runes[:n]),
		Line:   ss.line,
		Col:    ss.col,
	}

	ss.runes = ss.runes[n:]

	if newline {
		ss.line++
		ss.col = 0
	} else {
		ss.col += n
	}

	return
}

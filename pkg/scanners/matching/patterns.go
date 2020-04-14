package matching

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
)

// matcher implementations will return the number of terminals in a token, but
// only if the token appears next in the symbolStream else 0 is returned.
type matcher func(*symbolStream) int

// pattern represents a pattern for matching a specific representation of a
// lexeme. The pattern is provided as a matcher function.
type pattern struct {
	lexeme  lexeme.Lexeme
	matcher matcher
}

// patterns returns an array of all possible non-terminal symbols and their
// mapping to a lexeme. Longest and highest priority static symbols should be at
// the beginning of the array to ensure the correct token is scanned.
func patterns() []pattern {
	return []pattern{
		pattern{lexeme.LEXEME_NEWLINE, func(ss *symbolStream) int {
			return ss.countNewlineSymbols(0)
		}},
		pattern{lexeme.LEXEME_WHITESPACE, func(ss *symbolStream) int {
			// Returns the number of consecutive whitespace terminals.
			// Newlines are not counted as whitespace.
			return ss.countSymbolsWhile(0, func(i int, ru rune) bool {
				return !ss.isNewline(i) && unicode.IsSpace(ru)
			})
		}},
		pattern{lexeme.LEXEME_COMMENT, func(ss *symbolStream) int {
			if ss.isMatch(0, "//") {
				return ss.indexOfNextNewline(0)
			}
			return 0
		}},
		pattern{lexeme.LEXEME_BOOL, func(ss *symbolStream) int {
			return keywordMatcher(ss, "FALSE")
		}},
		pattern{lexeme.LEXEME_BOOL, func(ss *symbolStream) int {
			return keywordMatcher(ss, "TRUE")
		}},
		pattern{lexeme.LEXEME_END, func(ss *symbolStream) int {
			return keywordMatcher(ss, "END")
		}},
		pattern{lexeme.LEXEME_DO, func(ss *symbolStream) int {
			return keywordMatcher(ss, "DO")
		}},
		pattern{lexeme.LEXEME_FUNC, func(ss *symbolStream) int {
			return keywordMatcher(ss, "F")
		}},
		pattern{lexeme.LEXEME_ID, func(ss *symbolStream) int {
			return ss.countSymbolsWhile(0, func(i int, ru rune) bool {

				if unicode.IsLetter(ru) {
					return true
				}

				if i == 0 || ru != '_' {
					return false
				}

				return true
			})
		}},
		pattern{lexeme.LEXEME_ASSIGN, func(ss *symbolStream) int {
			return stringMatcher(ss, ":=")
		}},
		pattern{lexeme.LEXEME_RETURNS, func(ss *symbolStream) int {
			return stringMatcher(ss, "->")
		}},
		pattern{lexeme.LEXEME_LT_OR_EQU, func(ss *symbolStream) int {
			return stringMatcher(ss, "<=")
		}},
		pattern{lexeme.LEXEME_MT_OR_EQU, func(ss *symbolStream) int {
			return stringMatcher(ss, "=>")
		}},
		pattern{lexeme.LEXEME_PAREN_OPEN, func(ss *symbolStream) int {
			return stringMatcher(ss, "(")
		}},
		pattern{lexeme.LEXEME_PAREN_CLOSE, func(ss *symbolStream) int {
			return stringMatcher(ss, ")")
		}},
		pattern{lexeme.LEXEME_LIST_OPEN, func(ss *symbolStream) int {
			return stringMatcher(ss, "[")
		}},
		pattern{lexeme.LEXEME_LIST_CLOSE, func(ss *symbolStream) int {
			return stringMatcher(ss, "]")
		}},
		pattern{lexeme.LEXEME_DELIM, func(ss *symbolStream) int {
			return stringMatcher(ss, ",")
		}},
		pattern{lexeme.LEXEME_VOID, func(ss *symbolStream) int {
			return stringMatcher(ss, "_")
		}},
		pattern{lexeme.LEXEME_TERMINATOR, func(ss *symbolStream) int {
			return stringMatcher(ss, ";")
		}},
		pattern{lexeme.LEXEME_SPELL, func(ss *symbolStream) int {
			return stringMatcher(ss, "@")
		}},
		pattern{lexeme.LEXEME_ADD, func(ss *symbolStream) int {
			return stringMatcher(ss, "+")
		}},
		pattern{lexeme.LEXEME_SUBTRACT, func(ss *symbolStream) int {
			return stringMatcher(ss, "-")
		}},
		pattern{lexeme.LEXEME_MULTIPLY, func(ss *symbolStream) int {
			return stringMatcher(ss, "*")
		}},
		pattern{lexeme.LEXEME_DIVIDE, func(ss *symbolStream) int {
			return stringMatcher(ss, "/")
		}},
		pattern{lexeme.LEXEME_REMAINDER, func(ss *symbolStream) int {
			return stringMatcher(ss, "%")
		}},
		pattern{lexeme.LEXEME_AND, func(ss *symbolStream) int {
			return stringMatcher(ss, "&")
		}},
		pattern{lexeme.LEXEME_OR, func(ss *symbolStream) int {
			return stringMatcher(ss, "|")
		}},
		pattern{lexeme.LEXEME_EQU, func(ss *symbolStream) int {
			return stringMatcher(ss, "=")
		}},
		pattern{lexeme.LEXEME_NEQ, func(ss *symbolStream) int {
			return stringMatcher(ss, "#")
		}},
		pattern{lexeme.LEXEME_LT, func(ss *symbolStream) int {
			return stringMatcher(ss, "<")
		}},
		pattern{lexeme.LEXEME_MT, func(ss *symbolStream) int {
			return stringMatcher(ss, ">")
		}},
		pattern{lexeme.LEXEME_STRING, func(ss *symbolStream) int {

			const (
				PREFIX     = "`"
				SUFFIX     = "`"
				PREFIX_LEN = 1
				SUFFIX_LEN = 1
			)

			if !ss.isMatch(0, PREFIX) {
				return 0
			}

			n := ss.countSymbolsWhile(0, func(i int, ru rune) bool {

				if ss.isMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(ss, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		pattern{lexeme.LEXEME_TEMPLATE, func(ss *symbolStream) int {
			// As the name suggests, templates can be populated with the value of
			// identifiers, but the scanner is not concerned with parsing these. It does
			// need to watch out for escaped terminals that also represent the string
			// closer (suffix).

			const (
				PREFIX     = `"`
				SUFFIX     = `"`
				ESCAPE     = `/`
				PREFIX_LEN = 1
				SUFFIX_LEN = 1
			)

			if !ss.isMatch(0, PREFIX) {
				return 0
			}

			var prevEscaped bool

			n := ss.countSymbolsWhile(0, func(i int, ru rune) bool {

				escaped := prevEscaped
				prevEscaped = false

				if ss.isMatch(i, ESCAPE) {
					prevEscaped = !escaped
					return true
				}

				if !escaped && ss.isMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(ss, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		pattern{lexeme.LEXEME_FLOAT, func(ss *symbolStream) int {

			const (
				DELIM     = "."
				DELIM_LEN = len(DELIM)
			)

			n := integerMatcher(ss, 0)

			if n == 0 || n == ss.len() || !ss.isMatch(n, DELIM) {
				return 0
			}

			fractionalLen := integerMatcher(ss, n+DELIM_LEN)

			if fractionalLen == 0 {
				// One or many fractional digits must follow a delimiter. Zero following
				// digits is invalid syntax, so we must panic.
				panic(newErr(ss, n+DELIM_LEN,
					"Invalid syntax, expected digit after decimal point",
				))
			}

			return n + DELIM_LEN + fractionalLen
		}},
		pattern{lexeme.LEXEME_INT, func(ss *symbolStream) int {
			return integerMatcher(ss, 0)
		}},
	}
}

// keywordMatcher returns the number of terminal symbols in kw, but only if the
// next sequence of terminals matches the contents of kw and the terminal after
// is not a valid keyword terminal.
func keywordMatcher(ss *symbolStream, kw string) int {

	var WORD_LEN = len(kw)

	if stringMatcher(ss, kw) > 0 {
		if ss.len() == WORD_LEN || !unicode.IsLetter(ss.peekTerminal(WORD_LEN)) {
			return WORD_LEN
		}
	}

	return 0
}

// stringMatcher returns the number of terminal symbols in s, but only if the
// next sequence of terminals matches the contents of s.
func stringMatcher(ss *symbolStream, s string) int {

	if ss.len() >= len(s) && ss.isMatch(0, s) {
		return len(s)
	}

	return 0
}

// integerMatcher returns the number of terminal symbols of the next integer
// in the TerminalStream, but only if the next token is an integer else 0 is
// returned.
func integerMatcher(ss *symbolStream, start int) int {
	return ss.countSymbolsWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(ss *symbolStream, i int) {
	if ss.isNewline(i) {
		panic(newErr(ss, 0,
			"Newline encountered before a string or template was terminated",
		))
	}

	if i+1 == ss.len() {
		panic(newErr(ss, 0,
			"EOF encountered before a string or template was terminated",
		))
	}
}

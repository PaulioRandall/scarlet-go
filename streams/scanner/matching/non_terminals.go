package matching

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/terminal"
)

// matcher implementations will return the number of terminals in a token, but
// only if the token appears next in the TerminalStream else 0 is returned.
type matcher func(*terminal.TerminalStream) int

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
		pattern{lexeme.LEXEME_NEWLINE, func(ts *terminal.TerminalStream) int {
			return ts.CountNewlineSymbols(0)
		}},
		pattern{lexeme.LEXEME_WHITESPACE, func(ts *terminal.TerminalStream) int {
			// Returns the number of consecutive whitespace terminals.
			// Newlines are not counted as whitespace.
			return ts.CountSymbolsWhile(0, func(i int, ru rune) bool {
				return !ts.IsNewline(i) && unicode.IsSpace(ru)
			})
		}},
		pattern{lexeme.LEXEME_COMMENT, func(ts *terminal.TerminalStream) int {
			if ts.IsMatch(0, "//") {
				return ts.IndexOfNextNewline(0)
			}
			return 0
		}},
		pattern{lexeme.LEXEME_BOOL, func(ts *terminal.TerminalStream) int {
			return keywordMatcher(ts, "FALSE")
		}},
		pattern{lexeme.LEXEME_BOOL, func(ts *terminal.TerminalStream) int {
			return keywordMatcher(ts, "TRUE")
		}},
		pattern{lexeme.LEXEME_END, func(ts *terminal.TerminalStream) int {
			return keywordMatcher(ts, "END")
		}},
		pattern{lexeme.LEXEME_DO, func(ts *terminal.TerminalStream) int {
			return keywordMatcher(ts, "DO")
		}},
		pattern{lexeme.LEXEME_FUNC, func(ts *terminal.TerminalStream) int {
			return keywordMatcher(ts, "F")
		}},
		pattern{lexeme.LEXEME_ID, func(ts *terminal.TerminalStream) int {
			return ts.CountSymbolsWhile(0, func(i int, ru rune) bool {

				if unicode.IsLetter(ru) {
					return true
				}

				if i == 0 || ru != '_' {
					return false
				}

				return true
			})
		}},
		pattern{lexeme.LEXEME_ASSIGN, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, ":=")
		}},
		pattern{lexeme.LEXEME_RETURNS, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "->")
		}},
		pattern{lexeme.LEXEME_LT_OR_EQU, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "<=")
		}},
		pattern{lexeme.LEXEME_MT_OR_EQU, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "=>")
		}},
		pattern{lexeme.LEXEME_PAREN_OPEN, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "(")
		}},
		pattern{lexeme.LEXEME_PAREN_CLOSE, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, ")")
		}},
		pattern{lexeme.LEXEME_LIST_OPEN, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "[")
		}},
		pattern{lexeme.LEXEME_LIST_CLOSE, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "]")
		}},
		pattern{lexeme.LEXEME_DELIM, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, ",")
		}},
		pattern{lexeme.LEXEME_VOID, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "_")
		}},
		pattern{lexeme.LEXEME_TERMINATOR, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, ";")
		}},
		pattern{lexeme.LEXEME_SPELL, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "@")
		}},
		pattern{lexeme.LEXEME_ADD, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "+")
		}},
		pattern{lexeme.LEXEME_SUBTRACT, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "-")
		}},
		pattern{lexeme.LEXEME_MULTIPLY, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "*")
		}},
		pattern{lexeme.LEXEME_DIVIDE, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "/")
		}},
		pattern{lexeme.LEXEME_REMAINDER, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "%")
		}},
		pattern{lexeme.LEXEME_AND, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "&")
		}},
		pattern{lexeme.LEXEME_OR, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "|")
		}},
		pattern{lexeme.LEXEME_EQU, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "=")
		}},
		pattern{lexeme.LEXEME_NEQ, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "#")
		}},
		pattern{lexeme.LEXEME_LT, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, "<")
		}},
		pattern{lexeme.LEXEME_MT, func(ts *terminal.TerminalStream) int {
			return stringMatcher(ts, ">")
		}},
		pattern{lexeme.LEXEME_STRING, func(ts *terminal.TerminalStream) int {

			const (
				PREFIX     = "`"
				SUFFIX     = "`"
				PREFIX_LEN = 1
				SUFFIX_LEN = 1
			)

			if !ts.IsMatch(0, PREFIX) {
				return 0
			}

			n := ts.CountSymbolsWhile(0, func(i int, ru rune) bool {

				if ts.IsMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(ts, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		pattern{lexeme.LEXEME_TEMPLATE, func(ts *terminal.TerminalStream) int {
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

			if !ts.IsMatch(0, PREFIX) {
				return 0
			}

			var prevEscaped bool

			n := ts.CountSymbolsWhile(0, func(i int, ru rune) bool {

				escaped := prevEscaped
				prevEscaped = false

				if ts.IsMatch(i, ESCAPE) {
					prevEscaped = !escaped
					return true
				}

				if !escaped && ts.IsMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(ts, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		pattern{lexeme.LEXEME_FLOAT, func(ts *terminal.TerminalStream) int {

			const (
				DELIM     = "."
				DELIM_LEN = len(DELIM)
			)

			n := integerMatcher(ts, 0)

			if n == 0 || n == ts.Len() || !ts.IsMatch(n, DELIM) {
				return 0
			}

			fractionalLen := integerMatcher(ts, n+DELIM_LEN)

			if fractionalLen == 0 {
				// One or many fractional digits must follow a delimiter. Zero following
				// digits is invalid syntax, so we must panic.
				panic(newErr(ts, n+DELIM_LEN,
					"Invalid syntax, expected digit after decimal point",
				))
			}

			return n + DELIM_LEN + fractionalLen
		}},
		pattern{lexeme.LEXEME_INT, func(ts *terminal.TerminalStream) int {
			return integerMatcher(ts, 0)
		}},
	}
}

// keywordMatcher returns the number of terminal symbols in kw, but only if the
// next sequence of terminals matches the contents of kw and the terminal after
// is not a valid keyword terminal.
func keywordMatcher(ts *terminal.TerminalStream, kw string) int {

	var WORD_LEN = len(kw)

	if stringMatcher(ts, kw) > 0 {
		if ts.Len() == WORD_LEN || !unicode.IsLetter(ts.PeekTerminal(WORD_LEN)) {
			return WORD_LEN
		}
	}

	return 0
}

// stringMatcher returns the number of terminal symbols in s, but only if the
// next sequence of terminals matches the contents of s.
func stringMatcher(ts *terminal.TerminalStream, s string) int {

	if ts.Len() >= len(s) && ts.IsMatch(0, s) {
		return len(s)
	}

	return 0
}

// integerMatcher returns the number of terminal symbols of the next integer
// in the TerminalStream, but only if the next token is an integer else 0 is
// returned.
func integerMatcher(ts *terminal.TerminalStream, start int) int {
	return ts.CountSymbolsWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(ts *terminal.TerminalStream, i int) {
	if ts.IsNewline(i) {
		panic(newErr(ts, 0,
			"Newline encountered before a string or template was terminated",
		))
	}

	if i+1 == ts.Len() {
		panic(newErr(ts, 0,
			"EOF encountered before a string or template was terminated",
		))
	}
}

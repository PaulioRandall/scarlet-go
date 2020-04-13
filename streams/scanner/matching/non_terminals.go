package matching

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/symbol"
)

// matcher implementations will return the number of terminals in a token, but
// only if the token appears next in the TerminalStream else 0 is returned.
type matcher func(*symbol.TerminalStream) int

// nonTerminal represents a mapping between a lexeme and a Matcher function.
type nonTerminal struct {
	lexeme  lexeme.Lexeme
	matcher matcher
}

// nonTerminals returns an array of all possible non-terminal symbols and their
// mapping to a lexeme. Longest and highest priority static symbols should be at
// the beginning of the array to ensure the correct token is scanned.
func nonTerminals() []nonTerminal {
	return []nonTerminal{
		nonTerminal{lexeme.LEXEME_NEWLINE, func(ts *symbol.TerminalStream) int {
			return ts.CountNewlineSymbols(0)
		}},
		nonTerminal{lexeme.LEXEME_WHITESPACE, func(ts *symbol.TerminalStream) int {
			// Returns the number of consecutive whitespace terminals.
			// Newlines are not counted as whitespace.
			return ts.CountSymbolsWhile(0, func(i int, ru rune) bool {
				return !ts.IsNewline(i) && unicode.IsSpace(ru)
			})
		}},
		nonTerminal{lexeme.LEXEME_COMMENT, func(ts *symbol.TerminalStream) int {
			if ts.IsMatch(0, "//") {
				return ts.IndexOfNextNewline(0)
			}
			return 0
		}},
		nonTerminal{lexeme.LEXEME_BOOL, func(ts *symbol.TerminalStream) int {
			return keywordMatcher(ts, "FALSE")
		}},
		nonTerminal{lexeme.LEXEME_BOOL, func(ts *symbol.TerminalStream) int {
			return keywordMatcher(ts, "TRUE")
		}},
		nonTerminal{lexeme.LEXEME_END, func(ts *symbol.TerminalStream) int {
			return keywordMatcher(ts, "END")
		}},
		nonTerminal{lexeme.LEXEME_DO, func(ts *symbol.TerminalStream) int {
			return keywordMatcher(ts, "DO")
		}},
		nonTerminal{lexeme.LEXEME_FUNC, func(ts *symbol.TerminalStream) int {
			return keywordMatcher(ts, "F")
		}},
		nonTerminal{lexeme.LEXEME_ID, func(ts *symbol.TerminalStream) int {
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
		nonTerminal{lexeme.LEXEME_ASSIGN, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, ":=")
		}},
		nonTerminal{lexeme.LEXEME_RETURNS, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "->")
		}},
		nonTerminal{lexeme.LEXEME_LT_OR_EQU, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "<=")
		}},
		nonTerminal{lexeme.LEXEME_MT_OR_EQU, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "=>")
		}},
		nonTerminal{lexeme.LEXEME_PAREN_OPEN, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "(")
		}},
		nonTerminal{lexeme.LEXEME_PAREN_CLOSE, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, ")")
		}},
		nonTerminal{lexeme.LEXEME_LIST_OPEN, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "[")
		}},
		nonTerminal{lexeme.LEXEME_LIST_CLOSE, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "]")
		}},
		nonTerminal{lexeme.LEXEME_DELIM, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, ",")
		}},
		nonTerminal{lexeme.LEXEME_VOID, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "_")
		}},
		nonTerminal{lexeme.LEXEME_TERMINATOR, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, ";")
		}},
		nonTerminal{lexeme.LEXEME_SPELL, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "@")
		}},
		nonTerminal{lexeme.LEXEME_ADD, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "+")
		}},
		nonTerminal{lexeme.LEXEME_SUBTRACT, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "-")
		}},
		nonTerminal{lexeme.LEXEME_MULTIPLY, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "*")
		}},
		nonTerminal{lexeme.LEXEME_DIVIDE, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "/")
		}},
		nonTerminal{lexeme.LEXEME_REMAINDER, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "%")
		}},
		nonTerminal{lexeme.LEXEME_AND, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "&")
		}},
		nonTerminal{lexeme.LEXEME_OR, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "|")
		}},
		nonTerminal{lexeme.LEXEME_EQU, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "=")
		}},
		nonTerminal{lexeme.LEXEME_NEQ, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "#")
		}},
		nonTerminal{lexeme.LEXEME_LT, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, "<")
		}},
		nonTerminal{lexeme.LEXEME_MT, func(ts *symbol.TerminalStream) int {
			return stringMatcher(ts, ">")
		}},
		nonTerminal{lexeme.LEXEME_STRING, func(ts *symbol.TerminalStream) int {

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
		nonTerminal{lexeme.LEXEME_TEMPLATE, func(ts *symbol.TerminalStream) int {
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
		nonTerminal{lexeme.LEXEME_FLOAT, func(ts *symbol.TerminalStream) int {

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
		nonTerminal{lexeme.LEXEME_INT, func(ts *symbol.TerminalStream) int {
			return integerMatcher(ts, 0)
		}},
	}
}

// keywordMatcher returns the number of terminal symbols in kw, but only if the
// next sequence of terminals matches the contents of kw and the symbol after
// is not a valid keyword terminal.
func keywordMatcher(ts *symbol.TerminalStream, kw string) int {

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
func stringMatcher(ts *symbol.TerminalStream, s string) int {

	if ts.Len() >= len(s) && ts.IsMatch(0, s) {
		return len(s)
	}

	return 0
}

// integerMatcher returns the number of terminal symbols of the next integer
// in the TerminalStream, but only if the next token is an integer else 0 is
// returned.
func integerMatcher(ts *symbol.TerminalStream, start int) int {
	return ts.CountSymbolsWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(ts *symbol.TerminalStream, i int) {
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

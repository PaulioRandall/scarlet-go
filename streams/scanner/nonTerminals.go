package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/symbol"
)

// Matcher implementations will return the number of terminals token, but only
// if the token appears next in the SymbolStream else 0 is returned.
type Matcher func(symbol.SymbolStream) int

// nonTerminal represents a mapping between a lexeme and a Matcher function.
type nonTerminal struct {
	Lexeme  lexeme.Lexeme
	Matcher Matcher
}

// nonTerminals returns an array of all possible non-terminal symbols and their
// mapping to a lexeme. Longest and highest priority symbols should be at the
// beginning of the array to ensure the correct token is scanned.
func nonTerminals() []nonTerminal {
	return []nonTerminal{
		nonTerminal{lexeme.LEXEME_NEWLINE, func(ss symbol.SymbolStream) int {
			return ss.CountNewlineSymbols(0)
		}},
		nonTerminal{lexeme.LEXEME_WHITESPACE, func(ss symbol.SymbolStream) int {
			// Returns the number of consecutive whitespace terminals.
			// Newlines are not counted as whitespace.
			return ss.CountSymbolsWhile(0, func(i int, ru rune) bool {
				return !ss.IsNewline(i) && unicode.IsSpace(ru)
			})
		}},
		nonTerminal{lexeme.LEXEME_COMMENT, func(ss symbol.SymbolStream) int {
			if ss.IsMatch(0, "//") {
				return ss.IndexOfNextNewline(2)
			}
			return 0
		}},
		nonTerminal{lexeme.LEXEME_BOOL, func(ss symbol.SymbolStream) int {
			return keywordMatcher(ss, "FALSE")
		}},
		nonTerminal{lexeme.LEXEME_BOOL, func(ss symbol.SymbolStream) int {
			return keywordMatcher(ss, "TRUE")
		}},
		nonTerminal{lexeme.LEXEME_END, func(ss symbol.SymbolStream) int {
			return keywordMatcher(ss, "END")
		}},
		nonTerminal{lexeme.LEXEME_DO, func(ss symbol.SymbolStream) int {
			return keywordMatcher(ss, "DO")
		}},
		nonTerminal{lexeme.LEXEME_FUNC, func(ss symbol.SymbolStream) int {
			return keywordMatcher(ss, "F")
		}},
		nonTerminal{lexeme.LEXEME_ID, func(ss symbol.SymbolStream) int {
			return ss.CountSymbolsWhile(0, func(i int, ru rune) bool {

				if unicode.IsLetter(ru) {
					return true
				}

				if i == 0 || ru != '_' {
					return false
				}

				return true
			})
		}},
		nonTerminal{lexeme.LEXEME_ASSIGN, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, ":=")
		}},
		nonTerminal{lexeme.LEXEME_RETURNS, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "->")
		}},
		nonTerminal{lexeme.LEXEME_LT_OR_EQU, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "<=")
		}},
		nonTerminal{lexeme.LEXEME_MT_OR_EQU, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "=>")
		}},
		nonTerminal{lexeme.LEXEME_OPEN_PAREN, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "(")
		}},
		nonTerminal{lexeme.LEXEME_CLOSE_PAREN, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, ")")
		}},
		nonTerminal{lexeme.LEXEME_OPEN_GUARD, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "[")
		}},
		nonTerminal{lexeme.LEXEME_CLOSE_GUARD, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "]")
		}},
		nonTerminal{lexeme.LEXEME_OPEN_LIST, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "{")
		}},
		nonTerminal{lexeme.LEXEME_CLOSE_LIST, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "}")
		}},
		nonTerminal{lexeme.LEXEME_DELIM, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, ",")
		}},
		nonTerminal{lexeme.LEXEME_VOID, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "_")
		}},
		nonTerminal{lexeme.LEXEME_TERMINATOR, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, ";")
		}},
		nonTerminal{lexeme.LEXEME_SPELL, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "@")
		}},
		nonTerminal{lexeme.LEXEME_ADD, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "+")
		}},
		nonTerminal{lexeme.LEXEME_SUBTRACT, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "-")
		}},
		nonTerminal{lexeme.LEXEME_MULTIPLY, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "*")
		}},
		nonTerminal{lexeme.LEXEME_DIVIDE, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "/")
		}},
		nonTerminal{lexeme.LEXEME_REMAINDER, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "%")
		}},
		nonTerminal{lexeme.LEXEME_AND, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "&")
		}},
		nonTerminal{lexeme.LEXEME_OR, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "|")
		}},
		nonTerminal{lexeme.LEXEME_EQU, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "=")
		}},
		nonTerminal{lexeme.LEXEME_NEQ, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "#")
		}},
		nonTerminal{lexeme.LEXEME_LT, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, "<")
		}},
		nonTerminal{lexeme.LEXEME_MT, func(ss symbol.SymbolStream) int {
			return stringMatcher(ss, ">")
		}},
		// TODO: string templates
		nonTerminal{lexeme.LEXEME_STRING, func(ss symbol.SymbolStream) int {

			const (
				PREFIX     = "`"
				SUFFIX     = "`"
				PREFIX_LEN = 1
				SUFFIX_LEN = 1
			)

			if !ss.IsMatch(0, PREFIX) {
				return 0
			}

			n := ss.CountSymbolsWhile(0, func(i int, ru rune) bool {

				if ss.IsMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(ss, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		nonTerminal{lexeme.LEXEME_TEMPLATE, func(ss symbol.SymbolStream) int {
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

			if !ss.IsMatch(0, PREFIX) {
				return 0
			}

			var prevEscaped bool

			n := ss.CountSymbolsWhile(0, func(i int, ru rune) bool {

				escaped := prevEscaped
				prevEscaped = false

				if ss.IsMatch(i, ESCAPE) {
					prevEscaped = !escaped
					return true
				}

				if !escaped && ss.IsMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(ss, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		nonTerminal{lexeme.LEXEME_FLOAT, func(ss symbol.SymbolStream) int {

			const (
				DELIM     = "."
				DELIM_LEN = len(DELIM)
			)

			n := integerMatcher(ss, 0)

			if n == 0 || n == ss.Len() || !ss.IsMatch(n, DELIM) {
				return 0
			}

			fractionalLen := integerMatcher(ss, n+DELIM_LEN)

			if fractionalLen == 0 {
				// One or many fractional digits must follow a delimiter. Zero following
				// digits is invalid syntax, so we must panic.
				panic(terror(ss, n+DELIM_LEN,
					"Invalid syntax, expected digit after decimal point",
				))
			}

			return n + DELIM_LEN + fractionalLen
		}},
		nonTerminal{lexeme.LEXEME_INT, func(ss symbol.SymbolStream) int {
			return integerMatcher(ss, 0)
		}},
	}
}

// keywordMatcher returns the number of terminal symbols in kw, but only if the
// next sequence of terminals matches the contents of kw and the symbol after
// is not a valid keyword terminal.
func keywordMatcher(ss symbol.SymbolStream, kw string) int {

	var WORD_LEN = len(kw)

	if stringMatcher(ss, kw) > 0 {
		if ss.Len() == WORD_LEN || !unicode.IsLetter(ss.PeekSymbol(WORD_LEN)) {
			return WORD_LEN
		}
	}

	return 0
}

// stringMatcher returns the number of terminal symbols in s, but only if the
// next sequence of terminals matches the contents of s.
func stringMatcher(ss symbol.SymbolStream, s string) int {

	if ss.Len() >= len(s) && ss.IsMatch(0, s) {
		return len(s)
	}

	return 0
}

// integerMatcher returns the number of terminal symbols of the next integer
// in the SymbolStream, but only if the next token is an integer else 0 is
// returned.
func integerMatcher(ss symbol.SymbolStream, start int) int {
	return ss.CountSymbolsWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(ss symbol.SymbolStream, i int) {
	if ss.IsNewline(i) {
		panic(terror(ss, 0,
			"Newline encountered before a string or template was terminated",
		))
	}

	if i+1 == ss.Len() {
		panic(terror(ss, 0,
			"EOF encountered before a string or template was terminated",
		))
	}
}

package matching

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// matcher implementations will return the number of terminals in a token, but
// only if the token appears next in the symbolStream else 0 is returned.
type matcher func(*symbolStream) int

type pattern struct {
	tokenType token.TokenType
	matcher   matcher
}

func patterns() []pattern {
	return []pattern{
		pattern{token.NEWLINE, func(ss *symbolStream) int {
			return ss.countNewlineSymbols(0)
		}},
		pattern{token.WHITESPACE, func(ss *symbolStream) int {
			// Returns the number of consecutive whitespace terminals.
			// Newlines are not counted as whitespace.
			return ss.countSymbolsWhile(0, func(i int, ru rune) bool {
				return !ss.isNewline(i) && unicode.IsSpace(ru)
			})
		}},
		pattern{token.COMMENT, func(ss *symbolStream) int {
			if ss.isMatch(0, "//") {
				return ss.indexOfNextNewline(0)
			}
			return 0
		}},
		pattern{token.MATCH, func(ss *symbolStream) int {
			return matchWord(ss, "MATCH")
		}},
		pattern{token.BOOL, func(ss *symbolStream) int {
			return matchWord(ss, "FALSE")
		}},
		pattern{token.BOOL, func(ss *symbolStream) int {
			return matchWord(ss, "TRUE")
		}},
		pattern{token.LIST, func(ss *symbolStream) int {
			return matchWord(ss, "LIST")
		}},
		pattern{token.FIX, func(ss *symbolStream) int {
			return matchWord(ss, "FIX")
		}},
		pattern{token.EOF, func(ss *symbolStream) int {
			return matchWord(ss, "EOF")
		}},
		pattern{token.FUNC, func(ss *symbolStream) int {
			return matchWord(ss, "F")
		}},
		pattern{token.ID, func(ss *symbolStream) int {
			return ss.countSymbolsWhile(0, func(i int, ru rune) bool {

				if unicode.IsLetter(ru) {
					return true
				}

				return i != 0 && ru == '_'
			})
		}},
		pattern{token.ASSIGN, func(ss *symbolStream) int {
			return matchStr(ss, ":=")
		}},
		pattern{token.RETURNS, func(ss *symbolStream) int {
			return matchStr(ss, "->")
		}},
		pattern{token.LESS_THAN_OR_EQUAL, func(ss *symbolStream) int {
			return matchStr(ss, "<=")
		}},
		pattern{token.MORE_THAN_OR_EQUAL, func(ss *symbolStream) int {
			return matchStr(ss, ">=")
		}},
		pattern{token.BLOCK_OPEN, func(ss *symbolStream) int {
			return matchStr(ss, "{")
		}},
		pattern{token.BLOCK_CLOSE, func(ss *symbolStream) int {
			return matchStr(ss, "}")
		}},
		pattern{token.PAREN_OPEN, func(ss *symbolStream) int {
			return matchStr(ss, "(")
		}},
		pattern{token.PAREN_CLOSE, func(ss *symbolStream) int {
			return matchStr(ss, ")")
		}},
		pattern{token.GUARD_OPEN, func(ss *symbolStream) int {
			return matchStr(ss, "[")
		}},
		pattern{token.GUARD_CLOSE, func(ss *symbolStream) int {
			return matchStr(ss, "]")
		}},
		pattern{token.DELIM, func(ss *symbolStream) int {
			return matchStr(ss, ",")
		}},
		pattern{token.VOID, func(ss *symbolStream) int {
			return matchStr(ss, "_")
		}},
		pattern{token.TERMINATOR, func(ss *symbolStream) int {
			return matchStr(ss, ";")
		}},
		pattern{token.SPELL, func(ss *symbolStream) int {
			return matchStr(ss, "@")
		}},
		pattern{token.ADD, func(ss *symbolStream) int {
			return matchStr(ss, "+")
		}},
		pattern{token.SUBTRACT, func(ss *symbolStream) int {
			return matchStr(ss, "-")
		}},
		pattern{token.MULTIPLY, func(ss *symbolStream) int {
			return matchStr(ss, "*")
		}},
		pattern{token.DIVIDE, func(ss *symbolStream) int {
			return matchStr(ss, "/")
		}},
		pattern{token.REMAINDER, func(ss *symbolStream) int {
			return matchStr(ss, "%")
		}},
		pattern{token.AND, func(ss *symbolStream) int {
			return matchStr(ss, "&")
		}},
		pattern{token.OR, func(ss *symbolStream) int {
			return matchStr(ss, "|")
		}},
		pattern{token.EQUAL, func(ss *symbolStream) int {
			return matchStr(ss, "=")
		}},
		pattern{token.NOT_EQUAL, func(ss *symbolStream) int {
			return matchStr(ss, "#")
		}},
		pattern{token.LESS_THAN, func(ss *symbolStream) int {
			return matchStr(ss, "<")
		}},
		pattern{token.MORE_THAN, func(ss *symbolStream) int {
			return matchStr(ss, ">")
		}},
		pattern{token.STRING, func(ss *symbolStream) int {

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
		pattern{token.TEMPLATE, func(ss *symbolStream) int {
			// As the name suggests, templates can be populated with the value of
			// identifiers, but the scanner is not concerned with parsing these. It
			// does need to watch out for escaped terminals that also represent the
			// string closer (suffix).

			const (
				PREFIX     = `"`
				SUFFIX     = `"`
				ESCAPE     = `\`
				SUFFIX_LEN = 1
			)

			if !ss.isMatch(0, PREFIX) {
				return 0
			}

			escaped := true // Init true to escape prefix

			n := ss.countSymbolsWhile(0, func(i int, ru rune) bool {

				switch {
				case escaped:
					escaped = false

				case ss.isMatch(i, SUFFIX):
					return false

				case ss.isMatch(i, ESCAPE):
					escaped = true
					return true
				}

				checkForMissingTermination(ss, i)
				return true
			})

			return n + SUFFIX_LEN
		}},
		pattern{token.NUMBER, func(ss *symbolStream) int {

			const (
				DELIM     = "."
				DELIM_LEN = len(DELIM)
			)

			n := matchInt(ss, 0)

			if n == 0 || n == ss.len() || !ss.isMatch(n, DELIM) {
				return n
			}

			fractionalLen := matchInt(ss, n+DELIM_LEN)

			if fractionalLen == 0 {
				// One or many fractional digits must follow a delimiter.
				panic(newErr(ss, n+DELIM_LEN,
					"Invalid syntax, expected digit after decimal point",
				))
			}

			return n + DELIM_LEN + fractionalLen
		}},
	}
}

func matchWord(ss *symbolStream, kw string) int {

	var WORD_LEN = len(kw)

	if matchStr(ss, kw) > 0 {
		if ss.len() == WORD_LEN || !unicode.IsLetter(ss.peekTerminal(WORD_LEN)) {
			return WORD_LEN
		}
	}

	return 0
}

func matchStr(ss *symbolStream, s string) int {

	if ss.len() >= len(s) && ss.isMatch(0, s) {
		return len(s)
	}

	return 0
}

func matchInt(ss *symbolStream, start int) int {
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

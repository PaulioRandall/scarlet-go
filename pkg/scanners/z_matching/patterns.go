package z_matching

import (
	"unicode"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

// matcher implementations will return the number of terminals in a token, but
// only if the token appears next in the symbols else 0 is returned.
type matcher func(*symbols) int

type pattern struct {
	kind     Kind
	morpheme Morpheme
	matcher  matcher
}

func patterns() []pattern {
	return []pattern{
		pattern{K_NEWLINE, M_NEWLINE, func(s *symbols) int {
			return s.countNewlineSymbols(0)
		}},
		pattern{K_REDUNDANT, M_WHITESPACE, func(s *symbols) int {
			// Returns the number of consecutive whitespace terminals.
			// Newlines are not counted as whitespace.
			return s.countSymbolsWhile(0, func(i int, ru rune) bool {
				return !s.isNewline(i) && unicode.IsSpace(ru)
			})
		}},
		pattern{K_REDUNDANT, M_COMMENT, func(s *symbols) int {
			if s.isMatch(0, "//") {
				return s.indexOfNextNewline(0)
			}
			return 0
		}},
		pattern{K_KEYWORD, M_MATCH, func(s *symbols) int {
			return matchWord(s, "MATCH")
		}},
		pattern{K_KEYWORD, M_BOOL, func(s *symbols) int {
			return matchWord(s, "FALSE")
		}},
		pattern{K_KEYWORD, M_BOOL, func(s *symbols) int {
			return matchWord(s, "TRUE")
		}},
		pattern{K_KEYWORD, M_LIST, func(s *symbols) int {
			return matchWord(s, "LIST")
		}},
		pattern{K_KEYWORD, M_LOOP, func(s *symbols) int {
			return matchWord(s, "LOOP")
		}},
		pattern{K_KEYWORD, M_FIX, func(s *symbols) int {
			return matchWord(s, "FIX")
		}},
		pattern{K_KEYWORD, M_EOF, func(s *symbols) int {
			return matchWord(s, "EOF")
		}},
		pattern{K_KEYWORD, M_FUNC, func(s *symbols) int {
			return matchWord(s, "F")
		}},
		pattern{K_IDENTIFIER, M_IDENTIFIER, func(s *symbols) int {
			return s.countSymbolsWhile(0, func(i int, ru rune) bool {

				if unicode.IsLetter(ru) {
					return true
				}

				return i != 0 && ru == '_'
			})
		}},
		pattern{K_DELIMITER, M_ASSIGN, func(s *symbols) int {
			return matchStr(s, ":=")
		}},
		pattern{K_REFERENCE, M_LIST_END, func(s *symbols) int {
			return matchStr(s, ">>")
		}},
		pattern{K_REFERENCE, M_LIST_START, func(s *symbols) int {
			return matchStr(s, "<<")
		}},
		pattern{K_COMPARISON, M_LESS_THAN_OR_EQUAL, func(s *symbols) int {
			return matchStr(s, "<=")
		}},
		pattern{K_COMPARISON, M_MORE_THAN_OR_EQUAL, func(s *symbols) int {
			return matchStr(s, ">=")
		}},
		pattern{K_DELIMITER, M_BLOCK_OPEN, func(s *symbols) int {
			return matchStr(s, "{")
		}},
		pattern{K_DELIMITER, M_BLOCK_CLOSE, func(s *symbols) int {
			return matchStr(s, "}")
		}},
		pattern{K_DELIMITER, M_PAREN_OPEN, func(s *symbols) int {
			return matchStr(s, "(")
		}},
		pattern{K_DELIMITER, M_PAREN_CLOSE, func(s *symbols) int {
			return matchStr(s, ")")
		}},
		pattern{K_DELIMITER, M_GUARD_OPEN, func(s *symbols) int {
			return matchStr(s, "[")
		}},
		pattern{K_DELIMITER, M_GUARD_CLOSE, func(s *symbols) int {
			return matchStr(s, "]")
		}},
		pattern{K_KEYWORD, M_OUTPUT, func(s *symbols) int {
			return matchStr(s, "^")
		}},
		pattern{K_DELIMITER, M_DELIMITER, func(s *symbols) int {
			return matchStr(s, ",")
		}},
		pattern{K_IDENTIFIER, M_VOID, func(s *symbols) int {
			return matchStr(s, "_")
		}},
		pattern{K_DELIMITER, M_TERMINATOR, func(s *symbols) int {
			return matchStr(s, ";")
		}},
		pattern{K_KEYWORD, M_SPELL, func(s *symbols) int {
			return matchStr(s, "@")
		}},
		pattern{K_ARITHMETIC, M_ADD, func(s *symbols) int {
			return matchStr(s, "+")
		}},
		pattern{K_ARITHMETIC, M_SUBTRACT, func(s *symbols) int {
			return matchStr(s, "-")
		}},
		pattern{K_ARITHMETIC, M_MULTIPLY, func(s *symbols) int {
			return matchStr(s, "*")
		}},
		pattern{K_ARITHMETIC, M_DIVIDE, func(s *symbols) int {
			return matchStr(s, "/")
		}},
		pattern{K_ARITHMETIC, M_REMAINDER, func(s *symbols) int {
			return matchStr(s, "%")
		}},
		pattern{K_LOGIC, M_AND, func(s *symbols) int {
			return matchStr(s, "&")
		}},
		pattern{K_LOGIC, M_OR, func(s *symbols) int {
			return matchStr(s, "|")
		}},
		pattern{K_COMPARISON, M_EQUAL, func(s *symbols) int {
			return matchStr(s, "==")
		}},
		pattern{K_COMPARISON, M_NOT_EQUAL, func(s *symbols) int {
			return matchStr(s, "!=")
		}},
		pattern{K_COMPARISON, M_LESS_THAN, func(s *symbols) int {
			return matchStr(s, "<")
		}},
		pattern{K_COMPARISON, M_MORE_THAN, func(s *symbols) int {
			return matchStr(s, ">")
		}},
		pattern{K_LITERAL, M_STRING, func(s *symbols) int {

			const (
				PREFIX     = "`"
				SUFFIX     = "`"
				PREFIX_LEN = 1
				SUFFIX_LEN = 1
			)

			if !s.isMatch(0, PREFIX) {
				return 0
			}

			n := s.countSymbolsWhile(0, func(i int, ru rune) bool {

				if s.isMatch(i+PREFIX_LEN, SUFFIX) {
					return false
				}

				checkForMissingTermination(s, i)
				return true
			})

			return PREFIX_LEN + n + SUFFIX_LEN
		}},
		pattern{K_LITERAL, M_TEMPLATE, func(s *symbols) int {
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

			if !s.isMatch(0, PREFIX) {
				return 0
			}

			escaped := true // Init true to escape prefix

			n := s.countSymbolsWhile(0, func(i int, ru rune) bool {

				switch {
				case escaped:
					escaped = false

				case s.isMatch(i, SUFFIX):
					return false

				case s.isMatch(i, ESCAPE):
					escaped = true
					return true
				}

				checkForMissingTermination(s, i)
				return true
			})

			return n + SUFFIX_LEN
		}},
		pattern{K_LITERAL, M_NUMBER, func(s *symbols) int {

			const (
				DELIM     = "."
				DELIM_LEN = 1
			)

			n := matchInt(s, 0)

			if n == 0 || n == s.len() || !s.isMatch(n, DELIM) {
				return n
			}

			fractionalLen := matchInt(s, n+DELIM_LEN)

			if fractionalLen == 0 {
				// One or many fractional digits must follow a delimiter.
				panic(err(s, n, "Invalid syntax, expected digit after decimal point"))
			}

			return n + DELIM_LEN + fractionalLen
		}},
	}
}

func matchWord(s *symbols, kw string) int {

	var WORD_LEN = len(kw)

	if matchStr(s, kw) > 0 {
		if s.len() == WORD_LEN || !unicode.IsLetter(s.peekTerminal(WORD_LEN)) {
			return WORD_LEN
		}
	}

	return 0
}

func matchStr(s *symbols, str string) int {

	if s.len() >= len(str) && s.isMatch(0, str) {
		return len(str)
	}

	return 0
}

func matchInt(s *symbols, start int) int {
	return s.countSymbolsWhile(start, func(_ int, ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(s *symbols, i int) {
	if s.isNewline(i) {
		panic(err(s, i,
			"Newline encountered before a string or template was terminated",
		))
	}

	if i+1 == s.len() {
		panic(err(s, i,
			"EOF encountered before a string or template was terminated",
		))
	}
}

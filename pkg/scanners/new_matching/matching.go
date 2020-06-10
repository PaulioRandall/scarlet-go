package matching

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	mat "github.com/PaulioRandall/scarlet-go/pkg/scanners/new_matching/matcher"
)

func ScanAll(s string) ([]Token, error) {

	tks := []Token{}
	m := mat.New(s, patterns)

	for {

		v, e := m.Next()
		if e != nil {
			return nil, e
		}

		if v == nil {
			break
		}

		tk, _ := v.(Token)
		tks = append(tks, tk)
	}

	return tks, nil
}

type pattern struct {
	m Morpheme
	f mat.PatternMatcher
}

func (p pattern) Matcher() mat.PatternMatcher {
	return p.f
}

func (p pattern) OnMatch(value string, line, col int) interface{} {
	return NewToken(p.m, value, line, col)
}

var patterns = []mat.Pattern{
	pattern{NEWLINE, func(s *mat.Symbols) (int, error) {
		_, n := s.IsNewline(0)
		return n, nil
	}},
	pattern{WHITESPACE, func(s *mat.Symbols) (int, error) {
		// Returns the number of consecutive whitespace terminals.
		// Newlines are not counted as whitespace.
		return s.CountWhile(0, func(i int, ru rune) (bool, error) {

			if !unicode.IsSpace(ru) {
				return false, nil
			}

			if yes, _ := s.IsNewline(i); yes {
				return false, nil
			}

			return true, nil
		})
	}},
	pattern{COMMENT, func(s *mat.Symbols) (int, error) {

		PREFIX := "//"
		PREFIX_LEN := 2

		matched, e := s.Match(0, PREFIX)
		if e != nil || !matched {
			return 0, e
		}

		n, e := s.CountWhile(PREFIX_LEN, func(i int, ru rune) (bool, error) {
			ok, _ := s.IsNewline(i)
			return !ok, nil
		})

		if e != nil {
			return 0, e
		}

		return PREFIX_LEN + n, nil
	}},
	pattern{WHEN, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "when")
	}},
	pattern{WATCH, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "watch")
	}},
	pattern{BOOL, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "false")
	}},
	pattern{BOOL, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "true")
	}},
	pattern{LOOP, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "loop")
	}},
	pattern{DEF, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "def")
	}},
	pattern{FUNC, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "F")
	}},
	pattern{EXPR_FUNC, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "E")
	}},
	pattern{IDENTIFIER, func(s *mat.Symbols) (int, error) {
		return s.CountWhile(0, func(i int, ru rune) (bool, error) {

			if i == 0 {
				return unicode.IsLetter(ru), nil
			}

			if unicode.IsLetter(ru) || ru == '_' {
				return true, nil
			}

			return false, nil
		})
	}},
	pattern{ASSIGN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ":=")
	}},
	pattern{LESS_THAN_OR_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "<=")
	}},
	pattern{MORE_THAN_OR_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ">=")
	}},
	pattern{THEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "->")
	}},
	pattern{BLOCK_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "{")
	}},
	pattern{BLOCK_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "}")
	}},
	pattern{PAREN_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "(")
	}},
	pattern{PAREN_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ")")
	}},
	pattern{GUARD_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "[")
	}},
	pattern{GUARD_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "]")
	}},
	pattern{OUTPUT, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "^")
	}},
	pattern{DELIMITER, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ",")
	}},
	pattern{VOID, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "_")
	}},
	pattern{TERMINATOR, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ";")
	}},
	pattern{SPELL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "@")
	}},
	pattern{ADD, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "+")
	}},
	pattern{SUBTRACT, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "-")
	}},
	pattern{MULTIPLY, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "*")
	}},
	pattern{DIVIDE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "/")
	}},
	pattern{REMAINDER, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "%")
	}},
	pattern{AND, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "&")
	}},
	pattern{OR, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "|")
	}},
	pattern{EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "==")
	}},
	pattern{NOT_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "!=")
	}},
	pattern{LESS_THAN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "<")
	}},
	pattern{MORE_THAN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ">")
	}},
	pattern{STRING, func(s *mat.Symbols) (int, error) {

		const (
			PREFIX     = '"'
			SUFFIX     = '"'
			ESCAPE     = '\\'
			SUFFIX_LEN = 1
		)

		if s.At(0) != PREFIX {
			return 0, nil
		}

		escaped := true // Init true to escape prefix

		n, e := s.CountWhile(0, func(i int, ru rune) (bool, error) {

			resume := true

			if escaped {
				escaped = false
				goto FINALLY
			}

			if ru == SUFFIX {
				resume = false
				goto FINALLY
			}

			if ru == ESCAPE {
				escaped = true
			}

		FINALLY:
			e := checkForMissingTermination(s, i)
			return resume, e
		})

		if e != nil {
			return 0, e
		}

		return n + SUFFIX_LEN, nil
	}},
	pattern{NUMBER, func(s *mat.Symbols) (int, error) {

		const (
			DELIM     = '.'
			DELIM_LEN = 1
		)

		n, e := matchInt(s, 0)
		if e != nil {
			return 0, e
		}

		if n <= 0 || !s.Has(n+1) {
			return n, nil
		}

		if s.At(n) != DELIM {
			return n, nil
		}

		frac, e := matchInt(s, n+DELIM_LEN)
		if e != nil {
			return 0, e
		}

		if frac == 0 {
			// One or many fractional digits must follow a delimiter.
			line, col := s.Pos()
			return 0, err.New(
				"Expected digit after decimal point",
				err.Pos(line, col+n),
			)
		}

		return n + DELIM_LEN + frac, nil
	}},
}

func matchStr(s *mat.Symbols, str string) (int, error) {

	if !s.Has(len(str)) {
		return 0, nil
	}

	matched, e := s.Match(0, str)
	if e != nil || !matched {
		return 0, e
	}

	return len(str), nil
}

func matchWord(s *mat.Symbols, word string) (int, error) {

	size := len(word)
	if !s.Has(size) {
		return 0, nil
	}

	matched, e := s.Match(0, word)
	if e != nil || !matched {
		return 0, e
	}

	if !s.Has(size + 1) {
		return size, nil
	}

	ru := s.At(size)
	if unicode.IsLetter(ru) || ru == '_' {
		return 0, nil
	}

	return size, nil
}

func matchInt(s *mat.Symbols, start int) (int, error) {
	return s.CountWhile(start, func(_ int, ru rune) (bool, error) {
		return unicode.IsDigit(ru), nil
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(s *mat.Symbols, i int) error {

	line, col := s.Pos()

	if ok, _ := s.IsNewline(i); ok {
		return err.New(
			"Newline encountered before string was terminated",
			err.Pos(line, col+i),
		)
	}

	if s.Remaining()-i <= 0 {
		return err.New(
			"EOF encountered before string was terminated",
			err.Pos(line, col+i),
		)
	}

	return nil
}

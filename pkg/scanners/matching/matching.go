package matching

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	mat "github.com/PaulioRandall/scarlet-go/pkg/scanners/matching/matcher"
)

func ScanAll_(s string) ([]Token, error) {

	tks := []Token{}
	m := mat.New(s, patterns_)

	println("*****")

	for {

		v, e := m.Next()
		if e != nil {
			return nil, e
		}

		if v == nil {
			break
		}

		tk, _ := v.(Token)
		println(ToString(tk))
		tks = append(tks, tk)
	}

	return tks, nil
}

type pattern_ struct {
	m Morpheme
	f mat.PatternMatcher
}

func (p pattern_) Matcher() mat.PatternMatcher {
	return p.f
}

func (p pattern_) OnMatch(value string, line, col int) interface{} {
	return NewToken(p.m, value, line, col)
}

var patterns_ = []mat.Pattern{
	pattern_{NEWLINE, func(s *mat.Symbols) (int, error) {
		_, n := s.IsNewline(0)
		return n, nil
	}},
	pattern_{WHITESPACE, func(s *mat.Symbols) (int, error) {
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
	pattern_{COMMENT, func(s *mat.Symbols) (int, error) {

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

		return n, nil
	}},
	pattern_{MATCH, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "MATCH")
	}},
	pattern_{BOOL, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "FALSE")
	}},
	pattern_{BOOL, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "TRUE")
	}},
	pattern_{LIST, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "LIST")
	}},
	pattern_{LOOP, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "LOOP")
	}},
	pattern_{DEF, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "DEF")
	}},
	pattern_{FUNC, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "F")
	}},
	pattern_{EXPR_FUNC, func(s *mat.Symbols) (int, error) {
		return matchWord_(s, "E")
	}},
	pattern_{IDENTIFIER, func(s *mat.Symbols) (int, error) {
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
	pattern_{ASSIGN, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ":")
	}},
	pattern_{UPDATES, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "<-")
	}},
	pattern_{LIST_END, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ">>")
	}},
	pattern_{LIST_START, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "<<")
	}},
	pattern_{LESS_THAN_OR_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "<=")
	}},
	pattern_{MORE_THAN_OR_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ">=")
	}},
	pattern_{BLOCK_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "{")
	}},
	pattern_{BLOCK_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "}")
	}},
	pattern_{PAREN_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "(")
	}},
	pattern_{PAREN_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ")")
	}},
	pattern_{GUARD_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "[")
	}},
	pattern_{GUARD_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "]")
	}},
	pattern_{OUTPUT, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "^")
	}},
	pattern_{DELIMITER, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ",")
	}},
	pattern_{VOID, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "_")
	}},
	pattern_{TERMINATOR, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ";")
	}},
	pattern_{SPELL, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "@")
	}},
	pattern_{ADD, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "+")
	}},
	pattern_{SUBTRACT, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "-")
	}},
	pattern_{MULTIPLY, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "*")
	}},
	pattern_{DIVIDE, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "/")
	}},
	pattern_{REMAINDER, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "%")
	}},
	pattern_{AND, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "&")
	}},
	pattern_{OR, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "|")
	}},
	pattern_{EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "==")
	}},
	pattern_{NOT_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "!=")
	}},
	pattern_{LESS_THAN, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, "<")
	}},
	pattern_{MORE_THAN, func(s *mat.Symbols) (int, error) {
		return matchStr_(s, ">")
	}},
	pattern_{STRING, func(s *mat.Symbols) (int, error) {

		const (
			PREFIX     = '"'
			SUFFIX     = '"'
			ESCAPE     = '\\'
			SUFFIX_LEN = 1
		)

		ru, e := s.At(0)
		if e != nil || ru != PREFIX {
			return 0, e
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
			e := checkForMissingTermination_(s, i)
			return resume, e
		})

		return n + SUFFIX_LEN, nil
	}},
	pattern_{NUMBER, func(s *mat.Symbols) (int, error) {

		const (
			DELIM     = '.'
			DELIM_LEN = 1
		)

		n, e := matchInt_(s, 0)
		if e != nil {
			return 0, e
		}

		if n == 0 || s.Remaining()-n <= 0 {
			return n, nil
		}

		ru, e := s.At(n)
		if e != nil || ru != DELIM {
			return n, nil
		}

		frac, e := matchInt_(s, n+DELIM_LEN)
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

func matchStr_(s *mat.Symbols, str string) (int, error) {

	matched, e := s.Match(0, str)
	if e != nil || !matched {
		return 0, e
	}

	return len(str), nil
}

func matchWord_(s *mat.Symbols, word string) (int, error) {

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

	ru, e := s.At(size)
	if e != nil || unicode.IsLetter(ru) || ru == '_' {
		return 0, e
	}

	return size, nil
}

func matchInt_(s *mat.Symbols, start int) (int, error) {
	return s.CountWhile(start, func(_ int, ru rune) (bool, error) {
		return unicode.IsDigit(ru), nil
	})
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination_(s *mat.Symbols, i int) error {

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

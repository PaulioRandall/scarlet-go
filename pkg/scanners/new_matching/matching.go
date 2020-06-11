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
	ty TokenType
	f  mat.PatternMatcher
}

func (p pattern) Matcher() mat.PatternMatcher {
	return p.f
}

func (p pattern) OnMatch(value string, line, col int) interface{} {
	return NewToken(p.ty, value, line, col)
}

var patterns = []mat.Pattern{
	pattern{TK_NEWLINE, func(s *mat.Symbols) (int, error) {
		_, n := s.IsNewline(0)
		return n, nil
	}},
	pattern{TK_WHITESPACE, func(s *mat.Symbols) (int, error) {
		// Returns the number of consecutive whitespace terminals.
		// Newlines are not counted as whitespace.

		for i := 0; s.Has(i + 1); i++ {
			ru := s.At(i)

			if !unicode.IsSpace(ru) {
				return i, nil
			}

			if yes, _ := s.IsNewline(i); yes {
				return i, nil
			}
		}

		return s.Remaining(), nil
	}},
	pattern{TK_COMMENT, func(s *mat.Symbols) (int, error) {

		PREFIX := "//"
		PREFIX_LEN := 2

		matched, e := s.Match(0, PREFIX)
		if e != nil || !matched {
			return 0, e
		}

		n := PREFIX_LEN
		for exit := false; s.Has(n+1) && !exit; n++ {
			exit, _ = s.IsNewline(n)
		}

		return n, nil
	}},
	pattern{TK_WHEN, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "when")
	}},
	pattern{TK_WATCH, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "watch")
	}},
	pattern{TK_BOOL, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "false")
	}},
	pattern{TK_BOOL, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "true")
	}},
	pattern{TK_LOOP, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "loop")
	}},
	pattern{TK_DEFINITION, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "def")
	}},
	pattern{TK_FUNCTION, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "F")
	}},
	pattern{TK_EXPR_FUNC, func(s *mat.Symbols) (int, error) {
		return matchWord(s, "E")
	}},
	pattern{TK_IDENTIFIER, func(s *mat.Symbols) (int, error) {

		if !unicode.IsLetter(s.At(0)) {
			return 0, nil
		}

		i := 0
		for i++; s.Has(i + 1); i++ {
			ru := s.At(i)

			if !unicode.IsLetter(ru) && ru != '_' {
				break
			}
		}

		return i, nil
	}},
	pattern{TK_SPELL, func(s *mat.Symbols) (int, error) {

		if s.At(0) != '@' {
			return 0, nil
		}

		firstLetterError := func(i int) (int, error) {
			return 0, err.New(
				"Spell identifier or part must start with letter",
				err.Pos(s.Line(), s.Col()+i),
			)
		}

		i := 0

	SPELL_PART:

		i++
		if !s.Has(i + 1) {
			return firstLetterError(i)
		}

		if !unicode.IsLetter(s.At(i)) {
			return firstLetterError(i)
		}

		for i++; s.Has(i + 1); i++ {
			ru := s.At(i)

			if ru == '.' {
				goto SPELL_PART
			}

			if !unicode.IsLetter(ru) && ru != '_' {
				break
			}
		}

		return i, nil
	}},
	pattern{TK_ASSIGNMENT, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ":=")
	}},
	pattern{TK_LESS_THAN_OR_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "<=")
	}},
	pattern{TK_MORE_THAN_OR_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ">=")
	}},
	pattern{TK_THEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "->")
	}},
	pattern{TK_BLOCK_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "{")
	}},
	pattern{TK_BLOCK_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "}")
	}},
	pattern{TK_PAREN_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "(")
	}},
	pattern{TK_PAREN_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ")")
	}},
	pattern{TK_GUARD_OPEN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "[")
	}},
	pattern{TK_GUARD_CLOSE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "]")
	}},
	pattern{TK_OUTPUT, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "^")
	}},
	pattern{TK_DELIMITER, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ",")
	}},
	pattern{TK_VOID, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "_")
	}},
	pattern{TK_TERMINATOR, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ";")
	}},
	pattern{TK_PLUS, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "+")
	}},
	pattern{TK_MINUS, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "-")
	}},
	pattern{TK_MULTIPLY, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "*")
	}},
	pattern{TK_DIVIDE, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "/")
	}},
	pattern{TK_REMAINDER, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "%")
	}},
	pattern{TK_AND, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "&")
	}},
	pattern{TK_OR, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "|")
	}},
	pattern{TK_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "==")
	}},
	pattern{TK_NOT_EQUAL, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "!=")
	}},
	pattern{TK_LESS_THAN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, "<")
	}},
	pattern{TK_MORE_THAN, func(s *mat.Symbols) (int, error) {
		return matchStr(s, ">")
	}},
	pattern{TK_STRING, func(s *mat.Symbols) (int, error) {

		const (
			PREFIX     = '"'
			SUFFIX     = '"'
			ESCAPE     = '\\'
			SUFFIX_LEN = 1
		)

		if s.At(0) != PREFIX {
			return 0, nil
		}

		escaped := false // Init true to escape prefix

		i := 0
		for i++; s.Has(i + 1); i++ {
			ru := s.At(i)

			switch {
			case escaped:
				escaped = false

			case ru == SUFFIX:
				goto FINALLY

			case ru == ESCAPE:
				escaped = true
			}

			e := checkForMissingTermination(s, i)
			if e != nil {
				return 0, e
			}
		}

	FINALLY:
		return i + SUFFIX_LEN, nil
	}},
	pattern{TK_NUMBER, func(s *mat.Symbols) (int, error) {

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
			return 0, err.New(
				"Expected digit after decimal point",
				err.Pos(s.Line(), s.Col()+n),
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

	i := 0
	for s.Has(i+1) && unicode.IsDigit(s.At(i)) {
		i++
	}

	return i, nil
}

// checkForMissingTermination panics if a string or template is found to be
// unterminated.
func checkForMissingTermination(s *mat.Symbols, i int) error {

	if ok, _ := s.IsNewline(i); ok {
		return err.New(
			"Newline encountered before string was terminated",
			err.Pos(s.Line(), s.Col()+i),
		)
	}

	if s.Remaining()-i <= 0 {
		return err.New(
			"EOF encountered before string was terminated",
			err.Pos(s.Line(), s.Col()+i),
		)
	}

	return nil
}

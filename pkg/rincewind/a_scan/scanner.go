package scan

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

type scanner struct {
	itr  RuneItr
	buff rune
	line int
	col  int
}

func (scn *scanner) scan() (token.Tok, ScanFunc, error) {

	tk := token.Tok{}
	line, col := scn.line, scn.col

	e := next(scn, &tk)
	if e != nil {
		return token.Tok{}, nil, e
	}

	tk.Line = line
	tk.ColBegin = col
	tk.ColEnd = scn.col

	if tk.Sub == SUB_NEWLINE {
		scn.line++
		scn.col = 0
	}

	if scn.empty() {
		return tk, nil, nil
	}

	return tk, scn.scan, nil
}

func (scn *scanner) bufferNext() {

	buff, ok := scn.itr.Next()

	if ok || scn.buff != rune(0) {
		scn.col++
	}

	if ok {
		scn.buff = buff
		return
	}

	scn.buff = rune(0)
}

func (scn *scanner) hasNext() bool {
	return scn.buff != rune(0)
}

func (scn *scanner) empty() bool {
	return scn.buff == rune(0)
}

func (scn *scanner) next() rune {

	if scn.empty() {
		failNow("No symbols remain, should call `match`, `hasNext`, or `empty` first")
	}

	r := scn.buff
	scn.bufferNext()

	return r
}

func (scn *scanner) match(ru rune) bool {
	return scn.buff == ru
}

func (scn *scanner) notMatch(ru rune) bool {
	return !scn.match(ru)
}

func (scn *scanner) matchNewline() bool {
	return scn.buff == '\r' || scn.buff == '\n'
}

func (scn *scanner) matchSpace() bool {
	return unicode.IsSpace(scn.buff)
}

func (scn *scanner) matchLetter() bool {
	return unicode.IsLetter(scn.buff)
}

func (scn *scanner) matchDigit() bool {
	return unicode.IsDigit(scn.buff)
}

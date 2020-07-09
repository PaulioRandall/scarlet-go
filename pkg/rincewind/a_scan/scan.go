package scan

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

type SymbolItr interface {
	Next() (rune, bool)
}

type ScanFunc func() (token.Tok, ScanFunc, error)

func New(itr SymbolItr) ScanFunc {

	if itr == nil {
		perror.Panic("Non-nil SymbolItr required")
	}

	scn := &scanner{
		itr: itr,
		col: -1, // -1 so index is before first symbol
	}
	scn.bufferNext()

	if scn.empty() {
		return nil
	}

	return scn.scan
}

type scanner struct {
	itr  SymbolItr
	buff rune
	line int
	col  int
}

func (scn *scanner) scan() (token.Tok, ScanFunc, error) {

	tk := token.Tok{}
	line, col := scn.line, scn.col

	e := scanNext(scn, &tk)
	if e != nil {
		return token.Tok{}, nil, e
	}

	tk.Line = line
	tk.ColBegin = col
	tk.ColEnd = scn.col

	if tk.Sub == SU_NEWLINE {
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
		perror.Panic(
			"No symbols remain, should call `match`, `hasNext`, or `empty` first")
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

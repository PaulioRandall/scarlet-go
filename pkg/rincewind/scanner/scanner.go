package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type SymbolItr interface {
	Next() (rune, bool)
}

type ScanFunc func() (tok, ScanFunc, error)

func New(itr SymbolItr) ScanFunc {

	if itr == nil {
		perror.ProgPanic("Non-nil SymbolItr required")
	}

	scn := &scanner{itr: itr}
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

func (scn *scanner) bufferNext() {

	var ok bool
	scn.buff, ok = scn.itr.Next()

	if !ok {
		scn.buff = rune(0)
		return
	}

	scn.col++
}

func (scn *scanner) hasNext() bool {
	return scn.buff != rune(0)
}

func (scn *scanner) empty() bool {
	return scn.buff == rune(0)
}

func (scn *scanner) next() rune {

	if scn.empty() {
		perror.ProgPanic("No symbols remain, should call `hasNext` or `empty` first")
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

func (scn *scanner) scan() (tok, ScanFunc, error) {

	if scn.empty() {
		return tok{}, nil, nil
	}

	tk := tok{}
	line, col := scn.line, scn.col

	e := scanNext(scn, &tk)
	if e != nil {
		return tok{}, nil, e
	}

	tk.line = line
	tk.colBegin = col
	tk.colEnd = scn.col
	checkTok(tk, line, col) // TODO: Remove once tests have been created

	if tk.st == ST_NEWLINE {
		scn.line++
		scn.col = 0
	}

	if scn.empty() {
		return tk, nil, nil
	}

	return tk, scn.scan, nil
}

// TODO: Remove once tests have been created
func checkTok(tk tok, line, col int) {

	if tk.gt == GT_UNDEFINED {
		perror.ProgPanic("Missing GenType for token at %d:%d...", line, col)
	}

	if tk.st == ST_UNDEFINED {
		perror.ProgPanic("Missing SubType for token at %d:%d...", line, col)
	}

	if tk.raw == "" {
		perror.ProgPanic("Missing raw text for token at %d:%d...", line, col)
	}

	if tk.val == nil {
		perror.ProgPanic("Missing value for token at %d:%d...", line, col)
	}

	if tk.val == nil {
		perror.ProgPanic("Missing value for token at %d:%d...", line, col)
	}
}

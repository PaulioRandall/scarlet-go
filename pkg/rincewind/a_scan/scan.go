package scan

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type ScanFunc func() (token.Tok, ScanFunc, error)

type RuneStream interface {
	Next() (rune, bool)
}

type runeStream struct {
	runes []rune
	size  int
	i     int
}

func (rs *runeStream) Next() (rune, bool) {

	if rs.i >= rs.size {
		return rune(0), false
	}

	ru := rs.runes[rs.i]
	rs.i++
	return ru, true
}

func New(rs RuneStream) ScanFunc {

	if rs == nil {
		failNow("Non-nil RuneItr required")
	}

	scn := &scanner{
		rs:  rs,
		col: -1, // -1 so index is before the first symbol
	}
	scn.bufferNext()

	if scn.empty() {
		return nil
	}

	return scn.scan
}

func StreamAll(rs RuneStream) ([]token.Token, error) {

	var (
		e   error
		tk  token.Token
		tks = []token.Token{}
	)

	for f := New(rs); f != nil; {
		if tk, f, e = f(); e != nil {
			return nil, e
		}

		tks = append(tks, tk)
	}

	return tks, nil
}

func ScanAll(s string) ([]token.Token, error) {
	return StreamAll(&runeStream{
		runes: []rune(s),
		size:  len(s),
	})
}

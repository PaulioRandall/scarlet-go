package format

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func format(tks []token.Token) error {

	fmtr := &formatter{
		ts: token.NewStream(tks),
	}
	fmtr.bufferNext()

	// TODO
	return nil
}

type formatter struct {
	token.Queue
	ts                *token.Stream
	buff              token.Token
	indent            int
	wasTerminator     bool
	wasEmptyStatement bool
}

func (f *formatter) bufferNext() {
	f.buff = f.ts.Next()
}

func (f *formatter) empty() bool {
	return f.buff == nil && f.Empty()
}

func (f *formatter) match(props ...Prop) bool {

	if f.buff == nil {
		return false
	}

	return f.buff.Is(props...)
}

func (f *formatter) putBuff() {
	f.Put(f.buff)
	f.bufferNext()
}

func (f *formatter) putSpace() {

	tk := token.Tok{
		RawProps: []Prop{PR_REDUNDANT, PR_WHITESPACE},
		RawStr:   " ",
	}

	f.Put(tk)
}

func (f *formatter) putNewline() {

	tk := token.Tok{
		RawProps: []Prop{PR_TERMINATOR, PR_NEWLINE},
		RawStr:   "\n",
	}

	f.Put(tk)
}

func (f *formatter) accept(props ...Prop) bool {

	if f.match(props...) {
		f.putBuff()
		return true
	}

	return false
}

func (f *formatter) ignore(props ...Prop) bool {

	if f.match(props...) {
		f.bufferNext()
		return true
	}

	return false
}

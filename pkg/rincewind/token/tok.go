package token

import (
	"fmt"
)

func MagicToken(
	gen GenType,
	sub SubType,
	tk Token,
) Token {

	r := tok{
		gen: gen,
		sub: sub,
		raw: tk.Raw(),
	}

	r.lineBegin, r.colBegin = tk.Begin()
	r.lineEnd, r.colEnd = tk.End()

	return r
}

type tok struct {
	gen       GenType
	sub       SubType
	raw       string
	lineBegin int
	colBegin  int
	lineEnd   int
	colEnd    int
}

func (tk tok) GenType() GenType {
	return tk.gen
}

func (tk tok) SubType() SubType {
	return tk.sub
}

func (tk tok) Raw() string {
	return tk.raw
}

func (tk tok) Value() string {

	switch {
	case tk.gen == GE_SPELL:
		return tk.raw[1:]

	case tk.sub == SU_STRING:
		if len(tk.raw) == 2 {
			return ""
		}

		return tk.raw[1 : len(tk.raw)-1]
	}

	return tk.raw
}

func (tk tok) Begin() (int, int) {
	return tk.lineBegin, tk.colBegin
}

func (tk tok) End() (int, int) {
	return tk.lineEnd, tk.colEnd
}

func (tk tok) String() string {

	// +1 converts from line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s:%s %q`,
		tk.lineBegin+1,
		tk.colBegin,
		tk.lineEnd+1,
		tk.colEnd,
		tk.gen.String(),
		tk.sub.String(),
		tk.Value(),
	)
}

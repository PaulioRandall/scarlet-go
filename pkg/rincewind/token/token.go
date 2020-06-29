package token

import (
	"fmt"
)

type Token interface {
	fmt.Stringer
	GenType() GenType
	SubType() SubType
	Raw() string
	Val() interface{}
	Begin() (int, int)
	End() (int, int)
}

func NewToken(
	gt GenType,
	st SubType,
	raw string,
	val interface{},
	line, col int,
) Token {

	return tok{
		gt:       gt,
		st:       st,
		raw:      raw,
		val:      val,
		line:     line,
		colBegin: col,
		colEnd:   col + len(raw),
	}
}

type tok struct {
	gt       GenType
	st       SubType
	raw      string
	val      interface{}
	line     int
	colBegin int
	colEnd   int
}

func (tk tok) GenType() GenType {
	return tk.gt
}

func (tk tok) SubType() SubType {
	return tk.st
}

func (tk tok) Raw() string {
	return tk.raw
}

func (tk tok) Val() interface{} {
	return tk.val
}

func (tk tok) Begin() (int, int) {
	return tk.line, tk.colBegin
}

func (tk tok) End() (int, int) {
	return tk.line, tk.colEnd
}

func (tk tok) String() string {

	// +1 converts from line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s:%s %q`,
		tk.line+1,
		tk.colBegin,
		tk.line+1,
		tk.colEnd,
		tk.gt.String(),
		tk.st.String(),
		tk.val,
	)
}

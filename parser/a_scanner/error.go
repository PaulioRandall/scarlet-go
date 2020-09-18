package scanner

import (
	"fmt"
)

type scanErr struct {
	lineIdx, colIdx int
	msg             string
}

func newErr(lineIdx, colIdx int, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return scanErr{
		lineIdx: lineIdx,
		colIdx:  colIdx,
		msg:     msg,
	}
}

func (e scanErr) Error() string {
	return e.msg
}

func (e scanErr) Pos() (lineIdx, colIdx int) {
	return e.lineIdx, e.colIdx
}

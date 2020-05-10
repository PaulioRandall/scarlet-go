package matching

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
)

type scanErr struct {
	msg       string
	lineIndex int
	colIndex  int
	length    int
}

func newErr(s *symbols, colOffset int, msg string, args ...interface{}) err.Err {
	return &scanErr{
		lineIndex: s.line,
		colIndex:  s.col + colOffset,
		msg:       fmt.Sprintf(msg, args...),
	}
}

func (se scanErr) Error() string {
	return se.msg
}

func (se scanErr) Cause() error {
	return nil
}

func (se scanErr) LineIndex() int {
	return se.lineIndex
}

func (se scanErr) ColIndex() int {
	return se.colIndex
}

func (se scanErr) Length() int {
	return se.length
}

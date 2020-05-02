package recursive

import (
	"fmt"

	e "github.com/PaulioRandall/scarlet-go/pkg/err"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type parseErr struct {
	msg       string
	cause     error
	lineIndex int
	colIndex  int
	length    int
}

func err(f string, tk token.Token, msg string, args ...interface{}) e.Err {

	s := "[parser." + f + "] " + fmt.Sprintf(msg, args...)

	return &parseErr{
		msg:       s,
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

func unexpected(f string, tk token.Token, expected token.TokenType) e.Err {
	return err(f, tk, "Expected %v, got %s", expected, tk.String())
}

func (pe parseErr) Error() string {
	return pe.msg
}

func (pe parseErr) Cause() error {
	return pe.cause
}

func (pe parseErr) LineIndex() int {
	return pe.lineIndex
}

func (pe parseErr) ColIndex() int {
	return pe.colIndex
}

func (pe parseErr) Length() int {
	return pe.length
}

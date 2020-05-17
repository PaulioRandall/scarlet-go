package recursive

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func ParseAll(tks []token.Token) []st.Statement {
	p := &pipe{token.NewIterator(tks)}
	return parseStatements(p)
}

type parseErr struct {
	msg  string
	line int
	col  int
	len  int
}

func err(f string, tk token.Token, offset int, msg string, args ...interface{}) error {
	return &parseErr{
		msg:  "[parser." + f + "] " + fmt.Sprintf(msg, args...),
		line: tk.Line,
		col:  tk.Col + offset,
		len:  len(tk.Value),
	}
}

func unexpected(f string, tk token.Token, expected token.TokenType) error {
	return err(f, tk, 0, "Expected %v, got %s", expected, tk.String())
}

func (pe parseErr) Error() string {
	return pe.msg
}

func (pe parseErr) Line() int {
	return pe.line
}

func (pe parseErr) Col() int {
	return pe.col
}

func (pe parseErr) Len() int {
	return pe.len
}

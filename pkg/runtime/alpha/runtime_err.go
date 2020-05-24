package alpha

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type runtimeErr struct {
	msg  string
	line int
	col  int
	len  int
}

func err(f string, tk Token, msg string, args ...interface{}) error {
	return runtimeErr{
		msg:  "[runtime." + f + "] " + fmt.Sprintf(msg, args...),
		line: tk.Line(),
		col:  tk.Col(),
		len:  len(tk.Value()),
	}
}

func (re runtimeErr) Error() string {
	return re.msg
}

func (re runtimeErr) Line() int {
	return re.line
}

func (re runtimeErr) Col() int {
	return re.col
}

func (re runtimeErr) Len() int {
	return re.len
}

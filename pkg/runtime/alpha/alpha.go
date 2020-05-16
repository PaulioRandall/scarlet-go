package alpha

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func Run(ss []st.Statement) alphaContext {
	ctx := alphaContext{
		true,
		make(map[string]result),
		make(map[string]result),
		nil,
	}

	exeStatements(&ctx, ss)
	return ctx
}

type runtimeErr struct {
	msg  string
	line int
	col  int
	len  int
}

func err(f string, tk token.Token, msg string, args ...interface{}) error {
	return runtimeErr{
		msg:  "[runtime." + f + "] " + fmt.Sprintf(msg, args...),
		line: tk.Line,
		col:  tk.Col,
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

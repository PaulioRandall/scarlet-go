package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/runtime/alpha"
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

type Context interface {
	String() string
}

type Method string

const (
	DEFAULT Method = `DEFAULT_RUNTIME`
	ALPHA   Method = `ALPHA_RUNTIME`
)

func Run(s []st.Statement, m Method) Context {

	switch m {
	case DEFAULT, ALPHA:
		return alpha.Run(s)
	}

	panic(string(`Unknown runtime method '` + m + `'`))
}

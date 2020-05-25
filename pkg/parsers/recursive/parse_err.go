package recursive

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func errMsg(f string, exp string, act Token) string {
	return fmt.Sprintf(
		"[parser.%s] Expected %v, but got (%s)",
		f, exp, ToString(act),
	)
}

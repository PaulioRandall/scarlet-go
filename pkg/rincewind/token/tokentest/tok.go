package tokentest

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type tok struct {
	ge       GenType
	su       SubType
	raw      string
	line     int
	colBegin int
	colEnd   int
}

func (tk tok) GenType() GenType {
	return tk.ge
}

func (tk tok) SubType() SubType {
	return tk.su
}

func (tk tok) Raw() string {
	return tk.raw
}

func (tk tok) Value() string {

	switch {
	case tk.ge == GE_SPELL:
		return tk.raw[1:]

	case tk.su == SU_STRING:
		if len(tk.raw) == 2 {
			return ""
		}

		return tk.raw[1 : len(tk.raw)-1]
	}

	return tk.raw
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
		tk.ge.String(),
		tk.su.String(),
		tk.Value(),
	)
}

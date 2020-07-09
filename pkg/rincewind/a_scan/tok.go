package scan

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

type tok struct {
	gen        GenType
	sub        SubType
	raw        string
	line       int
	begin, end int
}

func (tk tok) GenType() GenType {
	return tk.gen
}

func (tk tok) SubType() SubType {
	return tk.sub
}

func (tk tok) Raw() string {
	return tk.raw
}

func (tk tok) Value() string {

	switch {
	case tk.gen == GE_SPELL:
		return tk.raw[1:]

	case tk.sub == SU_STRING:
		if len(tk.raw) == 2 {
			return ""
		}

		return tk.raw[1 : len(tk.raw)-1]
	}

	return tk.raw
}

func (tk tok) Begin() (int, int) {
	return tk.line, tk.begin
}

func (tk tok) End() (int, int) {
	return tk.line, tk.end
}

func (tk tok) String() string {

	// +1 converts from line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s:%s %q`,
		tk.line+1,
		tk.begin,
		tk.line+1,
		tk.end,
		tk.gen.String(),
		tk.sub.String(),
		tk.Value(),
	)
}

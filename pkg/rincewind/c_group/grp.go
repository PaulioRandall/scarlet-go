package group

import (
	"fmt"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/stat"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type grp struct {
	st  StatType
	tks []Token
}

func (g grp) Type() StatType {
	return g.st
}

func (g grp) Tokens() []Token {
	return g.tks
}

func (g grp) Begin() (int, int) {
	return g.tks[0].Begin()
}

func (g grp) End() (int, int) {
	lastIdx := len(g.tks) - 1
	return g.tks[lastIdx].End()
}

func (g grp) String() string {

	sb := strings.Builder{}

	line, col := g.Begin()
	endLine, endCol := g.End()
	s := fmt.Sprintf(`%d:%d %d:%d `, line, col, endLine, endCol)
	sb.WriteString(s)

	for _, tk := range g.tks {
		sb.WriteString(tk.GenType().String())
	}

	return sb.String()
}

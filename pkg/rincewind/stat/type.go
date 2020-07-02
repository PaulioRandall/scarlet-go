package stat

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type StatType int

func (st StatType) String() string {

	s, ok := types[st]

	if !ok {
		s = `ST_UNDEFINED`
	}

	return s
}

const (
	ST_UNDEFINED StatType = iota
	ST_SPELL_CALL
)

var types map[StatType]string = map[StatType]string{
	ST_UNDEFINED:  `ST_UNDEFINED`,
	ST_SPELL_CALL: `ST_SPELL_CALL`,
}

type Group interface {
	Type() StatType
	Tokens() []Token
	fmt.Stringer
	Snippet
}

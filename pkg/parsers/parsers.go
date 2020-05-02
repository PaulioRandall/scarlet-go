package parsers

import (
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/PaulioRandall/scarlet-go/pkg/parsers/recursive"
)

type Method string

const (
	DEFAULT          Method = `DEFAULT_PARSER`
	RECURSIVE_DECENT Method = `RECURSIVE_DECENT_PARSER`
)

func ParseAll(tks []token.Token, m Method) []statement.Statement {

	switch m {
	case DEFAULT, RECURSIVE_DECENT:
		return recursive.ParseAll(tks)
	}

	panic(string(`Unknown parsing method '` + m + `'`))
}

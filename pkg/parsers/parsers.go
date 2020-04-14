package parsers

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/statement"

	"github.com/PaulioRandall/scarlet-go/pkg/parsers/recursive"
)

// Method represents a scanning method
type Method string

const (
	DEFAULT          Method = `DEFAULT_PARSER`
	RECURSIVE_DECENT Method = `RECURSIVE_DECENT_PARSER`
)

func ParseAll(tks []lexeme.Token, m Method) statement.Statements {

	switch m {
	case DEFAULT, RECURSIVE_DECENT:
		return recursive.ParseAll(tks)
	}

	panic(string(`Unknown parsing method '` + m + `'`))
}

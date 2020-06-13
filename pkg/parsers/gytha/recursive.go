package gytha

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func ParseAll(tks []Token) []Statement {
	p := &pipe{NewIterator(tks)}
	return parseStatements(p)
}

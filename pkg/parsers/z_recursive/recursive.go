package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func ParseAll(tks []Token) []Statement {
	p := &pipe{NewIterator(tks)}
	return parseStatements(p)
}

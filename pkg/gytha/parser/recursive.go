package parser

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func ParseAll(tks []Token) []Statement {
	p := &pipe{NewIterator(tks)}
	return parseStatements(p)
}

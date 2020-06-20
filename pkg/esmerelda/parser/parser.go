package parser

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type TokenStream interface {
	Next() Token
	Peek() Token
	PeekBeyond() Token
}

func ParseStatements(ts TokenStream) ([]Expression, error) {
	p := newPipeline(ts)
	return statements(p)
}

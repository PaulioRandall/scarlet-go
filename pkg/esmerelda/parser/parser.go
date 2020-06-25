package parser

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type TokenStream interface {
	Next() Token
	Peek() Token
	PeekBeyond() Token
}

func ParseStatements(ts TokenStream) ([]Expr, error) {
	p := newPipeline(ts)
	return statements(p)
}

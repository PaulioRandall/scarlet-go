package parser

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type ParseFunc func() (Expr, ParseFunc, error)

type TokenStream interface {
	Next() Token
}

func New(ts TokenStream) ParseFunc {

	if ts == nil {
		panic("PROGRAMMERS ERROR! TokenStream is nil")
	}

	p := newPipeline(ts)

	if p.empty() {
		return nil
	}

	par := &parser{p}
	return par.parse
}

type parser struct {
	p *pipeline
}

func (par *parser) parse() (Expr, ParseFunc, error) {

	expr, e := statement(par.p)
	if e != nil {
		return nil, nil, e
	}

	if par.p.empty() {
		return expr, nil, nil
	}

	return expr, par.parse, nil
}

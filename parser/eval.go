package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/token"
)

// Expr represents an expression that produces a value when evaluated.
type Expr func(c ctx.Context) (ctx.Value, error)

func NewForValue(v ctx.Value) Expr {
	return func(ctx.Context) (ctx.Value, error) {
		return v, nil
	}
}

func NewForID(t token.Token) Expr {
	return func(c ctx.Context) (ctx.Value, error) {
		return c.Get(t.Value), nil
	}
}

func NewForListAccess(idEv, indexEv Expr) Expr {
	return func(c ctx.Context) (_ ctx.Value, _ error) {
		// TODO:
		return
	}
}

func NewForOperator(t token.Token) Expr {
	return func(c ctx.Context) (_ ctx.Value, _ error) {
		// TODO:
		return
	}
}

func NewForFuncCall(id Expr, params, returns []Expr, body []Expr) Expr {
	return func(parent ctx.Context) (_ ctx.Value, _ error) {
		// TODO:
		return
	}
}

func NewForSpellCall(id Expr, params []Expr) Expr {
	return func(parent ctx.Context) (_ ctx.Value, _ error) {
		// TODO:
		return
	}
}

func NewForAssign(ids, ex []Expr) Expr {
	return nil
}

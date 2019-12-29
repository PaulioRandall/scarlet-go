package parser

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/parser/ctx"
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// LITERAL          := BOOL | INT | REAL | STRING | TEMPLATE .
func matchLiteral(tc *TokenCollector) (_ eval.Expr, _ int) {

	t := tc.Read()
	var v ctx.Value

	switch t.Kind {
	case token.BOOL_LITERAL:
		v = parseBoolValue(t)
	case token.INT_LITERAL:
		fallthrough
	case token.REAL_LITERAL:
		v = parseNumValue(t)
	case token.STR_LITERAL:
		fallthrough
	case token.STR_TEMPLATE:
		// TODO: Treating as a string literal for now
		v = ctx.NewValue(ctx.STR, t.Value)
	default:
		tc.Unread(1)
		return
	}

	return eval.NewForValue(v), 1
}

// parseNumValue parses a number value and panics if the string cannot be
// parsed.
func parseNumValue(t token.Token) ctx.Value {

	n, e := strconv.ParseFloat(t.Value, 64)

	if e != nil {
		panic(NewParseErr("Could not parse number literal token value", e, t))
	}

	return ctx.NewValue(ctx.REAL, n)
}

// parseBoolValue parses a boolean value and panics if the string cannot be
// parsed.
func parseBoolValue(t token.Token) ctx.Value {

	var b bool

	if t.Value == `TRUE` {
		b = true
	} else if t.Value == `FALSE` {
		b = false
	} else {
		panic(NewParseErr("Could not parse boolean literal token value", nil, t))
	}

	return ctx.NewValue(ctx.BOOL, b)
}

package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// PARAM            := "\_" | ID_OR_ITEM | LITERAL .
func matchParam(tc *TokenCollector) (_ eval.Expr, _ int) {

	ex, n := matchIdOrItem(tc)
	if ex != nil {
		return ex, n
	}

	ex, n = matchLiteral(tc)
	if ex != nil {
		return ex, n
	}

	t := tc.Read()
	if t.Kind == token.VOID {
		return eval.NewForID(t), 1
	}

	tc.PutBack(1)
	return
}

// PARAM_LIST       := PARAM { "," PARAM } .
func matchParamList(tc *TokenCollector) (_ []eval.Expr, _ int) {

	ex, n := matchParam(tc)
	if ex == nil {
		return
	}

	params := []eval.Expr{ex}

	for {
		if tc.Read().Kind != token.DELIM {
			tc.PutBack(1)
			break
		}

		var i int
		ex, i = matchParam(tc)
		if ex == nil {
			panic(NewParseErr("Expected another parameter", nil, tc.Peek()))
		}

		params = append(params, ex)
		n += i + 1 // +1 for the delimiter
	}

	return params, n
}

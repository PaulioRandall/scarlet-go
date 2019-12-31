package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// parseAssign parses the next statement as an assignment statement. Assumes
// that the statement matches a valid assignment statement.
func parseAssign(tc *TokenCollector) (ids []eval.Expr) {

	_ = matchAssignStart(tc)

	for _, t := range tc.Take() {

		if t.Kind == token.ID || t.Kind == token.VOID {
			ids = append(ids, eval.NewForID(t))
		}
	}

	if 0 != matchAny(tc, token.FUNC) {
		tc.Clear(1)
		_ = parseFuncDef(tc)
		//_ = parseFuncBody(tc)
	}

	return
}

func parseFuncDef(tc *TokenCollector) (f eval.Expr) {

	// TODO: Match functions shouldn't increment the token index

	if 0 == matchAny(tc, token.OPEN_PAREN) {
		panic(NewParseErr("Expected opening parenthesis token", nil, tc.Peek()))
	}

	var params []eval.Expr
	if 0 != matchFuncParams(tc) {
		for _, t := range tc.Take() {
			if t.Kind == token.ID {
				params = append(params, eval.NewForID(t))
			}
		}
	}

	var returns []eval.Expr
	if 0 != matchFuncReturns(tc) {
		for _, t := range tc.Take() {
			if t.Kind == token.ID {
				returns = append(returns, eval.NewForID(t))
			}
		}
	}

	if 0 == matchAny(tc, token.CLOSE_PAREN) {
		panic(NewParseErr("Expected closing parenthesis token", nil, tc.Peek()))
	}

	return // TODO
}

// matchAssignStart := matchAssignIds ":=".
func matchAssignStart(tc *TokenCollector) (_ int) {

	n := matchAssignIds(tc)

	if 0 == n {
		return
	}

	if 0 == matchAny(tc, token.ASSIGN) {
		panic(NewParseErr("Expected ASSIGN token", nil, tc.Peek()))
	}

	return 1 + n
}

// ID_OR_VOID       := ID | "_" .
// matchAssignIds   := [ ID_OR_VOID { "," ID_OR_VOID } ].
func matchAssignIds(tc *TokenCollector) (n int) {

	if 0 == matchAny(tc, token.ID, token.VOID) {
		return
	}

	n++

	for 1 == matchAny(tc, token.DELIM) {
		n++

		if 0 == matchAny(tc, token.ID, token.VOID) {
			panic(NewParseErr("Expected ID token", nil, tc.Peek()))
		}

		n++
	}

	return n
}

// matchFuncParams   := [ ID { "," ID } ].
func matchFuncParams(tc *TokenCollector) (n int) {

	if 0 == matchAny(tc, token.ID) {
		return
	}

	n++

	for 1 == matchAny(tc, token.DELIM) {
		n++

		if 0 == matchAny(tc, token.ID) {
			panic(NewParseErr("Expected ID token", nil, tc.Peek()))
		}

		n++
	}

	return
}

// matchFuncReturns := [ "->" ID { "," ID } ].
func matchFuncReturns(tc *TokenCollector) (n int) {

	if 0 == matchAny(tc, token.RETURNS) {
		return
	}

	n++

	if 0 == matchAny(tc, token.ID) {
		panic(NewParseErr("Expected ID token", nil, tc.Peek()))
	}

	for 1 == matchAny(tc, token.DELIM) {
		n++

		if 0 == matchAny(tc, token.ID) {
			panic(NewParseErr("Expected ID token", nil, tc.Peek()))
		}

		n++
	}

	return
}

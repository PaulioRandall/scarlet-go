package parser

// PROGRAM          := STATEMENT BLOCK .
// BLOCK            := { STATEMENT } .
// STATEMENT        := ( ASSIGNMENT | INLINE_EXPR | GUARD | MATCH_BLOCK | WATCH_BLOCK ) NEWLINE .
// INLINE_STATEMENT := ( ASSIGNMENT | INLINE_EXPR ) NEWLINE .
// EXPR             := ID_OR_ITEM | INLINE_EXPR .
// INLINE_EXPR      := LITERAL | FUNC_CALL | SPELL | OPERATION .
// ASSIGNMENT       := [ "GLOBAL" ] ID_ARRAY ":=" ( LIST | EXPR | FUNC ) .
// FUNC             := "F" "(" PARAM_LIST [ "->" ID_ARRAY ] ")" BODY .
// GUARD            := "[" EXPR "]" BODY .
// OPERATION        := OPERAND OPERATOR { OPERAND OPERATOR } OPERAND .
// OPERAND          := [ "~" | "¬" ] ( ID_OR_ITEM | LITERAL | FUNC_CALL | SPELL ) .
// BODY             := INLINE_STATEMENT | ( "DO" NEWLINE BLOCK "END" ) .
// MATCH_BLOCK      := "MATCH" NEWLINE MATCH_CASE { MATCH_CASE } "END" .
// MATCH_CASE       := EXPR BODY NEWLINE .
// WATCH_BLOCK      := "WATCH" ID { "," ID } NEWLINE BLOCK "END" .
// LIST             := "{" LIST_ITEMS [ "," [ NEWLINE ] ] "}" .
// LIST_ITEMS       := EXPR { "," [ NEWLINE ] EXPR } .

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// OPERATOR         := NUM_OPERATOR | BOOL_OPERATOR | CMP_OPERATOR .
// NUM_OPERATOR     := "+" | "-" | "\*" | "/" | "%" .
// BOOL_OPERATOR    := "|" | "&" .
// CMP_OPERATOR     := "=" | "#" | "<" | ">" | "<=" | ">=" .
func matchOperator(tc *TokenCollector) (_ eval.Expr, _ int) {

	t := tc.Read()

	if t.Kind != token.OPERATOR {
		tc.Unread(1)
		return
	}

	return eval.NewForOperator(t), 1
}

// FUNC_CALL        := ID "(" PARAM_LIST ")" .
func matchFuncCall(tc *TokenCollector) (_ eval.Expr, _ int) {

	exs, n := matchCall(tc)
	if exs == nil {
		return
	}

	return eval.NewForFuncCall(exs[0], exs[1:]), n
}

// SPELL            := "@" FUNC_CALL .
func matchSpellCall(tc *TokenCollector) (_ eval.Expr, _ int) {

	n, t := 1, tc.Read()
	if t.Kind != token.SPELL {
		tc.Unread(n)
		return
	}

	exs, i := matchCall(tc)
	if exs == nil {
		panic(NewParseErr("Expected ID and parameter tokens", nil, tc.Peek()))
	}

	n += i
	return eval.NewForSpellCall(exs[0], exs[1:]), n
}

// FUNC_CALL        := ID "(" PARAM_LIST ")" .
func matchCall(tc *TokenCollector) (_ []eval.Expr, _ int) {

	var idExpr eval.Expr
	var paramExpr []eval.Expr
	var ex []eval.Expr
	var i int

	n, t := 1, tc.Read()
	if t.Kind != token.ID {
		goto NO_MATCH
	}

	n, t = n+1, tc.Read()
	if t.Kind != token.OPEN_PAREN {
		goto NO_MATCH
	}

	idExpr = eval.NewForID(t)
	paramExpr, i = matchParamList(tc)

	n, t = n+1, tc.Read()
	if t.Kind != token.CLOSE_PAREN {
		panic(NewParseErr("Expected closing parentheses", nil, t))
	}

	ex = []eval.Expr{idExpr}
	if paramExpr != nil {
		n += i
		ex = append(ex, paramExpr...)
	}

	return ex, n

NO_MATCH:
	tc.Unread(n)
	return
}

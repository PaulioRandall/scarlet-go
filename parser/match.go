package parser

// COMMENT          := "//" * Any visible unicode character * NEWLINE .
// PROGRAM          := STATEMENT BLOCK .
// BLOCK            := { STATEMENT } .
// STATEMENT        := ( ASSIGNMENT | INLINE_EXPR | GUARD | MATCH_BLOCK | WATCH_BLOCK ) NEWLINE .
// INLINE_STATEMENT := ( ASSIGNMENT | INLINE_EXPR ) NEWLINE .
// EXPR             := ID_OR_ITEM | INLINE_EXPR .
// INLINE_EXPR      := LITERAL | FUNC_CALL | SPELL | OPERATION .
// SPELL            := "@" FUNC_CALL .
// FUNC_CALL        := ID "(" PARAM_LIST ")" .
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
	"github.com/PaulioRandall/scarlet-go/token"
)

// matchSeq attempts to pattern match the sequence of token kinds to the tokens
// read by the TokenCollector. If they all match the number of tokens is
// returned else zero is returned.
func matchSeq(tc *TokenCollector, ks ...token.Kind) (_ int) {

	var t token.Token
	var n int

	for _, k := range ks {
		n, t = n+1, tc.Read()

		if t.Kind != k {
			tc.Unread(n)
			return
		}
	}

	return n
}

// matchAny attempts to match any of the token kinds with the next token read
// by the TokenCollector. If a match is found 1 is returned else 0 is.
func matchAny(tc *TokenCollector, ks ...token.Kind) (_ int) {

	t := tc.Read()

	for _, k := range ks {
		if t.Kind == k {
			return 1
		}
	}

	tc.Unread(1)
	return 0
}

// matchEither calls each matcher function in order until a match is found or
// non-match. If one matches the number of tokens returned by the matcher is
// returned else zero is returned.
func matchEither(tc *TokenCollector, ms ...matcher) (n int) {

	for _, m := range ms {
		n = m(tc)

		if n > 0 {
			break
		}
	}

	return
}

// matchItemAccess  := "[" ( ID | INT) ) "]" .
func matchItemAccess(tc *TokenCollector) (_ int) {

	n := matchAny(tc, token.OPEN_GUARD)
	n += matchAny(tc, token.ID, token.INT_LITERAL)
	n += matchAny(tc, token.CLOSE_GUARD)

	if n != 3 {
		tc.Unread(n)
		return
	}

	return n
}

// matchIdOrItem   := ID [ matchItemAccess ] .
func matchIdOrItem(tc *TokenCollector) (_ int) {

	if 0 == matchAny(tc, token.ID) {
		return
	}

	return 1 + matchItemAccess(tc)
}

// matchLiteral     := BOOL | INT | REAL | STRING | TEMPLATE .
func matchLiteral(tc *TokenCollector) int {
	return matchAny(tc,
		token.BOOL_LITERAL,
		token.INT_LITERAL,
		token.REAL_LITERAL,
		token.STR_LITERAL,
		token.STR_TEMPLATE,
	)
}

// matchCallStart   := ID "("
func matchCallStart(tc *TokenCollector) (_ int) {
	return matchSeq(tc, token.ID, token.OPEN_PAREN)
}

// matchParam       := VOID | matchLiteral | matchIdOrItem .
func matchParam(tc *TokenCollector) int {

	if 1 == matchAny(tc, token.VOID) {
		return 1
	}

	return matchEither(tc,
		matchLiteral,
		matchIdOrItem,
	)
}

// PARAM_LIST       := matchParam { "," matchParam } .
func matchParamList(tc *TokenCollector) (_ int) {

	n := matchParam(tc)
	if n == 0 {
		return
	}

	for 1 == matchAny(tc, token.DELIM) {
		n++

		i := matchParam(tc)
		if i == 0 {
			panic(NewParseErr("Expected another parameter", nil, tc.Peek()))
		}

		n += i
	}

	return n
}

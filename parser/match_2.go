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
	"github.com/PaulioRandall/scarlet-go/token"
)

// match attempts to pattern match the list of token kinds to the tokens read
// by the TokenCollector. If they all match the number of tokens is returned
// else zero is returned.
func match(tc *TokenCollector, ks ...token.Kind) (_ int) {

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

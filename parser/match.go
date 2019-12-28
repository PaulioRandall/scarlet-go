package parser

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
// PARAM_LIST       := [ PARAM ] { "," ( PARAM ) } .
// PARAM            := "\_" | ID_OR_ITEM | LITERAL .
// OPERATOR         := NUM_OPERATOR | BOOL_OPERATOR | CMP_OPERATOR .
// CMP_OPERATOR     := "=" | "#" | "<" | ">" | "<=" | ">=" .
// BOOL_OPERATOR    := "|" | "&" .
// NUM_OPERATOR     := "+" | "-" | "\*" | "/" | "%" .
// LIST             := "{" LIST_ITEMS [ "," [ NEWLINE ] ] "}" .
// LIST_ITEMS       := EXPR { "," [ NEWLINE ] EXPR } .

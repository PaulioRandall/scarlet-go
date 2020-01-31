# Grammer

## Production Rules (WSN)

P := {  
- COMMENT          := "//" * Any visible unicode character * NEWLINE .
- PROGRAM          := STATEMENT BLOCK .
- BLOCK            := { STATEMENT } .
- STATEMENT        := ( ASSIGNMENT | INLINE_EXPR | GUARD | MATCH_BLOCK | WATCH_BLOCK ) NEWLINE .
- INLINE_STATEMENT := ( ASSIGNMENT | INLINE_EXPR ) NEWLINE .
- EXPR             := ID_OR_ITEM | INLINE_EXPR .
- INLINE_EXPR      := LITERAL | FUNC_CALL | SPELL | OPERATION .
- SPELL            := "@" FUNC_CALL .
- FUNC_CALL        := ID "(" PARAM_LIST ")" .
- ASSIGNMENT       := [ "STICKY" ] ID_ARRAY ":=" ( LIST | EXPR | FUNC ) .
- FUNC             := "F" "(" PARAM_LIST [ "->" ID_ARRAY ] ")" BODY .
- GUARD            := "[" EXPR "]" BODY .
- OPERATION        := OPERAND OPERATOR { OPERAND OPERATOR } OPERAND .
- OPERAND          := [ "~" | "¬" ] ( ID_OR_ITEM | LITERAL | FUNC_CALL | SPELL ) .
- BODY             := INLINE_STATEMENT | ( "DO" NEWLINE BLOCK "END" ) .
- MATCH_BLOCK      := "MATCH" NEWLINE MATCH_CASE { MATCH_CASE } "END" .
- MATCH_CASE       := EXPR BODY NEWLINE .
- WATCH_BLOCK      := "WATCH" ID { "," ID } NEWLINE BLOCK "END" .
- PARAM_LIST       := [ PARAM { "," PARAM } ] .
- PARAM            := "\_" | ID_USAGE | LITERAL .
- ID_ARRAY         := ID_OR_VOID { "," ID_OR_VOID } .
- ID_OR_VOID       := ID | "\_" .
- ID               := LETTER { "\_" | LETTER } .
- OPERATOR         := NUM_OPERATOR | BOOL_OPERATOR | CMP_OPERATOR .
- CMP_OPERATOR     := "=" | "#" | "<" | ">" | "<=" | ">=" .
- BOOL_OPERATOR    := "|" | "&" .
- NUM_OPERATOR     := "+" | "-" | "\*" | "/" | "%" .
- LITERAL          := BOOL | INT | REAL | STRING | TEMPLATE .
- ID_OR_ITEM       := ID [ ITEM_ACCESS ] .
- ITEM_ACCESS      := "[" ( ID | INT ) "]" .
- LIST             := "{" LIST_ITEMS [ "," [ NEWLINE ] ] "}" .
- LIST_ITEMS       := EXPR { "," [ NEWLINE ] EXPR } .
- BOOL             := "TRUE" | "FALSE" .
- REAL             := INT [ "." INT ] .
- INT              := DIGIT { DIGIT } .
- DIGIT            := * Unicode category Nd (0-9) * .
- STRING           := "\`" * Any visible unicode character * "\`" .
- TEMPLATE         := '"' * Any control or visible unicode character * '"' .
- LETTER           := * Unicode category L (letter) * .
- NEWLINE          := * LF or CRLF * .
}

## Start

S := {  
- PROGRAM  
}

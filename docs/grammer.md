# Grammer

G := (V, T, P, S)  

## Non-terminals

V := {  
  COMMENT  
  PROGRAM  
  BLOCK  
  STATEMENT  
  INLINE_STATEMENT  
  EXPR  
  INLINE_EXPR  
  SPELL  
  FUNC_CALL  
  ASSIGNMENT  
  FUNC  
  GUARD  
  OPERATION  
  OPERAND  
  BODY  
  MATCH_BLOCK  
  MATCH_CASE  
  WATCH_BLOCK  
  PARAM_LIST  
  PARAM  
  ID_USAGE  
  ID_ARRAY  
  ID_OR_VOID  
  ID  
  OPERATOR  
  CMP_OPERATOR  
  BOOL_OPERATOR  
  NUM_OPERATOR  
  LITERAL  
  LIST_ACCESS  
  LIST  
  LIST_ITEMS  
  BOOL  
  NUMBER  
  INTEGER  
  DIGIT  
  STRING  
  TEMPLATE  
  LETTER  
  NEWLINE  
}

## Terminals

T := {  
  '@'  
  '('  
  ')'  
  ':'  
  ','  
  '"'  
  '\`'  
  '='  
  '#'  
  '<'  
  '>'  
  '+'  
  '-'  
  '\*'  
  '/'  
  '%'  
  '~'  
  '¬'  
  '|'  
  '&'  
  '\_'  
  '['  
  ']'  
  * Any control or visible unicode character *  
}

## Production Rules (WSN)

P := {  
  COMMENT          := "//" * Any visible unicode character * NEWLINE .  
  PROGRAM          := STATEMENT BLOCK .  
  BLOCK            := { STATEMENT } .  
  STATEMENT        := ( ASSIGNMENT | INLINE_EXPR | GUARD | MATCH_BLOCK | WATCH_BLOCK ) NEWLINE .  
  INLINE_STATEMENT := ( ASSIGNMENT | INLINE_EXPR ) NEWLINE .  
  EXPR             := ID_USAGE | INLINE_EXPR .  
  INLINE_EXPR      := LITERAL | FUNC_CALL | SPELL | OPERATION .  
  SPELL            := "@" FUNC_CALL .  
  FUNC_CALL        := ID "(" PARAM_LIST ")" .  
  ASSIGNMENT       := [ "GLOBAL" ] ID_ARRAY ":=" ( LIST | EXPR | FUNC ) .  
  FUNC             := "F" "(" PARAM_LIST [ "->" ID_ARRAY ] ")" BODY .  
  GUARD            := "[" EXPR "]" BODY .  
  OPERATION        := OPERAND OPERATOR { OPERAND OPERATOR } OPERAND .  
  OPERAND          := [ "~" | "¬" ] ( ID_USAGE | LITERAL | FUNC_CALL | SPELL ) .  
  BODY             := INLINE_STATEMENT | ( "DO" NEWLINE BLOCK "END" ) .  
  MATCH_BLOCK      := "MATCH" NEWLINE MATCH_CASE { MATCH_CASE } "END" .  
  MATCH_CASE       := EXPR BODY NEWLINE .  
  WATCH_BLOCK      := "WATCH" ID { "," ID } NEWLINE BLOCK "END" .  
  PARAM_LIST       := [ PARAM ] { "," ( PARAM ) } .  
  PARAM            := "\_" | ID_USAGE | LITERAL .  
  ID_USAGE         := ID [ LIST_ACCESS ] .  
  ID_ARRAY         := ID_OR_VOID { "," ID_OR_VOID } .  
  ID_OR_VOID       := ID | "\_" .  
  ID               := LETTER { "\_" | LETTER } .  
  OPERATOR         := NUM_OPERATOR | BOOL_OPERATOR | CMP_OPERATOR .  
  CMP_OPERATOR     := "=" | "#" | "<" | ">" | "<=" | ">=" .  
  BOOL_OPERATOR    := "|" | "&" .  
  NUM_OPERATOR     := "+" | "-" | "\*" | "/" | "%" .  
  LITERAL          := BOOL | NUMBER | STRING | TEMPLATE.  
  LIST_ACCESS      := "(" ( ID | INTEGER ) ")" .  
  LIST             := "{" LIST_ITEMS [ "," [ NEWLINE ] ] "}" .  
  LIST_ITEMS       := EXPR { "," [ NEWLINE ] EXPR } .  
  BOOL             := "TRUE" | "FALSE" .  
  NUMBER           := INTEGER [ "." INTEGER ] .  
  INTEGER          := DIGIT { DIGIT } .  
  DIGIT            := * Unicode category Nd (0-9) * .  
  STRING           := "\`" * Any visible unicode character * "\`" .  
  TEMPLATE         := '"' * Any control or visible unicode character * '"' .  
  LETTER           := * Unicode category L (letter) * .  
  NEWLINE          := * LF or CRLF * .  
}

## Start

S := {  
  PROGRAM  
}

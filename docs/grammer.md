# Grammer

G := (V, T, P, S)

## Non-terminals

V := {
  PROGRAM
  STATEMENT
  SPELL
  CALL
  ASSIGNMENT
  FUNC
  PARAMS
  BLOCK
  IDS
  ID
  LETTER
}

## Terminals

T := {
  '@'
  '('
  ')'
  ':'
  '='
  ','
  * Unicode category L (letter) *
}

## Production Rules (WSN)

P := {
  PROGRAM    := STATEMENT { STATEMENT } .
  STATEMENT  := ( ASSIGNMENT | CALL | SPELL ) .
  SPELL      := "@" CALL .
  CALL       := ID PARAMS .
  ASSIGNMENT := IDS ":=" FUNC .
  FUNC       := "F" PARAMS IDS BLOCK .
  PARAMS     := "(" IDS ")" .
  BLOCK      := "DO" { STATEMENT } "END" .
  IDS        := ID { "," ID } .
  ID         := LETTER { LETTER } .
  LETTER     := * Unicode category L (letter) * .
}

## Start

S := {
  PROGRAM
}

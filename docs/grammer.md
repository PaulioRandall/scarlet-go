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
  NEWLINE
}

## Terminals

T := {
  "@"
  "("
  ")"
  ":"
  "="
  ","
  * Unicode category L (letter) *
  * LF or CRLF *
}

## Production Rules (WSN)

P := {
  PROGRAM    := STATEMENT { STATEMENT } .
  STATEMENT  := ( ASSIGNMENT | CALL | SPELL ) NEWLINE .
  SPELL      := "@" CALL .
  CALL       := ID PARAMS .
  ASSIGNMENT := IDS ":=" FUNC .
  FUNC       := "FUNC" PARAMS IDS BLOCK .
  PARAMS     := "(" IDS ")" .
  BLOCK      := "DO" NEWLINE { STATEMENT } "END" NEWLINE .
  IDS        := ID { "," ID } .
  ID         := LETTER { LETTER } .
  LETTER     := * Unicode category L (letter) * .
  NEWLINE    := * LF or CRLF * .
}

## Start

S := {
  PROGRAM
}

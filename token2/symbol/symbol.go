// Symbol package defines the literal non-terminal symbols used in source code.
package symbol

// Symbol represents a non-terminal symbol.
type Symbol string

const (
	TERMINATOR    Symbol = ";"
	LF                   = "\n"
	CRLF                 = "\r\n"
	TRUE                 = "true"
	FALSE                = "false"
	LOOP                 = "loop"
	STRING_PREFIX        = `"`
	STRING_SUFFIX        = `"`
	NUMBER_DELIM         = "."
	SPELL_PREFIX         = "@"
	SPELL_DELIM          = "."
	DELIM                = ","
	L_PAREN              = "("
	R_PAREN              = ")"
	L_SQUARE             = "["
	R_SQUARE             = "]"
	L_CURLY              = "{"
	R_CURLY              = "}"
	ASSIGN               = ":="
	VOID                 = "_"
	ADD                  = "+"
	SUB                  = "-"
	MUL                  = "*"
	DIV                  = "/"
	REM                  = "%"
	AND                  = "&&"
	OR                   = "||"
	LESS                 = "<"
	MORE                 = ">"
	LESS_EQUAL           = "<="
	MORE_EQUAL           = ">="
	EQUAL                = "=="
	NOT_EQUAL            = "!="
)

// Symbol package defines the terminal and non-terminal symbols used in
// scanning source code.
package symbol

// TerminalSymbol are used almost exclusively in scanning.
type TerminalSymbol string

const (
	UNDEFINED_TERMINAL TerminalSymbol = ""
	CR                                = '\r'
	LF                                = '\n'
	STRING_PREFIX                     = '"'
	STRING_SUFFIX                     = '"'
	STRING_ESCAPE                     = '\\'
	NUMBER_FRAC_DELIM                 = '.'
	SPELL_NAME_DELIM                  = '.'
	VOID                              = '_'
)

// NonTerminalSymbol are used almost exclusively in scanning.
type NonTerminalSymbol string

const (
	UNDEFINED_NONTERMINAL NonTerminalSymbol = ""
	COMMENT_PREFIX                          = "#"
	TERMINATOR                              = ";"
	CRLF                                    = "\r\n"
	TRUE                                    = "true"
	FALSE                                   = "false"
	LOOP                                    = "loop"
	SPELL_PREFIX                            = "@"
	DELIM                                   = ","
	L_PAREN                                 = "("
	R_PAREN                                 = ")"
	L_SQUARE                                = "["
	R_SQUARE                                = "]"
	L_CURLY                                 = "{"
	R_CURLY                                 = "}"
	ASSIGN                                  = ":="
	ADD                                     = "+"
	SUB                                     = "-"
	MUL                                     = "*"
	DIV                                     = "/"
	REM                                     = "%"
	AND                                     = "&&"
	OR                                      = "||"
	LESS                                    = "<"
	MORE                                    = ">"
	LESS_EQUAL                              = "<="
	MORE_EQUAL                              = ">="
	EQUAL                                   = "=="
	NOT_EQUAL                               = "!="
)

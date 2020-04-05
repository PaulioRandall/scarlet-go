package token

// Kind represents a token type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	SOF         Kind = `SOF`
	EOF         Kind = `EOF`
	COMMENT     Kind = `COMMENT`
	WHITESPACE  Kind = `WHITESPACE`
	NEWLINE     Kind = `NEWLINE`
	FUNC        Kind = `FUNC`
	DO          Kind = `DO`
	INLINE      Kind = `INLINE`
	MATCH       Kind = `MATCH`
	END         Kind = `END`
	ID          Kind = `ID`
	DELIM       Kind = `DELIM`
	ASSIGN      Kind = `ASSIGN`
	RETURNS     Kind = `RETURNS`
	OPEN_PAREN  Kind = `OPEN_PAREN`
	CLOSE_PAREN Kind = `CLOSE_PAREN`
	OPEN_GUARD  Kind = `OPEN_GUARD`
	CLOSE_GUARD Kind = `CLOSE_GUARD`
	OPEN_LIST   Kind = `OPEN_LIST`
	CLOSE_LIST  Kind = `CLOSE_LIST`
	SPELL       Kind = `SPELL`
	STR         Kind = `STR`
	TEMPLATE    Kind = `TEMPLATE`
	INT         Kind = `INT`
	REAL        Kind = `REAL`
	BOOL        Kind = `BOOL`
	NOT         Kind = `NOT`
	ADD         Kind = `ADD`
	SUBTRACT    Kind = `SUBTRACT`
	MULTIPLY    Kind = `MULTIPLY`
	DIVIDE      Kind = `DIVIDE`
	MOD         Kind = `MOD`
	AND         Kind = `AND`
	OR          Kind = `OR`
	EQU         Kind = `EQUAL`
	NEQ         Kind = `NOT_EQUAL`
	LT          Kind = `LESS_THAN`
	LT_OR_EQU   Kind = `LESS_THAN_OR_EQUAL`
	MT          Kind = `MORE_THAN`
	MT_OR_EQU   Kind = `MORE_THAN_OR_EQUAL`
	VOID        Kind = `VOID`
	TERMINATOR  Kind = `TERMINATOR`
)

// Collection of terminal and non-terminal symbols for ease of reference and to
// speed up syntax experimentation.
const (
	NON_TERMINAL_FUNCTION              string = `F`
	NON_TERMINAL_NORMAL_BLOCK_START    string = `DO`
	NON_TERMINAL_MATCH_BLOCK_START     string = `MATCH`
	NON_TERMINAL_BLOCK_END             string = `END`
	NON_TERMINAL_TRUE                  string = `TRUE`
	NON_TERMINAL_FALSE                 string = `FALSE`
	TERMINAL_CARRIAGE_RETURN           rune   = '\r'
	TERMINAL_LINEFEED                  rune   = '\n'
	TERMINAL_COMMENT_START             rune   = '/'
	TERMINAL_FRACTIONAL_DELIM          rune   = '.'
	TERMINAL_STRING_START              rune   = '`'
	TERMINAL_STRING_END                rune   = '`'
	TERMINAL_TEMPLATE_START            rune   = '"'
	TERMINAL_TEMPLATE_END              rune   = '"'
	TERMINAL_TEMPLATE_ESCAPE           rune   = '\\'
	TERMINAL_WORD_UNDERSCORE           rune   = '_'
	NON_TERMINAL_ASSIGNMENT            string = `:=`
	NON_TERMINAL_RETURN_PARAMS         string = `->`
	TERMINAL_OPEN_PAREN                rune   = '('
	TERMINAL_CLOSE_PAREN               rune   = ')'
	TERMINAL_OPEN_GUARD                rune   = '['
	TERMINAL_CLOSE_GUARD               rune   = ']'
	TERMINAL_OPEN_LIST                 rune   = '{'
	TERMINAL_CLOSE_LIST                rune   = '}'
	TERMINAL_LIST_DELIM                rune   = ','
	TERMINAL_VOID_VALUE                rune   = '_'
	TERMINAL_STATEMENT_TERMINATOR      rune   = ';'
	TERMINAL_SPELL_PREFIX              rune   = '@'
	TERMINAL_UNIVERSAL_NEGATION        rune   = '~'
	TERMINAL_TEA_DRINKING_NEGATION     rune   = '¬'
	TERMINAL_MATH_ADDITION             rune   = '+'
	TERMINAL_MATH_SUBTRACTION          rune   = '-'
	TERMINAL_MATH_MULTIPLICATION       rune   = '*'
	TERMINAL_MATH_DIVISION             rune   = '/'
	TERMINAL_MATH_REMAINDER            rune   = '%'
	TERMINAL_LOGICAL_AND               rune   = '&'
	TERMINAL_LOGICAL_OR                rune   = '|'
	TERMINAL_EQUALITY                  rune   = '='
	TERMINAL_UNEQUALITY                rune   = '#'
	NON_TERMINAL_LESS_THAN_OR_EQUAL    string = "<="
	NON_TERMINAL_GREATER_THAN_OR_EQUAL string = "=>"
	TERMINAL_LESS_THAN                 rune   = '<'
	TERMINAL_MORE_THAN                 rune   = '>'
)

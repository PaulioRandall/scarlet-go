package token

// Kind represents a token type.
type Kind string

const (
	KIND_UNDEFINED Kind = ``
	// ------------------
	KIND_SOF         Kind = `SOF`
	KIND_EOF         Kind = `EOF`
	KIND_COMMENT     Kind = `COMMENT`
	KIND_WHITESPACE  Kind = `WHITESPACE`
	KIND_NEWLINE     Kind = `NEWLINE`
	KIND_FUNC        Kind = `FUNCTION_KEYWORD`
	KIND_MATCH       Kind = `MATCH_BLOCK_KEYWORD`
	KIND_INLINE      Kind = `INLINE_BLOCK_IMPLIED`
	KIND_ID          Kind = `ID`
	KIND_DELIM       Kind = `DELIMITER`
	KIND_ASSIGN      Kind = `ASSIGNMENT`
	KIND_RETURNS     Kind = `RETURNS`
	KIND_DO          Kind = `BLOCK_OPEN`
	KIND_END         Kind = `BLOCK_CLOSE`
	KIND_OPEN_PAREN  Kind = `PAREN_OPEN`
	KIND_CLOSE_PAREN Kind = `PAREN_CLOSE`
	KIND_OPEN_GUARD  Kind = `GUARD_OPEN`
	KIND_CLOSE_GUARD Kind = `GUARD_CLOSE`
	KIND_OPEN_LIST   Kind = `LIST_OPEN`
	KIND_CLOSE_LIST  Kind = `LIST_CLOSE`
	// TODO: kind these
	SPELL      Kind = `SPELL`
	STR        Kind = `STR`
	TEMPLATE   Kind = `TEMPLATE`
	INT        Kind = `INT`
	REAL       Kind = `REAL`
	BOOL       Kind = `BOOL`
	NOT        Kind = `NOT`
	ADD        Kind = `ADD`
	SUBTRACT   Kind = `SUBTRACT`
	MULTIPLY   Kind = `MULTIPLY`
	DIVIDE     Kind = `DIVIDE`
	MOD        Kind = `MOD`
	AND        Kind = `AND`
	OR         Kind = `OR`
	EQU        Kind = `EQUAL`
	NEQ        Kind = `NOT_EQUAL`
	LT         Kind = `LESS_THAN`
	LT_OR_EQU  Kind = `LESS_THAN_OR_EQUAL`
	MT         Kind = `MORE_THAN`
	MT_OR_EQU  Kind = `MORE_THAN_OR_EQUAL`
	VOID       Kind = `VOID`
	TERMINATOR Kind = `TERMINATOR`
)

// Collection of terminal and non-terminal symbols for ease of reference and to
// speed up syntax experimentation.
const (
	LEXEME_FUNCTION         string = `F`
	LEXEME_BLOCK_START      string = `DO`
	LEXEME_BLOCK_END        string = `END`
	LEXEME_TRUE             string = `TRUE`
	LEXEME_FALSE            string = `FALSE`
	LEXEME_NEWLINE_LF       string = "\n"
	LEXEME_NEWLINE_CRLF     string = "\r\n"
	LEXEME_COMMENT_START    string = "/"
	LEXEME_FRACTIONAL_DELIM string = "."
	// TODO: lexeme theses
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

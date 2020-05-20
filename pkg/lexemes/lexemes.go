package lexemes

type Lexeme int

// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.

const (
	UNDEFINED = iota
	// ------------------
	ANY
	ANOTHER
	// ------------------
	EOF
	// ------------------
	COMMENT
	WHITESPACE
	NEWLINE
	FUNC
	FIX
	ID
	DELIM
	ASSIGN
	OUTPUT
	BLOCK_OPEN
	BLOCK_CLOSE
	PAREN_OPEN
	PAREN_CLOSE
	LIST
	MATCH
	GUARD_OPEN
	GUARD_CLOSE
	LOOP
	SPELL
	STRING
	TEMPLATE
	NUMBER
	BOOL
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
	REMAINDER
	AND
	OR
	EQUAL
	NOT_EQUAL
	LESS_THAN
	LESS_THAN_OR_EQUAL
	MORE_THAN
	MORE_THAN_OR_EQUAL
	VOID
	TERMINATOR
	LIST_START
	LIST_END
)

var values map[Lexeme]string = map[Lexeme]string{
	UNDEFINED:          ``,
	ANOTHER:            `ANOTHER`,
	EOF:                `EOF`,
	COMMENT:            `COMMENT`,
	WHITESPACE:         `WHITESPACE`,
	NEWLINE:            `NEWLINE`,
	FUNC:               `FUNC`,
	FIX:                `FIX`,
	ID:                 `ID`,
	DELIM:              `DELIM`,
	ASSIGN:             `ASSIGN`,
	OUTPUT:             `OUTPUT`,
	BLOCK_OPEN:         `BLOCK_OPEN`,
	BLOCK_CLOSE:        `BLOCK_CLOSE`,
	PAREN_OPEN:         `PAREN_OPEN`,
	PAREN_CLOSE:        `PAREN_CLOSE`,
	LIST:               `LIST`,
	MATCH:              `MATCH`,
	GUARD_OPEN:         `GUARD_OPEN`,
	GUARD_CLOSE:        `GUARD_CLOSE`,
	LOOP:               `LOOP`,
	SPELL:              `SPELL`,
	STRING:             `STRING`,
	TEMPLATE:           `TEMPLATE`,
	NUMBER:             `NUMBER`,
	BOOL:               `BOOL`,
	ADD:                `ADD`,
	SUBTRACT:           `SUBTRACT`,
	MULTIPLY:           `MULTIPLY`,
	DIVIDE:             `DIVIDE`,
	REMAINDER:          `REMAINDER`,
	AND:                `AND`,
	OR:                 `OR`,
	EQUAL:              `EQUAL`,
	NOT_EQUAL:          `NOT_EQUAL`,
	LESS_THAN:          `LESS_THAN`,
	LESS_THAN_OR_EQUAL: `LESS_THAN_OR_EQUAL`,
	MORE_THAN:          `MORE_THAN`,
	MORE_THAN_OR_EQUAL: `MORE_THAN_OR_EQUAL`,
	VOID:               `VOID`,
	TERMINATOR:         `TERMINATOR`,
	LIST_START:         `LIST_START`,
	LIST_END:           `LIST_END`,
}

func String(i Lexeme) string {
	return values[i]
}

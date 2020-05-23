package z_token

type Morpheme int

// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.

const (
	M_UNDEFINED Morpheme = iota
	// ------------------
	M_ANY
	M_ANOTHER
	// ------------------
	M_EOF
	// ------------------
	M_COMMENT            // K_REDUNDANT
	M_WHITESPACE         // K_REDUNDANT
	M_NEWLINE            // K_NEWLINE
	M_FUNC               // K_KEYWORD
	M_FIX                // K_KEYWORD
	M_LIST               // K_KEYWORD
	M_MATCH              // K_KEYWORD
	M_LOOP               // K_KEYWORD
	M_SPELL              // K_KEYWORD
	M_IDENTIFIER         // K_IDENTIFIER
	M_VOID               // K_IDENTIFIER
	M_DELIM              // K_DELIMITER
	M_ASSIGN             // K_DELIMITER
	M_OUTPUT             // K_DELIMITER
	M_BLOCK_OPEN         // K_DELIMITER
	M_BLOCK_CLOSE        // K_DELIMITER
	M_PAREN_OPEN         // K_DELIMITER
	M_PAREN_CLOSE        // K_DELIMITER
	M_GUARD_OPEN         // K_DELIMITER
	M_GUARD_CLOSE        // K_DELIMITER
	M_TERMINATOR         // K_DELIMITER
	M_STRING             // K_LITERAL
	M_TEMPLATE           // K_LITERAL
	M_NUMBER             // K_LITERAL
	M_BOOL               // K_LITERAL
	M_ADD                // K_ARITHMETIC
	M_SUBTRACT           // K_ARITHMETIC
	M_MULTIPLY           // K_ARITHMETIC
	M_DIVIDE             // K_ARITHMETIC
	M_REMAINDER          // K_ARITHMETIC
	M_AND                // K_LOGIC
	M_OR                 // K_LOGIC
	M_EQUAL              // K_COMPARISON
	M_NOT_EQUAL          // K_COMPARISON
	M_LESS_THAN          // K_COMPARISON
	M_LESS_THAN_OR_EQUAL // K_COMPARISON
	M_MORE_THAN          // K_COMPARISON
	M_MORE_THAN_OR_EQUAL // K_COMPARISON
	M_LIST_START         // K_REFERENCE
	M_LIST_END           // K_REFERENCE
)

var morphemes map[Morpheme]string = map[Morpheme]string{
	M_UNDEFINED:          ``,
	M_ANOTHER:            `ANOTHER`,
	M_EOF:                `EOF`,
	M_COMMENT:            `COMMENT`,
	M_WHITESPACE:         `WHITESPACE`,
	M_NEWLINE:            `NEWLINE`,
	M_FUNC:               `FUNC`,
	M_FIX:                `FIX`,
	M_LIST:               `LIST`,
	M_MATCH:              `MATCH`,
	M_LOOP:               `LOOP`,
	M_SPELL:              `SPELL`,
	M_IDENTIFIER:         `ID`,
	M_VOID:               `VOID`,
	M_DELIM:              `DELIM`,
	M_ASSIGN:             `ASSIGN`,
	M_OUTPUT:             `OUTPUT`,
	M_BLOCK_OPEN:         `BLOCK_OPEN`,
	M_BLOCK_CLOSE:        `BLOCK_CLOSE`,
	M_PAREN_OPEN:         `PAREN_OPEN`,
	M_PAREN_CLOSE:        `PAREN_CLOSE`,
	M_GUARD_OPEN:         `GUARD_OPEN`,
	M_GUARD_CLOSE:        `GUARD_CLOSE`,
	M_TERMINATOR:         `TERMINATOR`,
	M_STRING:             `STRING`,
	M_TEMPLATE:           `TEMPLATE`,
	M_NUMBER:             `NUMBER`,
	M_BOOL:               `BOOL`,
	M_ADD:                `ADD`,
	M_SUBTRACT:           `SUBTRACT`,
	M_MULTIPLY:           `MULTIPLY`,
	M_DIVIDE:             `DIVIDE`,
	M_REMAINDER:          `REMAINDER`,
	M_AND:                `AND`,
	M_OR:                 `OR`,
	M_EQUAL:              `EQUAL`,
	M_NOT_EQUAL:          `NOT_EQUAL`,
	M_LESS_THAN:          `LESS_THAN`,
	M_LESS_THAN_OR_EQUAL: `LESS_THAN_OR_EQUAL`,
	M_MORE_THAN:          `MORE_THAN`,
	M_MORE_THAN_OR_EQUAL: `MORE_THAN_OR_EQUAL`,
	M_LIST_START:         `LIST_START`,
	M_LIST_END:           `LIST_END`,
}

func (m Morpheme) String() string {
	return morphemes[m]
}

func (m Morpheme) Precedence() int {
	switch m {
	case M_MULTIPLY, M_DIVIDE, M_REMAINDER:
		return 6 // Multiplicative

	case M_ADD, M_SUBTRACT:
		return 5 // Additive

	case M_LESS_THAN, M_LESS_THAN_OR_EQUAL, M_MORE_THAN, M_MORE_THAN_OR_EQUAL:
		return 4 // Relational

	case M_EQUAL, M_NOT_EQUAL:
		return 3 // Equalitive

	case M_AND:
		return 2

	case M_OR:
		return 1
	}

	return 0
}

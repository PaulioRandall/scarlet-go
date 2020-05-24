package z_token

type Morpheme int

// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.

const (
	UNDEFINED Morpheme = iota
	// ------------------
	ANY
	ANOTHER
	// ------------------
	COMMENT            // K_REDUNDANT
	WHITESPACE         // K_REDUNDANT
	NEWLINE            // K_NEWLINE
	FUNC               // K_KEYWORD
	FIX                // K_KEYWORD
	LIST               // K_KEYWORD
	MATCH              // K_KEYWORD
	LOOP               // K_KEYWORD
	SPELL              // K_KEYWORD
	OUTPUT             // K_KEYWORD
	IDENTIFIER         // K_IDENTIFIER
	VOID               // K_IDENTIFIER
	DELIMITER          // K_DELIMITER
	ASSIGN             // K_DELIMITER
	BLOCK_OPEN         // K_DELIMITER
	BLOCK_CLOSE        // K_DELIMITER
	PAREN_OPEN         // K_DELIMITER
	PAREN_CLOSE        // K_DELIMITER
	GUARD_OPEN         // K_DELIMITER
	GUARD_CLOSE        // K_DELIMITER
	TERMINATOR         // K_DELIMITER
	STRING             // K_LITERAL
	TEMPLATE           // K_LITERAL
	NUMBER             // K_LITERAL
	BOOL               // K_LITERAL
	ADD                // K_ARITHMETIC
	SUBTRACT           // K_ARITHMETIC
	MULTIPLY           // K_ARITHMETIC
	DIVIDE             // K_ARITHMETIC
	REMAINDER          // K_ARITHMETIC
	AND                // K_LOGIC
	OR                 // K_LOGIC
	EQUAL              // K_COMPARISON
	NOT_EQUAL          // K_COMPARISON
	LESS_THAN          // K_COMPARISON
	LESS_THAN_OR_EQUAL // K_COMPARISON
	MORE_THAN          // K_COMPARISON
	MORE_THAN_OR_EQUAL // K_COMPARISON
	LIST_START         // K_REFERENCE
	LIST_END           // K_REFERENCE
)

var morphemes map[Morpheme]string = map[Morpheme]string{
	UNDEFINED:          ``,
	ANOTHER:            `ANOTHER`,
	COMMENT:            `COMMENT`,
	WHITESPACE:         `WHITESPACE`,
	NEWLINE:            `NEWLINE`,
	FUNC:               `FUNC`,
	FIX:                `FIX`,
	LIST:               `LIST`,
	MATCH:              `MATCH`,
	LOOP:               `LOOP`,
	SPELL:              `SPELL`,
	OUTPUT:             `OUTPUT`,
	IDENTIFIER:         `ID`,
	VOID:               `VOID`,
	DELIMITER:          `DELIM`,
	ASSIGN:             `ASSIGN`,
	BLOCK_OPEN:         `BLOCK_OPEN`,
	BLOCK_CLOSE:        `BLOCK_CLOSE`,
	PAREN_OPEN:         `PAREN_OPEN`,
	PAREN_CLOSE:        `PAREN_CLOSE`,
	GUARD_OPEN:         `GUARD_OPEN`,
	GUARD_CLOSE:        `GUARD_CLOSE`,
	TERMINATOR:         `TERMINATOR`,
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
	LIST_START:         `LIST_START`,
	LIST_END:           `LIST_END`,
}

func (m Morpheme) String() string {
	return morphemes[m]
}

func (m Morpheme) Precedence() int {
	switch m {
	case MULTIPLY, DIVIDE, REMAINDER:
		return 6 // Multiplicative

	case ADD, SUBTRACT:
		return 5 // Additive

	case LESS_THAN, LESS_THAN_OR_EQUAL, MORE_THAN, MORE_THAN_OR_EQUAL:
		return 4 // Relational

	case EQUAL, NOT_EQUAL:
		return 3 // Equalitive

	case AND:
		return 2

	case OR:
		return 1
	}

	return 0
}

func (m Morpheme) Redundant() bool {
	return m == UNDEFINED || m == WHITESPACE || m == COMMENT
}

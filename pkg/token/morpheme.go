package token

import (
	"strings"
)

type Morpheme int

// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.

const (
	UNDEFINED Morpheme = iota
	// ------------------
	ANY
	ANOTHER
	// ------------------
	COMMENT
	WHITESPACE
	NEWLINE
	FUNC
	EXPR_FUNC
	DEF
	LIST
	WHEN
	WATCH
	LOOP
	SPELL
	OUTPUT
	IDENTIFIER
	VOID
	DELIMITER
	ASSIGN
	UPDATES
	THEN
	BLOCK_OPEN
	BLOCK_CLOSE
	PAREN_OPEN
	PAREN_CLOSE
	GUARD_OPEN
	GUARD_CLOSE
	TERMINATOR
	STRING
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
	LIST_START
	LIST_END
)

var morphemes map[Morpheme]string = map[Morpheme]string{
	UNDEFINED:          ``,
	ANOTHER:            `ANOTHER`,
	COMMENT:            `COMMENT`,
	WHITESPACE:         `WHITESPACE`,
	NEWLINE:            `NEWLINE`,
	FUNC:               `FUNCTION`,
	EXPR_FUNC:          `EXPR_FUNC`,
	DEF:                `DEFINE`,
	LIST:               `LIST`,
	WHEN:               `WHEN`,
	WATCH:              `WATCH`,
	LOOP:               `LOOP`,
	SPELL:              `SPELL`,
	OUTPUT:             `OUTPUT`,
	IDENTIFIER:         `ID`,
	VOID:               `VOID`,
	DELIMITER:          `DELIM`,
	ASSIGN:             `ASSIGN`,
	UPDATES:            `UPDATES`,
	THEN:               `THEN`,
	BLOCK_OPEN:         `BLOCK_OPEN`,
	BLOCK_CLOSE:        `BLOCK_CLOSE`,
	PAREN_OPEN:         `PAREN_OPEN`,
	PAREN_CLOSE:        `PAREN_CLOSE`,
	GUARD_OPEN:         `GUARD_OPEN`,
	GUARD_CLOSE:        `GUARD_CLOSE`,
	TERMINATOR:         `TERMINATOR`,
	STRING:             `STRING`,
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

func OperatorTypes() []Morpheme {
	return []Morpheme{
		MULTIPLY,
		DIVIDE,
		REMAINDER,
		ADD,
		SUBTRACT,
		LESS_THAN,
		LESS_THAN_OR_EQUAL,
		MORE_THAN,
		MORE_THAN_OR_EQUAL,
		EQUAL,
		NOT_EQUAL,
		AND,
		OR,
	}
}

func (m Morpheme) Redundant() bool {
	return m == UNDEFINED || m == WHITESPACE || m == COMMENT
}

func JoinMorphemes(ms ...Morpheme) string {

	sb := strings.Builder{}

	for i, m := range ms {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(m.String())
	}

	return sb.String()
}

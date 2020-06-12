package token

import (
	"strings"
)

type TokenType int

// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.

const (
	TK_UNDEFINED TokenType = iota
	// ------------------
	TK_ANY
	TK_ANOTHER
	// ------------------
	TK_COMMENT
	TK_WHITESPACE
	TK_EXIT
	TK_NEWLINE
	TK_FUNCTION
	TK_EXPR_FUNC
	TK_DEFINITION
	TK_LIST
	TK_WHEN
	TK_WATCH
	TK_LOOP
	TK_SPELL
	TK_OUTPUT
	TK_OUTPUTS
	TK_IDENTIFIER
	TK_VOID
	TK_EXISTS
	TK_DELIMITER
	TK_ASSIGNMENT
	TK_UPDATES
	TK_THEN
	TK_BLOCK_OPEN
	TK_BLOCK_CLOSE
	TK_PAREN_OPEN
	TK_PAREN_CLOSE
	TK_GUARD_OPEN
	TK_GUARD_CLOSE
	TK_TERMINATOR
	TK_STRING
	TK_NUMBER
	TK_BOOL
	TK_PLUS
	TK_MINUS
	TK_MULTIPLY
	TK_DIVIDE
	TK_REMAINDER
	TK_AND
	TK_OR
	TK_EQUAL
	TK_NOT_EQUAL
	TK_LESS_THAN
	TK_LESS_THAN_OR_EQUAL
	TK_MORE_THAN
	TK_MORE_THAN_OR_EQUAL
	TK_LIST_START
	TK_LIST_END
)

var types map[TokenType]string = map[TokenType]string{
	TK_UNDEFINED:          ``,
	TK_ANOTHER:            `ANOTHER`,
	TK_COMMENT:            `COMMENT`,
	TK_WHITESPACE:         `WHITESPACE`,
	TK_EXIT:               `EXIT`,
	TK_NEWLINE:            `NEWLINE`,
	TK_FUNCTION:           `FUNCTION`,
	TK_EXPR_FUNC:          `EXPR_FUNCTION`,
	TK_DEFINITION:         `DEFINE`,
	TK_LIST:               `LIST`,
	TK_WHEN:               `WHEN`,
	TK_WATCH:              `WATCH`,
	TK_LOOP:               `LOOP`,
	TK_SPELL:              `SPELL`,
	TK_OUTPUT:             `OUTPUT`,
	TK_OUTPUTS:            `OUTPUTS`,
	TK_IDENTIFIER:         `ID`,
	TK_VOID:               `VOID`,
	TK_EXISTS:             `EXISTS`,
	TK_DELIMITER:          `DELIM`,
	TK_ASSIGNMENT:         `ASSIGNMENT`,
	TK_UPDATES:            `UPDATES`,
	TK_THEN:               `THEN`,
	TK_BLOCK_OPEN:         `BLOCK_OPEN`,
	TK_BLOCK_CLOSE:        `BLOCK_CLOSE`,
	TK_PAREN_OPEN:         `PAREN_OPEN`,
	TK_PAREN_CLOSE:        `PAREN_CLOSE`,
	TK_GUARD_OPEN:         `GUARD_OPEN`,
	TK_GUARD_CLOSE:        `GUARD_CLOSE`,
	TK_TERMINATOR:         `TERMINATOR`,
	TK_STRING:             `STRING`,
	TK_NUMBER:             `NUMBER`,
	TK_BOOL:               `BOOL`,
	TK_PLUS:               `PLUS`,
	TK_MINUS:              `MINUS`,
	TK_MULTIPLY:           `MULTIPLY`,
	TK_DIVIDE:             `DIVIDE`,
	TK_REMAINDER:          `REMAINDER`,
	TK_AND:                `AND`,
	TK_OR:                 `OR`,
	TK_EQUAL:              `EQUAL`,
	TK_NOT_EQUAL:          `NOT_EQUAL`,
	TK_LESS_THAN:          `LESS_THAN`,
	TK_LESS_THAN_OR_EQUAL: `LESS_THAN_OR_EQUAL`,
	TK_MORE_THAN:          `MORE_THAN`,
	TK_MORE_THAN_OR_EQUAL: `MORE_THAN_OR_EQUAL`,
	TK_LIST_START:         `LIST_START`,
	TK_LIST_END:           `LIST_END`,
}

func (ty TokenType) String() string {
	return types[ty]
}

func (ty TokenType) Precedence() int {
	switch ty {
	case TK_MULTIPLY, TK_DIVIDE, TK_REMAINDER:
		return 6 // Multiplicative

	case TK_PLUS, TK_MINUS:
		return 5 // Additive

	case TK_LESS_THAN, TK_LESS_THAN_OR_EQUAL, TK_MORE_THAN, TK_MORE_THAN_OR_EQUAL:
		return 4 // Relational

	case TK_EQUAL, TK_NOT_EQUAL:
		return 3 // Equalitive

	case TK_AND:
		return 2

	case TK_OR:
		return 1
	}

	return 0
}

func OperatorTypes() []TokenType {
	return []TokenType{
		TK_MULTIPLY,
		TK_DIVIDE,
		TK_REMAINDER,
		TK_PLUS,
		TK_MINUS,
		TK_LESS_THAN,
		TK_LESS_THAN_OR_EQUAL,
		TK_MORE_THAN,
		TK_MORE_THAN_OR_EQUAL,
		TK_EQUAL,
		TK_NOT_EQUAL,
		TK_AND,
		TK_OR,
	}
}

func (ty TokenType) Redundant() bool {
	return ty == TK_UNDEFINED || ty == TK_WHITESPACE || ty == TK_COMMENT
}

func JoinTypes(tys ...TokenType) string {

	sb := strings.Builder{}

	for i, ty := range tys {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(ty.String())
	}

	return sb.String()
}

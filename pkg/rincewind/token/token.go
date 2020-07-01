package token

import (
	"fmt"
)

type Token interface {
	GenType() GenType
	SubType() SubType
	Raw() string
	Value() string
	Snippet
	fmt.Stringer
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type GenType int // General type
type SubType int

func (ge GenType) String() string {
	return genTypes[ge]
}

func (su SubType) String() string {
	return subTypes[su]
}

const (
	GE_UNDEFINED GenType = 0
	SU_UNDEFINED SubType = 0
	// ------------------
	GE_WHITESPACE GenType = iota
	// ------------------
	GE_TERMINATOR GenType = iota
	SU_TERMINATOR SubType = iota
	SU_NEWLINE
	// ------------------
	GE_IDENTIFIER GenType = iota
	SU_IDENTIFIER SubType = iota
	SU_VOID
	// ------------------
	GE_LITERAL GenType = iota
	SU_BOOL    SubType = iota
	SU_NUMBER
	SU_STRING
	// ------------------
	GE_SPELL GenType = iota
	// ------------------
	GE_DELIMITER   GenType = iota
	SU_VALUE_DELIM SubType = iota
	// ------------------
	GE_BRACKET    GenType = iota
	SU_PAREN_OPEN SubType = iota
	SU_PAREN_CLOSE
)

var genTypes = map[GenType]string{
	GE_WHITESPACE: `whitespace`,
	GE_TERMINATOR: `terminator`,
	GE_IDENTIFIER: `identifier`,
	GE_LITERAL:    `literal`,
	GE_SPELL:      `spell`,
	GE_BRACKET:    `bracket`,
}

var subTypes = map[SubType]string{
	SU_IDENTIFIER:  `identifier`,
	SU_VOID:        `void`,
	SU_TERMINATOR:  `terminator`,
	SU_NEWLINE:     `newline`,
	SU_BOOL:        `bool`,
	SU_NUMBER:      `number`,
	SU_STRING:      `string`,
	SU_PAREN_OPEN:  `paren_open`,
	SU_PAREN_CLOSE: `paren_close`,
}

/*
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
*/

package token

import (
	"fmt"
)

type Token interface {
	GenType() GenType
	SubType() SubType
	Raw() string
	Val() interface{}
	Snippet
	fmt.Stringer
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type GenType int // General type
type SubType int

func (gt GenType) String() string {
	return genTypes[gt]
}

func (st SubType) String() string {
	return subTypes[st]
}

const (
	GT_UNDEFINED GenType = 0
	ST_UNDEFINED SubType = 0
	// ------------------
	GT_WHITESPACE GenType = iota
	// ------------------
	GT_TERMINATOR GenType = iota
	ST_TERMINATOR SubType = iota
	ST_NEWLINE
	// ------------------
	GT_IDENTIFIER GenType = iota
	ST_IDENTIFIER SubType = iota
	ST_VOID
	// ------------------
	GT_LITERAL GenType = iota
	ST_BOOL    SubType = iota
	ST_NUMBER
	ST_STRING
	// ------------------
	GT_SPELL GenType = iota
	// ------------------
	GT_DELIMITER   GenType = iota
	ST_VALUE_DELIM SubType = iota
	// ------------------
	GT_BRACKET    GenType = iota
	ST_PAREN_OPEN SubType = iota
	ST_PAREN_CLOSE
)

var genTypes = map[GenType]string{
	GT_WHITESPACE: `whitespace`,
	GT_TERMINATOR: `terminator`,
	GT_IDENTIFIER: `identifier`,
	GT_LITERAL:    `literal`,
	GT_SPELL:      `spell`,
	GT_BRACKET:    `bracket`,
}

var subTypes = map[SubType]string{
	ST_IDENTIFIER:  `identifier`,
	ST_VOID:        `void`,
	ST_TERMINATOR:  `terminator`,
	ST_NEWLINE:     `newline`,
	ST_BOOL:        `bool`,
	ST_NUMBER:      `number`,
	ST_STRING:      `string`,
	ST_PAREN_OPEN:  `paren_open`,
	ST_PAREN_CLOSE: `paren_close`,
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

package statement

import (
//"strings"
)

type Kind int

const (
	ST_UNDEFINED Kind = iota
	ST_VOID
	ST_IDENTIFIER
	ST_LITERAL
	ST_LIST_ACCESSOR
	ST_LIST_CONSTRUCTOR
	ST_NEGATION
	ST_OPERATION
	ST_ASSIGNMENT
	ST_BLOCK
	ST_PARAMETERS
	ST_FUNCTION
)

var kinds map[Kind]string = map[Kind]string{
	ST_UNDEFINED:        ``,
	ST_VOID:             `Void`,
	ST_IDENTIFIER:       `Identifier`,
	ST_LITERAL:          `Literal`,
	ST_LIST_ACCESSOR:    `List-Accessor`,
	ST_LIST_CONSTRUCTOR: `List-Constructor`,
	ST_NEGATION:         `Negation`,
	ST_OPERATION:        `Operation`,
	ST_ASSIGNMENT:       `Assignment`,
	ST_BLOCK:            `Block`,
	ST_PARAMETERS:       `Parameters`,
	ST_FUNCTION:         `Function`,
}

func (k Kind) String() string {
	return kinds[k]
}

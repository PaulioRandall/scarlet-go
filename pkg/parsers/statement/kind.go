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
	ST_NEGATION
	ST_OPERATION
	ST_ASSIGNMENT
	ST_ASSIGNMENT_BLOCK
	ST_BLOCK
	ST_EXPRESSION_FUNCTION
	ST_PARAMETERS
	ST_FUNCTION
	ST_WATCH
	ST_GUARD
	ST_MATCH_CASE
	ST_MATCH
	ST_LOOP
)

var kinds map[Kind]string = map[Kind]string{
	ST_UNDEFINED:           ``,
	ST_VOID:                `Void`,
	ST_IDENTIFIER:          `Identifier`,
	ST_LITERAL:             `Literal`,
	ST_LIST_ACCESSOR:       `List-Accessor`,
	ST_NEGATION:            `Negation`,
	ST_OPERATION:           `Operation`,
	ST_ASSIGNMENT:          `Assignment`,
	ST_ASSIGNMENT_BLOCK:    `Assignment-Block`,
	ST_BLOCK:               `Block`,
	ST_EXPRESSION_FUNCTION: `Expression-Function`,
	ST_PARAMETERS:          `Parameters`,
	ST_FUNCTION:            `Function`,
	ST_WATCH:               `Watch`,
	ST_GUARD:               `Guard`,
	ST_MATCH_CASE:          `MatchCase`,
	ST_MATCH:               `Match`,
	ST_LOOP:                `Loop`,
}

func (k Kind) String() string {
	return kinds[k]
}
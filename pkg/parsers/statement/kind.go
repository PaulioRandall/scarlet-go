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
)

var kinds map[Kind]string = map[Kind]string{
	ST_UNDEFINED:  ``,
	ST_VOID:       `Void`,
	ST_IDENTIFIER: `Identifier`,
	ST_LITERAL:    `Literal`,
}

func (k Kind) String() string {
	return kinds[k]
}

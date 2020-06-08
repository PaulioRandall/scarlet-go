package statement

import (
//"strings"
)

type Kind int

const (
	ST_UNDEFINED Kind = iota
	ST_VOID
	ST_IDENTIFIER
)

var kinds map[Kind]string = map[Kind]string{
	ST_UNDEFINED:  ``,
	ST_VOID:       `Void`,
	ST_IDENTIFIER: `Identifier`,
}

func (k Kind) String() string {
	return kinds[k]
}

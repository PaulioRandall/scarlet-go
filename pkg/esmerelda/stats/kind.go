package stats

type Kind int

const (
	ST_UNDEFINED Kind = iota
	ST_EXIT
	ST_VOID
	ST_IDENTIFIER
	ST_LITERAL
	ST_COLLECTION_ACCESSOR
	ST_NEGATION
	ST_OPERATION
	ST_ASSIGNMENT
	ST_ASSIGNMENT_BLOCK
	ST_BLOCK
	ST_EXPR_FUNC
	ST_PARAMETERS
	ST_FUNCTION
	ST_FUNCTION_CALL
	ST_WATCH
	ST_GUARD
	ST_WHEN_CASE
	ST_WHEN
	ST_LOOP
	ST_SPELL_CALL
	ST_EXISTS
)

var kinds map[Kind]string = map[Kind]string{
	ST_UNDEFINED:           ``,
	ST_EXIT:                `Exit`,
	ST_VOID:                `Void`,
	ST_IDENTIFIER:          `Identifier`,
	ST_LITERAL:             `Literal`,
	ST_COLLECTION_ACCESSOR: `Collection-Accessor`,
	ST_NEGATION:            `Negation`,
	ST_OPERATION:           `Operation`,
	ST_ASSIGNMENT:          `Assignment`,
	ST_ASSIGNMENT_BLOCK:    `Assignment-Block`,
	ST_BLOCK:               `Block`,
	ST_EXPR_FUNC:           `Expression-Function`,
	ST_PARAMETERS:          `Parameters`,
	ST_FUNCTION:            `Function`,
	ST_FUNCTION_CALL:       `Function-Call`,
	ST_WATCH:               `Watch`,
	ST_GUARD:               `Guard`,
	ST_WHEN_CASE:           `WhenCase`,
	ST_WHEN:                `When`,
	ST_LOOP:                `Loop`,
	ST_SPELL_CALL:          `Spell-Call`,
	ST_EXISTS:              `Exists`,
}

func (k Kind) String() string {
	return kinds[k]
}

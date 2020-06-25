package stats

type Kind int

const (
	ST_UNDEFINED Kind = iota
	ST_EXIT
	ST_VOID
	ST_IDENTIFIER
	ST_LITERAL
	ST_CONTAINER_ITEM
	ST_NEGATION
	ST_OPERATION
	ST_ASSIGNMENT
	ST_ASSIGNMENT_BLOCK
	ST_BLOCK
	ST_EXPR_FUNC
	ST_PARAMETERS
	ST_FUNC_DEF
	ST_FUNC_CALL
	ST_WATCH
	ST_GUARD
	ST_WHEN_CASE
	ST_WHEN
	ST_LOOP
	ST_SPELL_CALL
	ST_EXISTS
)

var kinds map[Kind]string = map[Kind]string{
	ST_UNDEFINED:        ``,
	ST_EXIT:             `Exit`,
	ST_VOID:             `Void`,
	ST_IDENTIFIER:       `Identifier`,
	ST_LITERAL:          `Literal`,
	ST_CONTAINER_ITEM:   `Container-Item`,
	ST_NEGATION:         `Negation`,
	ST_OPERATION:        `Operation`,
	ST_ASSIGNMENT:       `Assignment`,
	ST_ASSIGNMENT_BLOCK: `Assignment-Block`,
	ST_BLOCK:            `Block`,
	ST_EXPR_FUNC:        `Expression-Function`,
	ST_PARAMETERS:       `Parameters`,
	ST_FUNC_DEF:         `Function-Def`,
	ST_FUNC_CALL:        `Function-Call`,
	ST_WATCH:            `Watch`,
	ST_GUARD:            `Guard`,
	ST_WHEN_CASE:        `WhenCase`,
	ST_WHEN:             `When`,
	ST_LOOP:             `Loop`,
	ST_SPELL_CALL:       `Spell-Call`,
	ST_EXISTS:           `Exists`,
}

func (k Kind) String() string {
	return kinds[k]
}

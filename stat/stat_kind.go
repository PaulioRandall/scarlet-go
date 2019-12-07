package stat

// Kind represents a statement type.
type Kind int

const (
	UNDEFINED Kind = iota
	// ------------------
	ASSIGN_FUNC
	ASSIGN_ID
	INVOKE_SPELL
	INVOKE_FUNC
)

var kindNames map[Kind]string = map[Kind]string{
	ASSIGN_FUNC:  `ASSIGN_FUNC`,
	ASSIGN_ID:    `ASSIGN_ID`,
	INVOKE_SPELL: `ASSIGN_SPELL`,
	INVOKE_FUNC:  `ASSIGN_FUNC`,
}

// Name returns the name of the token type.
func (k Kind) Name() string {
	s := kindNames[k]
	if s == `` {
		return `UNDEFINED`
	}
	return s
}

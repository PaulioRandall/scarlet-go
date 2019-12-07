package stat

// StatKind represents a statement type.
type StatKind int

const (
	UNDEFINED StatKind = iota
	// ------------------
	ASSIGN_FUNC
	ASSIGN_ID
	INVOKE_SPELL
	INVOKE_FUNC
)

var kindNames map[StatKind]string = map[StatKind]string{
	ASSIGN_FUNC:  `ASSIGN_FUNC`,
	ASSIGN_ID:    `ASSIGN_ID`,
	INVOKE_SPELL: `ASSIGN_SPELL`,
	INVOKE_FUNC:  `ASSIGN_FUNC`,
}

// Name returns the name of the token type.
func (sk StatKind) Name() string {
	s := kindNames[sk]
	if s == `` {
		return `UNDEFINED`
	}
	return s
}

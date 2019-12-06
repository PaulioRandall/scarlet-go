package token

// Kind represents a token type.
type Kind int

const (
	UNDEFINED Kind = iota
	// ------------------
	WHITESPACE
	NEWLINE
	FUNC
	DO
	END
	ID
	ID_DELIM
	ASSIGN
	OPEN_PAREN
	CLOSE_PAREN
	SPELL
)

var kindNames map[Kind]string = map[Kind]string{
	WHITESPACE:  `WHITESPACE`,
	NEWLINE:     `NEWLINE`,
	FUNC:        `FUNC`,
	DO:          `DO`,
	END:         `END`,
	ID:          `ID`,
	ID_DELIM:    `ID_DELIM`,
	ASSIGN:      `ASSIGN`,
	OPEN_PAREN:  `OPEN_PAREN`,
	CLOSE_PAREN: `CLOSE_PAREN`,
	SPELL:       `SPELL`,
}

// Name returns the name of the token type.
func (k Kind) Name() string {
	s := kindNames[k]
	if s == `` {
		return `UNDEFINED`
	}
	return s
}

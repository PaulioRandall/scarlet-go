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
	ASSIGN
	OPEN_PAREN
	CLOSE_PAREN
	ID_DELIM
)

var kindNames map[Kind]string = map[Kind]string{
	WHITESPACE:  `WHITESPACE`,
	NEWLINE:     `NEWLINE`,
	FUNC:        `FUNC`,
	DO:          `DO`,
	END:         `END`,
	ID:          `ID`,
	ASSIGN:      `ASSIGN`,
	OPEN_PAREN:  `OPEN_PAREN`,
	CLOSE_PAREN: `CLOSE_PAREN`,
	ID_DELIM:    `ID_DELIM`,
}

// Name returns the name of the token type.
func (k Kind) Name() string {
	return kindNames[k]
}

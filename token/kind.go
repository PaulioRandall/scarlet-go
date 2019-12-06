package token

// Kind represents a token type.
type Kind int

const (
	UNDEFINED Kind = iota
	// ------------------
	WHITESPACE
	NEWLINE
	FUNC
	END
	ID
)

var kindNames map[Kind]string = map[Kind]string{
	WHITESPACE: `WHITESPACE`,
	NEWLINE:    `NEWLINE`,
	FUNC:       `FUNC`,
	END:        `END`,
	ID:         `ID`,
}

// Name returns the name of the token type.
func (k Kind) Name() string {
	return kindNames[k]
}

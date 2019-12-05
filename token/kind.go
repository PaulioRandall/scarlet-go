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

// kindNames is a mapping of Kinds to their string name
var kindNames map[Kind]string = map[Kind]string{
	UNDEFINED:  `UNDEFINED`,
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

// FindWordKind identifies the kind of the word string or UNDEFINED if one
// could not be identified.
func FindWordKind(s string) Kind {
	for k, v := range kindNames {
		if v == s {
			return k
		}
	}

	return UNDEFINED
}

package token

// Kind represents a token type.
type Kind int

const (
	UNDEFINED Kind = iota
	// ------------------
	WHITESPACE
	NEWLINE
	PROCEDURE
	END
)

// Name returns the name of the token type.
func (k Kind) Name() string {
	return KindName(k)
}

// KindName returns the name of the input kind.
func KindName(k Kind) string {
	switch k {
	case UNDEFINED:
		return `UNDEFINED`
	case NEWLINE:
		return `NEWLINE`
	case WHITESPACE:
		return `WHITESPACE`
	case PROCEDURE:
		return `PROCEDURE`
	case END:
		return `END`
	}

	return `--UNKOWN--`
}

// FindWordKind identifies the kind of the word string.
func FindWordKind(s string) (k Kind) {

	switch s {
	case `PROCEDURE`:
		k = PROCEDURE
	case `END`:
		k = END
	default:
		k = UNDEFINED
	}

	return
}

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
)

// Name returns the name of the token type.
func (k Kind) Name() string {
	switch k {
	case UNDEFINED:
		return `UNDEFINED`
	case NEWLINE:
		return `NEWLINE`
	case WHITESPACE:
		return `WHITESPACE`
	case FUNC:
		return `FUNC`
	case END:
		return `END`
	default:
		return `--UNKOWN--`
	}
}

// FindWordKind identifies the kind of the word string.
func FindWordKind(s string) Kind {
	switch s {
	case `FUNC`:
		return FUNC
	case `END`:
		return END
	default:
		return UNDEFINED
	}
}

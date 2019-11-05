package token

// Kind represents a token type.
type Kind int

const (
	UNDEFINED Kind = iota
	// ------------------
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
	case PROCEDURE:
		return `PROCEDURE`
	case END:
		return `END`
	}

	return `--UNKOWN--`
}

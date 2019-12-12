package token

// Kind represents a token type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	COMMENT     Kind = `COMMENT`
	WHITESPACE  Kind = `WHITESPACE`
	NEWLINE     Kind = `NEWLINE`
	FUNC        Kind = `FUNC`
	DO          Kind = `DO`
	END         Kind = `END`
	ID          Kind = `ID`
	ID_DELIM    Kind = `ID_DELIM`
	ASSIGN      Kind = `ASSIGN`
	OPEN_PAREN  Kind = `OPEN_PAREN`
	CLOSE_PAREN Kind = `CLOSE_PAREN`
	SPELL       Kind = `SPELL`
	STR_LITERAL Kind = `STR_LITERAL`
)

// Token represents a grammer token within a source file.
type Token interface {

	// Kind returns the type of the token.
	Kind() Kind

	// Value returns the string representing the token in source.
	Value() string

	// Where returns the token location within the source.
	Where() Snippet
}

// tokenImpl is a simple implementation of the Token interface.
type tokenImpl struct {
	k Kind
	v string
	s Snippet
}

// NewToken creates a new token.
func NewToken(k Kind, v string, line, start, end int) Token {
	return tokenImpl{
		k: k,
		v: v,
		s: NewSnippet(line, start, end),
	}
}

// TokenBySnippet creates a new token using a Snippet as the location parameter.
func TokenBySnippet(k Kind, v string, s Snippet) Token {
	return tokenImpl{
		k: k,
		v: v,
		s: s,
	}
}

// Kind satisfies the Token interface.
func (t tokenImpl) Kind() Kind {
	return t.k
}

// Value satisfies the Token interface.
func (t tokenImpl) Value() string {
	return t.v
}

// Where satisfies the Token interface.
func (t tokenImpl) Where() Snippet {
	return t.s
}

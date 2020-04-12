package snippet

type Kind string

const (
	SNIPPET_UNDEFINED Kind = ``
	SNIPPET_EOF       Kind = `EOF_`
	SNIPPET_FUNC      Kind = `FUNC`
)

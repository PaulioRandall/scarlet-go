package z_token

type Kind int

const (
	K_UNDEFINED Kind = iota
	// ------------------
	K_ANY
	K_ANOTHER
	// ------------------
	K_EOF
	// ------------------
	K_REDUNDANT
	K_NEWLINE
	K_KEYWORD
	K_IDENTIFIER
	K_DELIMITER
	K_LITERAL
	K_ARITHMETIC
	K_LOGIC
	K_COMPARISON
	K_REFERENCE
)

var kinds map[Kind]string = map[Kind]string{
	K_UNDEFINED:  ``,
	K_ANOTHER:    `ANOTHER`,
	K_EOF:        `EOF`,
	K_REDUNDANT:  `REDUNDANT`,
	K_NEWLINE:    `NEWLINE`,
	K_KEYWORD:    `KEYWORD`,
	K_IDENTIFIER: `IDENTIFIER`,
	K_DELIMITER:  `DELIMITER`,
	K_LITERAL:    `LITERAL`,
	K_ARITHMETIC: `ARITHMETIC`,
	K_LOGIC:      `LOGIC`,
	K_COMPARISON: `COMPARISON`,
	K_REFERENCE:  `REFERENCE`,
}

func (k Kind) String() string {
	return kinds[k]
}

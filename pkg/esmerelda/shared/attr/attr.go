package attr

type Attr int

func (at Attr) String() string {
	return attrs[at]
}

const (
	ATTR_UNDEFINED Attr = iota
	ATTR_ANY
	// ------------------
	ATTR_REDUNDANT // Whitespace or comment
	ATTR_WHITESPACE
	ATTR_COMMENT
	ATTR_TERMINATOR // newline or semicolon
	ATTR_NEWLINE
	ATTR_LITERAL // bool, number, or string
	ATTR_TERM    // Literal or identifier
	ATTR_BOOL    // true
	ATTR_NUMBER  // 1
	ATTR_STRING  // "abc"
	ATTR_IDENTIFIER
	ATTR_ASSIGNEE  // Identifier, void, or list item
	ATTR_VOID      // _
	ATTR_SPELL     // @
	ATTR_PARAMETER // Magic token
	ATTR_DELIMITER // Comma or bracket
	ATTR_SEPARATOR // ,
	ATTR_PARENTHESIS
	ATTR_OPENER
	ATTR_CLOSER
)

var attrs = map[Attr]string{
	ATTR_ANY: `ANY`,
	// ------------------
	ATTR_REDUNDANT:   `REDUNDANT`,
	ATTR_WHITESPACE:  `WHITESPACE`,
	ATTR_COMMENT:     `COMMENT`,
	ATTR_TERMINATOR:  `TERMINATOR`,
	ATTR_NEWLINE:     `NEWLINE`,
	ATTR_LITERAL:     `LITERAL`,
	ATTR_TERM:        `TERM`,
	ATTR_BOOL:        `BOOL`,
	ATTR_NUMBER:      `NUMBER`,
	ATTR_STRING:      `STRING`,
	ATTR_IDENTIFIER:  `IDENTIFIER`,
	ATTR_ASSIGNEE:    `ASSIGNEE`,
	ATTR_VOID:        `VOID`,
	ATTR_SPELL:       `SPELL`,
	ATTR_PARAMETER:   `PARAMETER`,
	ATTR_DELIMITER:   `DELIMITER`,
	ATTR_SEPARATOR:   `SEPARATOR`,
	ATTR_PARENTHESIS: `PARENTHESIS`,
	ATTR_OPENER:      `OPENER`,
	ATTR_CLOSER:      `CLOSER`,
}

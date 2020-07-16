package prop

import (
	"strings"
)

type Prop int

type Entity interface {
	Props() []Prop
	Is(Prop) bool
	IsNot(Prop) bool
}

func (p Prop) String() string {
	return props[p]
}

func JoinProps(infix string, props ...Prop) string {

	sb := strings.Builder{}

	for i, p := range props {
		if i != 0 {
			sb.WriteString(infix)
		}

		sb.WriteString(p.String())
	}

	return sb.String()
}

const (
	PR_UNDEFINED Prop = iota
	PR_ANY
	// -----------------
	PR_REDUNDANT // Whitespace or comment
	PR_WHITESPACE
	PR_COMMENT
	PR_TERMINATOR // newline or semicolon
	PR_NEWLINE
	PR_LITERAL // bool, number, or string
	PR_TERM    // Literal or identifier
	PR_BOOL    // true
	PR_NUMBER  // 1
	PR_STRING  // "abc"
	PR_IDENTIFIER
	PR_ASSIGNEE  // Identifier, void, or list item
	PR_VOID      // _
	PR_CALLABLE  // Spell or function
	PR_SPELL     // @
	PR_PARAMETER // Magic token
	PR_DELIMITER // Comma or bracket
	PR_SEPARATOR // ,
	PR_PARENTHESIS
	PR_OPENER
	PR_CLOSER
)

var props = map[Prop]string{
	PR_ANY: `ANY`,
	// ------------------
	PR_REDUNDANT:   `REDUNDANT`,
	PR_WHITESPACE:  `WHITESPACE`,
	PR_COMMENT:     `COMMENT`,
	PR_TERMINATOR:  `TERMINATOR`,
	PR_NEWLINE:     `NEWLINE`,
	PR_LITERAL:     `LITERAL`,
	PR_TERM:        `TERM`,
	PR_BOOL:        `BOOL`,
	PR_NUMBER:      `NUMBER`,
	PR_STRING:      `STRING`,
	PR_IDENTIFIER:  `IDENTIFIER`,
	PR_ASSIGNEE:    `ASSIGNEE`,
	PR_VOID:        `VOID`,
	PR_CALLABLE:    `CALLABLE`,
	PR_SPELL:       `SPELL`,
	PR_PARAMETER:   `PARAMETER`,
	PR_DELIMITER:   `DELIMITER`,
	PR_SEPARATOR:   `SEPARATOR`,
	PR_PARENTHESIS: `PARENTHESIS`,
	PR_OPENER:      `OPENER`,
	PR_CLOSER:      `CLOSER`,
}

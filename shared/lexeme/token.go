package lexeme

import (
	"strings"
)

type Token int

const (
	UNDEFINED Token = iota
	// -----------------
	WHITESPACE
	COMMENT     // # comment
	TERMINATOR  // ;
	NEWLINE     // \n
	BOOL        // true | false
	NUMBER      // 1
	STRING      // "abc"
	IDENTIFIER  // abc
	SPELL       // @abc
	SEPARATOR   // ,
	LEFT_PAREN  // (
	RIGHT_PAREN // )
	CALLABLE    // Magic token, tells compiler that spell or func args are coming
	ASSIGNMENT  // :=
	ADD         // +
	SUB         // -
	MUL         // *
	DIV         // /
	REM         // %
)

var tokens = map[Token]string{
	WHITESPACE:  `WHITESPACE`,
	COMMENT:     `COMMENT`,
	TERMINATOR:  `TERMINATOR`,
	NEWLINE:     `NEWLINE`,
	BOOL:        `BOOL`,
	NUMBER:      `NUMBER`,
	STRING:      `STRING`,
	IDENTIFIER:  `IDENTIFIER`,
	SPELL:       `SPELL`,
	SEPARATOR:   `SEPARATOR`,
	LEFT_PAREN:  `LEFT_PAREN`,
	RIGHT_PAREN: `RIGHT_PAREN`,
	CALLABLE:    `CALLABLE`,
	ASSIGNMENT:  `ASSIGNMENT`,
	ADD:         `ADDITION`,
	SUB:         `SUBTRACTION`,
	MUL:         `MULTIPLICATION`,
	DIV:         `DIVISION`,
	REM:         `REMAINDER`,
}

func (tk Token) IsAny(others ...Token) bool {

	for _, o := range others {
		if tk == o {
			return true
		}
	}

	return false
}

func (tk Token) IsRedundant() bool {
	return tk == WHITESPACE || tk == COMMENT
}

func (tk Token) IsTerminator() bool {
	return tk == TERMINATOR || tk == NEWLINE
}

func (tk Token) IsLiteral() bool {
	return tk == BOOL || tk == NUMBER || tk == STRING
}

func (tk Token) IsTerm() bool {
	return tk == IDENTIFIER || tk.IsLiteral()
}

func (tk Token) IsAssignee() bool {
	return tk == IDENTIFIER
}

func (tk Token) IsOperator() bool {
	return tk == ADD ||
		tk == SUB ||
		tk == MUL ||
		tk == DIV ||
		tk == REM
}

func (tk Token) IsOpener() bool {
	return tk == LEFT_PAREN
}

func (tk Token) IsCloser() bool {
	return tk == RIGHT_PAREN
}

func (tk Token) String() string {
	return tokens[tk]
}

func JoinTokens(infix string, tks ...Token) string {

	sb := strings.Builder{}

	for i, tk := range tks {
		if i != 0 {
			sb.WriteString(infix)
		}

		sb.WriteString(tk.String())
	}

	return sb.String()
}

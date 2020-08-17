package lexeme

import (
	"strings"
)

type Token int

const (
	UNDEFINED Token = iota
	// -----------------
	SPACE
	COMMENT    // # comment
	TERMINATOR // ;
	NEWLINE    // \n
	BOOL       // true | false
	NUMBER     // 1
	STRING     // "abc"
	IDENT      // abc
	SPELL      // @abc
	DELIM      // ,
	L_PAREN    // (
	R_PAREN    // )
	L_SQUARE   // [
	R_SQUARE   // ]
	L_CURLY    // {
	R_CURLY    //	}
	ASSIGN     // :=
	ADD        // +
	SUB        // -
	MUL        // *
	DIV        // /
	REM        // %
	AND        // &&
	OR         // ||
	LESS       // <
	MORE       // >
	LESS_EQUAL // <=
	MORE_EQUAL // >=
	EQUAL      // ==
	NOT_EQUAL  // !=
	GUARD      // Magic: Indicates the subsequent block is conditional
)

var tokens = map[Token]string{
	SPACE:      `SPACE`,
	COMMENT:    `COMMENT`,
	TERMINATOR: `TERMINATOR`,
	NEWLINE:    `NEWLINE`,
	BOOL:       `BOOL`,
	NUMBER:     `NUMBER`,
	STRING:     `STRING`,
	IDENT:      `IDENT`,
	SPELL:      `SPELL`,
	DELIM:      `DELIM`,
	L_PAREN:    `L_PAREN`,
	R_PAREN:    `R_PAREN`,
	L_SQUARE:   `L_SQUARE`,
	R_SQUARE:   `R_SQUARE`,
	L_CURLY:    `L_CURLY`,
	R_CURLY:    `R_CURLY`,
	ASSIGN:     `ASSIGN`,
	ADD:        `ADD`,
	SUB:        `SUB`,
	MUL:        `MUL`,
	DIV:        `DIV`,
	REM:        `REM`,
	AND:        `AND`,
	OR:         `OR`,
	LESS:       `LESS`,
	MORE:       `MORE`,
	LESS_EQUAL: `LESS_EQUAL`,
	MORE_EQUAL: `MORE_EQUAL`,
	EQUAL:      `EQUAL`,
	NOT_EQUAL:  `NOT_EQUAL`,
	GUARD:      `GUARD`,
}

func (tk Token) Precedence() int {
	switch tk {
	case L_PAREN, R_PAREN:
		return 7
	case MUL, DIV, REM:
		return 6
	case ADD, SUB:
		return 5
	case LESS, MORE, LESS_EQUAL, MORE_EQUAL:
		return 4
	case EQUAL, NOT_EQUAL:
		return 3
	case AND:
		return 2
	case OR:
		return 1
	}

	return 0
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
	return tk == SPACE || tk == COMMENT
}

func (tk Token) IsTerminator() bool {
	return tk == TERMINATOR || tk == NEWLINE
}

func (tk Token) IsLiteral() bool {
	return tk == BOOL || tk == NUMBER || tk == STRING
}

func (tk Token) IsTerm() bool {
	return tk == IDENT || tk.IsLiteral()
}

func (tk Token) IsAssignee() bool {
	return tk == IDENT
}

func (tk Token) IsOperator() bool {
	return tk == MUL ||
		tk == DIV ||
		tk == REM ||
		tk == ADD ||
		tk == SUB ||
		tk == LESS ||
		tk == MORE ||
		tk == LESS_EQUAL ||
		tk == MORE_EQUAL ||
		tk == EQUAL ||
		tk == NOT_EQUAL ||
		tk == AND ||
		tk == OR
}

func (tk Token) IsOpener() bool {
	return tk == L_PAREN || tk == L_SQUARE
}

func (tk Token) IsCloser() bool {
	return tk == R_PAREN || tk == R_SQUARE
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

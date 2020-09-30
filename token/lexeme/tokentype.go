package lexeme

import (
	"strings"
)

type TokenType string

const (
	UNDEFINED TokenType = ``
	// -----------------
	SPACE      TokenType = `SPACE`
	COMMENT    TokenType = `COMMENT`    // # comment
	TERMINATOR TokenType = `TERMINATOR` // ;
	NEWLINE    TokenType = `NEWLINE`    // \n
	BOOL       TokenType = `BOOL`       /*Retired!*/
	TRUE       TokenType = `TRUE`       // true
	FALSE      TokenType = `FALSE`      // false
	NUMBER     TokenType = `NUMBER`     // 1
	STRING     TokenType = `STRING`     // "abc"
	IDENT      TokenType = `IDENT`      // abc
	SPELL      TokenType = `SPELL`      // @abc
	GUARD      TokenType = `GUARD`      /*Retired!*/
	LOOP       TokenType = `LOOP`       // loop
	DELIM      TokenType = `DELIM`      // ,
	L_PAREN    TokenType = `L_PAREN`    // (
	R_PAREN    TokenType = `R_PAREN`    // )
	L_SQUARE   TokenType = `L_SQUARE`   // [
	R_SQUARE   TokenType = `R_SQUARE`   // ]
	L_CURLY    TokenType = `L_CURLY`    // {
	R_CURLY    TokenType = `R_CURLY`    //	}
	ASSIGN     TokenType = `ASSIGN`     // :=
	VOID       TokenType = `VOID`       // _
	ADD        TokenType = `ADD`        // +
	SUB        TokenType = `SUB`        // -
	MUL        TokenType = `MUL`        // *
	DIV        TokenType = `DIV`        // /
	REM        TokenType = `REM`        // %
	AND        TokenType = `AND`        // &&
	OR         TokenType = `OR`         // ||
	LESS       TokenType = `LESS`       // <
	MORE       TokenType = `MORE`       // >
	LESS_EQUAL TokenType = `LESS_EQUAL` // <=
	MORE_EQUAL TokenType = `MORE_EQUAL` // >=
	EQUAL      TokenType = `EQUAL`      // ==
	NOT_EQUAL  TokenType = `NOT_EQUAL`  // !=
)

func Identify(s string) TokenType {
	switch s {
	case "TRUE":
		return TRUE
	case "FALSE":
		return FALSE
	case "LOOP":
		return LOOP
	}

	return IDENT
}

func (tk TokenType) Precedence() int {
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

func (tk TokenType) IsAny(others ...TokenType) bool {

	for _, o := range others {
		if tk == o {
			return true
		}
	}

	return false
}

func (tk TokenType) IsRedundant() bool {
	return tk == SPACE || tk == COMMENT
}

func (tk TokenType) IsTerminator() bool {
	return tk == TERMINATOR || tk == NEWLINE
}

func (tk TokenType) IsLiteral() bool {
	return tk == BOOL || tk == NUMBER || tk == STRING
}

func (tk TokenType) IsBool() bool {
	return tk == TRUE || tk == FALSE
}

func (tk TokenType) IsTerm() bool {
	return tk == IDENT || tk.IsLiteral()
}

func (tk TokenType) IsAssignee() bool {
	return tk == IDENT || tk == VOID
}

func (tk TokenType) IsOperator() bool {
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

func (tk TokenType) IsOpener() bool {
	return tk == L_PAREN || tk == L_SQUARE || tk == L_CURLY
}

func (tk TokenType) IsCloser() bool {
	return tk == R_PAREN || tk == R_SQUARE || tk == R_CURLY
}

func (tk TokenType) String() string {
	return string(tk)
}

func JoinTokens(infix string, tks ...TokenType) string {

	sb := strings.Builder{}

	for i, tk := range tks {
		if i != 0 {
			sb.WriteString(infix)
		}

		sb.WriteString(tk.String())
	}

	return sb.String()
}

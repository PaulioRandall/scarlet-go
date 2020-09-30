package token

import (
	"strings"
)

type Token string

const (
	UNDEFINED Token = ``
	// -----------------
	SPACE      Token = `SPACE`
	COMMENT    Token = `COMMENT`    // # comment
	TERMINATOR Token = `TERMINATOR` // ;
	NEWLINE    Token = `NEWLINE`    // \n
	BOOL       Token = `BOOL`       /*Retired!*/
	TRUE       Token = `TRUE`       // true
	FALSE      Token = `FALSE`      // false
	NUMBER     Token = `NUMBER`     // 1
	STRING     Token = `STRING`     // "abc"
	IDENT      Token = `IDENT`      // abc
	SPELL      Token = `SPELL`      // @abc
	GUARD      Token = `GUARD`      /*Retired!*/
	LOOP       Token = `LOOP`       // loop
	DELIM      Token = `DELIM`      // ,
	L_PAREN    Token = `L_PAREN`    // (
	R_PAREN    Token = `R_PAREN`    // )
	L_SQUARE   Token = `L_SQUARE`   // [
	R_SQUARE   Token = `R_SQUARE`   // ]
	L_CURLY    Token = `L_CURLY`    // {
	R_CURLY    Token = `R_CURLY`    //	}
	ASSIGN     Token = `ASSIGN`     // :=
	VOID       Token = `VOID`       // _
	ADD        Token = `ADD`        // +
	SUB        Token = `SUB`        // -
	MUL        Token = `MUL`        // *
	DIV        Token = `DIV`        // /
	REM        Token = `REM`        // %
	AND        Token = `AND`        // &&
	OR         Token = `OR`         // ||
	LESS       Token = `LESS`       // <
	MORE       Token = `MORE`       // >
	LESS_EQUAL Token = `LESS_EQUAL` // <=
	MORE_EQUAL Token = `MORE_EQUAL` // >=
	EQUAL      Token = `EQUAL`      // ==
	NOT_EQUAL  Token = `NOT_EQUAL`  // !=
)

func IdentifyWord(s string) Token {
	switch s {
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "loop":
		return LOOP
	}

	return IDENT
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
	return tk == NUMBER || tk == STRING || tk.IsBool()
}

func (tk Token) IsBool() bool {
	return tk == TRUE || tk == FALSE
}

func (tk Token) IsTerm() bool {
	return tk == IDENT || tk.IsLiteral()
}

func (tk Token) IsAssignee() bool {
	return tk == IDENT || tk == VOID
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
	return tk == L_PAREN || tk == L_SQUARE || tk == L_CURLY
}

func (tk Token) IsCloser() bool {
	return tk == R_PAREN || tk == R_SQUARE || tk == R_CURLY
}

func (tk Token) String() string {
	return string(tk)
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

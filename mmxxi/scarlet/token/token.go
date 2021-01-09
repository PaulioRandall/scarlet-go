package token

import (
	"strings"
)

type Token int

const (
	UNDEFINED Token = iota

	// Redundant
	SPACE   // whitespace
	COMMENT // # comment

	// Identifiers
	IDENT // identifier

	// Literals
	BOOL // 'true' or 'false'
	NUM  // 1
	STR  // "string"

	// Keywords
	T_BOOL // B
	T_NUM  // N
	T_STR  // S
	E_FUNC // E
	FUNC   // F
	LOOP   // loop
	MATCH  // match
	TYPE   // type

	// Operators
	DEFINE // :=
	ASSIGN // <-
	OUTPUT // ->

	ADD // +
	SUB // -, ¬
	MUL // *
	DIV // /
	REM // %

	AND // &&
	OR  // ||

	EQU // ==
	NEQ // !=
	LT  // <
	MT  // >
	LTE // <=
	MTE // >=

	NOT // !
	QUE // ?

	// Delimiters
	TERMINATOR // ';' or '\n'
	SPELL      // @
	DELIM      // ,
	REF        // :

	L_PAREN // (
	R_PAREN // )
	L_BRACK // [
	R_BRACK // ]
	L_BRACE // {
	R_BRACE // }
)

// IdentifyWord returns the Token represented by the 's'. If 's' does not match
// a keyword then the IDENT Token is returned.
func IdentifyWord(s string) Token {
	switch s {
	case "true", "false":
		return BOOL
	case "B":
		return T_BOOL
	case "N":
		return T_NUM
	case "S":
		return T_STR
	case "E":
		return E_FUNC
	case "F":
		return FUNC
	case "loop":
		return LOOP
	case "match":
		return MATCH
	case "type":
		return TYPE
	}

	return IDENT
}

// Precedence returns a number representing the priority of the Token in
// comparison to other Tokens whereby a higher number signifies a greater
// precedence. Upon equal precedence, left always has priority.
func (tk Token) Precedence() int {
	switch tk {
	case L_PAREN, R_PAREN, L_BRACK, R_BRACK, L_BRACE, R_BRACE:
		return 7
	case MUL, DIV, REM:
		return 6
	case ADD, SUB:
		return 5
	case LT, MT, LTE, MTE:
		return 4
	case EQU, NEQ:
		return 3
	case AND:
		return 2
	case OR:
		return 1
	}

	return 0
}

// IsRedundant returns true if the Token is redundant to the parsing process.
func (tk Token) IsRedundant() bool {
	return tk == SPACE || tk == COMMENT
}

// IsLiteral returns true if the Token represents a literal or constant value
// such as a bool, number, or string.
func (tk Token) IsLiteral() bool {
	return tk == BOOL || tk == NUM || tk == STR
}

// IsTerm returns true if the Token represents a literal or an identifier.
func (tk Token) IsTerm() bool {
	return tk == IDENT || tk.IsLiteral()
}

// IsOpener returns true if the Token represents an opening bracket of any sort.
func (tk Token) IsOpener() bool {
	return tk == L_PAREN || tk == L_BRACK || tk == L_BRACE
}

// IsCloser returns true if the Token represents an closing bracket of any sort.
func (tk Token) IsCloser() bool {
	return tk == R_PAREN || tk == R_BRACK || tk == R_BRACE
}

// IsOperator returns true if the Token represents a arithmetic, logical, or
// boolean operator. All operators have a precedence of 1 or greater.
func (tk Token) IsOperator() bool {
	return tk.IsInfix() || tk.IsPrefix() || tk.IsPostfix()
}

// IsPrefix returns true if the Token represents a prefix operator.
func (tk Token) IsPrefix() bool {
	return tk.IsOpener() || tk == NOT
}

// IsPostfix returns true if the Token represents a postfix operator.
func (tk Token) IsPostfix() bool {
	return tk.IsCloser() || tk == QUE
}

// IsInfix returns true if the Token represents an infix operator.
func (tk Token) IsInfix() bool {
	return tk == ADD ||
		tk == SUB ||
		tk == MUL ||
		tk == DIV ||
		tk == REM ||
		tk == EQU ||
		tk == NEQ ||
		tk == LT ||
		tk == MT ||
		tk == LTE ||
		tk == MTE ||
		tk == AND ||
		tk == OR
}

// IsType returns true if the Token represents a type.
func (tk Token) IsType() bool {
	return tk == T_BOOL || tk == T_NUM || tk == T_STR
}

// String returns the human readable string representation of the Token.
func (tk Token) String() string {
	switch tk {

	// Redundant
	case SPACE:
		return "SPACE"
	case COMMENT:
		return "COMMENT"

		// Identifiers
	case IDENT:
		return "IDENT"

		// Literals
	case BOOL:
		return "BOOL"
	case NUM:
		return "NUM"
	case STR:
		return "STR"

		// Keywords
	case T_BOOL:
		return "T_BOOL"
	case T_NUM:
		return "T_NUM"
	case T_STR:
		return "T_STR"
	case E_FUNC:
		return "E_FUNC"
	case FUNC:
		return "FUNC"
	case LOOP:
		return "LOOP"
	case MATCH:
		return "MATCH"
	case TYPE:
		return "TYPE"

		// Operators
	case DEFINE:
		return "DEFINE"
	case ASSIGN:
		return "ASSIGN"
	case OUTPUT:
		return "OUTPUT"

	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case REM:
		return "REM"

	case AND:
		return "AND"
	case OR:
		return "OR"

	case EQU:
		return "EQU"
	case NEQ:
		return "NEQ"
	case LT:
		return "LT"
	case MT:
		return "MT"
	case LTE:
		return "LTE"
	case MTE:
		return "MTE"

	case NOT:
		return "NOT"
	case QUE:
		return "QUE"

		// Delimiters
	case TERMINATOR:
		return "TERMINATOR"
	case SPELL:
		return "SPELL"
	case DELIM:
		return "DELIM"
	case REF:
		return "REF"

	case L_PAREN:
		return "L_BRACE"
	case R_PAREN:
		return "L_BRACE"
	case L_BRACK:
		return "L_BRACE"
	case R_BRACK:
		return "R_BRACE"
	case L_BRACE:
		return "L_BRACE"
	case R_BRACE:
		return "R_BRACE"
	}

	return "UNDEFINED"
}

// Join concaternates the string representations of each Token in 'tks'
// inserting the supplied infix between items.
func Join(infix string, tks ...Token) string {

	sb := strings.Builder{}

	for i, tk := range tks {
		if i != 0 {
			sb.WriteString(infix)
		}

		sb.WriteString(tk.String())
	}

	return sb.String()
}

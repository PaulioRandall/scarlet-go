package token

import (
	"strings"
)

type Token int

const (
	UNDEFINED Token = iota

	// Reference
	SOF
	EOF

	// Redundant
	SPACE   // whitespace
	COMMENT // # comment

	// Identifiers
	IDENT // identifier
	VOID  // _

	// Literals
	BOOL // true or false
	NUM  // 1
	STR  // "string"

	// Keywords
	E_FUNC // expression function
	FUNC   // function
	LOOP   // loop
	MATCH  // match
	TYPE   // type

	// Operators
	ASSIGN // <-
	DEFINE // :=
	INTO   // ->

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
	LEQ // <=
	MEQ // >=

	NOT // !
	QUE // ?

	// Delimiters
	TERMINATOR // ;, '\n'
	SPELL      // @
	DELIM      // ,
	REF        // :
	DOT        // .

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
	case "E":
		return E_FUNC
	case "F":
		return FUNC
	case "loop":
		return LOOP
	case "match":
		return MATCH
	case "B":
		return BOOL
	case "N":
		return NUM
	case "S":
		return STR
	case "type":
		return TYPE
	}

	return IDENT
}

// Precedence returns a number representing the priority of the Token in
// comparison to other Tokens whereby a higher number signifies a greater
// precedence.
func (tk Token) Precedence() int {
	switch tk {
	case L_PAREN, R_PAREN:
		return 7
	case MUL, DIV, REM:
		return 6
	case ADD, SUB:
		return 5
	case LT, MT, LEQ, MEQ:
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

// IsAssignee returns true of the Token can be used as the target of an
// assignment.
func (tk Token) IsAssignee() bool {
	return tk == IDENT || tk == VOID
}

// IsOperator returns true if the Token represents a arithmetic, logical, or
// boolean operator. All operators have a precedence of 1 or greater.
//
// Note that parenthesis are not currently defined as operators unlike other
// compilers might. This is just the way the compiler is built and may be
// subject to change later.
func (tk Token) IsOperator() bool {
	return tk.IsInfixOperator() || tk == NOT || tk == QUE
}

// IsInfixOperator returns true if the Token represents an infix operator.
func (tk Token) IsInfixOperator() bool {
	return tk == ADD ||
		tk == SUB ||
		tk == MUL ||
		tk == DIV ||
		tk == REM ||
		tk == EQU ||
		tk == NEQ ||
		tk == LT ||
		tk == MT ||
		tk == LEQ ||
		tk == MEQ ||
		tk == AND ||
		tk == OR
}

// IsOpener returns true if the Token represents an opening bracket of any sort.
func (tk Token) IsOpener() bool {
	return tk == L_PAREN || tk == L_BRACK || tk == L_BRACE
}

// IsCloser returns true if the Token represents an closing bracket of any sort.
func (tk Token) IsCloser() bool {
	return tk == R_PAREN || tk == R_BRACK || tk == R_BRACE
}

// String returns the human readable string representation of the Token.
func (tk Token) String() string {
	switch tk {

	// Reference
	case SOF:
		return "SOF"
	case EOF:
		return "EOF"

		// Redundant
	case SPACE:
		return "SPACE"
	case COMMENT:
		return "COMMENT"

		// Identifiers
	case IDENT:
		return "IDENT"
	case VOID:
		return "VOID"

		// Literals
	case BOOL:
		return "BOOL"
	case NUM:
		return "NUM"
	case STR:
		return "STR"

		// Keywords
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
	case ASSIGN:
		return "ASSIGN"
	case DEFINE:
		return "DEFINE"
	case INTO:
		return "INTO"

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
	case LEQ:
		return "LEQ"
	case MEQ:
		return "MEQ"

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
	case DOT:
		return "DOT"

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

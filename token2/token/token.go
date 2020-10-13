// Token package defines the finite set of token types with utility functions
// for determining if a token is in a predefined subset.
package token

import (
	"strings"
)

// Token represents one of the finite set of symbols used in the parsing process
// to ensure correct syntax, build a model of the program logic (parse tree
// etc), and determine how a Lexeme's value should be parsed and used.
//
// Each Lexeme scanned from source code has a corrisponding Token constant from
// the finite set below; undefined refers to a zero or invalid token. After
// scanning and evaluation the Token constant is the primary means for parsing
// the code into a set of instructions.
type Token string

const (
	UNDEFINED  Token = ``
	SPACE            = `SPACE`      // Whitespace
	COMMENT          = `COMMENT`    // # comment
	NEWLINE          = `NEWLINE`    // \n
	TERMINATOR       = `TERMINATOR` // ;
	TRUE             = `TRUE`       // true
	FALSE            = `FALSE`      // false
	LOOP             = `LOOP`       // loop
	NUMBER           = `NUMBER`     // 1
	STRING           = `STRING`     // "string"
	IDENT            = `IDENT`      // identifier
	SPELL            = `SPELL`      // @spell
	DELIM            = `DELIM`      // ,
	L_PAREN          = `L_PAREN`    // (
	R_PAREN          = `R_PAREN`    // )
	L_SQUARE         = `L_SQUARE`   // [
	R_SQUARE         = `R_SQUARE`   // ]
	L_CURLY          = `L_CURLY`    // {
	R_CURLY          = `R_CURLY`    //	}
	ASSIGN           = `ASSIGN`     // :=
	VOID             = `VOID`       // _
	ADD              = `ADD`        // +
	SUB              = `SUB`        // -
	MUL              = `MUL`        // *
	DIV              = `DIV`        // /
	REM              = `REM`        // %
	AND              = `AND`        // &&
	OR               = `OR`         // ||
	LESS             = `LESS`       // <
	MORE             = `MORE`       // >
	LESS_EQUAL       = `LESS_EQUAL` // <=
	MORE_EQUAL       = `MORE_EQUAL` // >=
	EQUAL            = `EQUAL`      // ==
	NOT_EQUAL        = `NOT_EQUAL`  // !=
)

// IdentifyWord returns the Token represented by the 's'. If 's' does not match
// a keyword then the IDENT Token is returned.
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

// Match returns true if the receiving Token is equal to any of the argument
// tokens.
func (tk Token) Match(tks ...Token) bool {
	for _, o := range tks {
		if tk == o {
			return true
		}
	}
	return false
}

// IsRedundant returns true if the Token is redundant to the parsing process.
func (tk Token) IsRedundant() bool {
	return tk == SPACE || tk == COMMENT
}

// IsTerminator returns true if the Token terminates a statement.
func (tk Token) IsTerminator() bool {
	return tk == TERMINATOR || tk == NEWLINE
}

// IsBool returns true if the Token represents a literal true or false.
func (tk Token) IsBool() bool {
	return tk == TRUE || tk == FALSE
}

// IsLiteral returns true if the Token represents a literal or constant value
// such as a bool, number, or string.
func (tk Token) IsLiteral() bool {
	return tk == NUMBER || tk == STRING || tk.IsBool()
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
	return tk.IsUnaryOperator() || tk.IsBinaryOperator()
}

// IsUnaryOperator returns true if the Token represents a unary operator.
func (tk Token) IsUnaryOperator() bool {
	return false
}

// IsBinaryOperator returns true if the Token represents a binary operator.
func (tk Token) IsBinaryOperator() bool {
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

// IsOpener returns true if the Token represents an opening bracket of any sort.
func (tk Token) IsOpener() bool {
	return tk == L_PAREN || tk == L_SQUARE || tk == L_CURLY
}

// IsCloser returns true if the Token represents an closing bracket of any sort.
func (tk Token) IsCloser() bool {
	return tk == R_PAREN || tk == R_SQUARE || tk == R_CURLY
}

// String returns the human readable string representation of the Token.
func (tk Token) String() string {
	return string(tk)
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

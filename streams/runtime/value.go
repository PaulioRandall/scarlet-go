package runtime

import (
//	"fmt"
//"strconv"

//"github.com/PaulioRandall/scarlet-go/lexeme"
)

// Value represents a value within the executing program, either the value of
// a variable or an intermediate within a statement.
type Value interface {
	Get() interface{}

	String() string
}

type Void struct{}

func (_ Void) Get() interface{} {
	return nil
}

func (_ Void) String() string {
	return "(VOID) _"
}

/*
// NewValue creates a new value from a lexeme.
func NewValue(tk lexeme.Token) Value {

	var (
		k Kind
		v interface{}
	)

	switch tk.Lexeme {
	case lexeme.LEXEME_STRING, lexeme.LEXEME_TEMPLATE:
		k, v = STR, tk.Value

	case lexeme.LEXEME_BOOL:
		k, v = BOOL, (tk.Value == "TRUE")

	case lexeme.LEXEME_INT:
		k, v = INT, parseNum(INT, tk)

	case lexeme.LEXEME_FLOAT:
		k, v = REAL, parseNum(REAL, tk)

	case lexeme.LEXEME_VOID:
		k = VOID

	default:
		panic(newTkErr(tk, "An UNDEFINED token may not be converted to a Value"))
	}

	return Value{
		k: k,
		v: v,
	}
}

// parseNum parses an INT or REAL token value into its Go counterpart.
func parseNum(k Kind, tk lexeme.Token) (v interface{}) {

	var e error

	if k == INT {
		v, e = strconv.ParseInt(tk.Value, 10, 64)
	} else if k == REAL {
		v, e = strconv.ParseFloat(tk.Value, 64)
	} else {
		panic(newTkErr(tk, "SANITY CHECK! Illegal number type, cannot parse"))
	}

	if e != nil {
		panic(newTkError(e, tk, "SANITY CHECK! Could not parse integer token"))
	}

	return
}

// String returns a human readable string representation of the value.
func (v Value) String() string {
	switch v.k {
	case FUNC:
		return v.v.(funcValue).String()
	case STR:
		return "\"" + v.v.(string) + "\""
	default:
		return fmt.Sprintf("%v", v.v)
	}
}
*/

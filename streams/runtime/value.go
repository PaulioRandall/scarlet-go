package runtime

import (
	//	"fmt"
	"strconv"
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

type Bool bool

func (b Bool) Get() interface{} {
	return bool(b)
}

func (b Bool) String() string {
	return "(BOOL) " + strconv.FormatBool(bool(b))
}

type Int int64

func (i Int) Get() interface{} {
	return int64(i)
}

func (i Int) String() string {
	return "(INT) " + strconv.FormatInt(int64(i), 10)
}

type Float float64

func (f Float) Get() interface{} {
	return float64(f)
}

func (f Float) String() string {
	return "(FLOAT) " + strconv.FormatFloat(float64(f), 'f', -1, 64)
}

type String string

func (s String) Get() interface{} {
	return string(s)
}

func (s String) String() string {
	return "(STRING) " + string(s)
}

type Template string

func (t Template) Get() interface{} {
	return string(t)
}

func (t Template) String() string {
	return "(TEMPLATE) " + string(t)
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

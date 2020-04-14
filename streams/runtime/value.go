package runtime

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/lexeme"
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

func valueOf(tk lexeme.Token) Value {

	switch tk.Lexeme {
	case lexeme.LEXEME_VOID:
		return Void{}

	case lexeme.LEXEME_BOOL:
		return Bool(tk.Value == `TRUE`)

	case lexeme.LEXEME_INT:
		return parseInt(tk)

	case lexeme.LEXEME_FLOAT:
		return parseFloat(tk)

	case lexeme.LEXEME_STRING:
		return String(tk.Value)

	case lexeme.LEXEME_TEMPLATE:
		return Template(tk.Value)
	}
	//panic(newTkErr(tk, "An UNDEFINED token may not be converted to a Value"))
	panic(string("TODO: Create Err `Invalid value type " + tk.Lexeme + "`"))
}

// parseInt parses an INT token into an Int value.
func parseInt(tk lexeme.Token) Int {
	i, e := strconv.ParseInt(tk.Value, 10, 64)

	if e != nil {
		//panic(newTkError(e, tk, "SANITY CHECK! Could not parse integer token"))
		panic(string("TODO: Create Err `Could not parse integer`"))
	}

	return Int(i)
}

// parseFloat parses an FLOAT token into an Float value.
func parseFloat(tk lexeme.Token) Float {
	f, e := strconv.ParseFloat(tk.Value, 64)

	if e != nil {
		//panic(newTkError(e, tk, "SANITY CHECK! Could not parse integer token"))
		panic(string("TODO: Create Err `Could not parse float`"))
	}

	return Float(f)
}

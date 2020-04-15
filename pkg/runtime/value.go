package runtime

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/pkg/token"
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

func valueOf(tk token.Token) Value {

	switch tk.Type {
	case token.VOID:
		return Void{}

	case token.BOOL:
		return Bool(tk.Value == `TRUE`)

	case token.INT:
		return parseInt(tk)

	case token.FLOAT:
		return parseFloat(tk)

	case token.STRING:
		return String(tk.Value)

	case token.TEMPLATE:
		return Template(tk.Value)
	}

	panic(err("parseFloat", tk, "Invalid value type (%s)", tk.String()))
}

// parseInt parses an INT token into an Int value.
func parseInt(tk token.Token) Int {
	i, e := strconv.ParseInt(tk.Value, 10, 64)

	if e != nil {
		panic(err("parseInt", tk, "Could not parse integer"))
	}

	return Int(i)
}

// parseFloat parses an FLOAT token into an Float value.
func parseFloat(tk token.Token) Float {
	f, e := strconv.ParseFloat(tk.Value, 64)

	if e != nil {
		panic(err("parseFloat", tk, "Could not parse float"))
	}

	return Float(f)
}

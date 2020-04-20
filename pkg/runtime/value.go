package runtime

import (
	"strconv"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// TODO: Consider renaming this to `Result`

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

type Number float64

func (n Number) Get() interface{} {
	return float64(n)
}

func (n Number) ToInt() int64 {
	return int64(float64(n))
}

func (n Number) String() string {
	return "(NUMBER) " + strconv.FormatFloat(float64(n), 'f', -1, 64)
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

type List []Value

func (l List) Get() interface{} {
	return []Value(l)
}

func (l List) String() string {
	s := "(LIST) " + "{"
	for i, item := range []Value(l) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + "}"
}

type Tuple []Value

func (t Tuple) Get() interface{} {
	return []Value(t)
}

func (t Tuple) String() string {
	s := "(TUPLE) " + "("
	for i, item := range []Value(t) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + ")"
}

type Function st.FuncDef

func (f Function) Get() interface{} {
	return st.FuncDef(f)
}

func (f Function) String() string {

	s := "(FUNCTION) F("

	if f.Input != nil {
		for i, item := range f.Input {
			if i != 0 {
				s += ", "
			}

			s += item.Value
		}
		s += " "
	}

	if f.Output != nil {
		s += "-> "

		for i, item := range f.Output {
			if i != 0 {
				s += ", "
			}

			s += item.Value
		}
	}

	return s + ")"
}

func valueOf(tk token.Token) Value {

	switch tk.Type {
	case token.VOID:
		return Void{}

	case token.BOOL:
		return Bool(tk.Value == `TRUE`)

	case token.NUMBER:
		return parseFloat(tk)

	case token.STRING:
		return String(tk.Value)

	case token.TEMPLATE:
		return Template(tk.Value)
	}

	panic(err("valueOf", tk, "Invalid value type (%s)", tk.String()))
}

// parseFloat parses an NUMBER token into a Number value.
func parseFloat(tk token.Token) Number {
	f, e := strconv.ParseFloat(tk.Value, 64)

	if e != nil {
		panic(err("parseFloat", tk, "Could not parse number"))
	}

	return Number(f)
}

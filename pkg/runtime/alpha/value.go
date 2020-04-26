package alpha

import (
	"strconv"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// TODO: Consider renaming this to `Result`

// value represents a value within the executing program, either the value of
// a variable or an intermediate within a statement.
type value interface {
	Get() interface{}

	String() string
}

type voidLiteral struct{}

func (_ voidLiteral) Get() interface{} {
	return nil
}

func (_ voidLiteral) String() string {
	return "(VOID) _"
}

type boolLiteral bool

func (b boolLiteral) Get() interface{} {
	return bool(b)
}

func (b boolLiteral) String() string {
	return "(BOOL) " + strconv.FormatBool(bool(b))
}

type numberLiteral float64

func (n numberLiteral) Get() interface{} {
	return float64(n)
}

func (n numberLiteral) ToInt() int64 {
	return int64(float64(n))
}

func (n numberLiteral) String() string {
	return "(NUMBER) " + strconv.FormatFloat(float64(n), 'f', -1, 64)
}

type stringLiteral string

func (s stringLiteral) Get() interface{} {
	return string(s)
}

func (s stringLiteral) String() string {
	return "(STRING) " + string(s)
}

type templateLiteral string

func (t templateLiteral) Get() interface{} {
	return string(t)
}

func (t templateLiteral) String() string {
	return "(TEMPLATE) " + string(t)
}

type listLiteral []value

func (l listLiteral) Get() interface{} {
	return []value(l)
}

func (l listLiteral) String() string {
	s := "(LIST) " + "{"
	for i, item := range []value(l) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + "}"
}

type tuple []value

func (t tuple) Get() interface{} {
	return []value(t)
}

func (t tuple) String() string {
	s := "(TUPLE) " + "("
	for i, item := range []value(t) {
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

func valueOf(tk token.Token) value {

	switch tk.Type {
	case token.VOID:
		return voidLiteral{}

	case token.BOOL:
		return boolLiteral(tk.Value == `TRUE`)

	case token.NUMBER:
		return parseFloat(tk)

	case token.STRING:
		return stringLiteral(tk.Value)

	case token.TEMPLATE:
		return templateLiteral(tk.Value)
	}

	panic(err("valueOf", tk, "Invalid value type (%s)", tk.String()))
}

// parseFloat parses an NUMBER token into a numberLiteral.
func parseFloat(tk token.Token) numberLiteral {
	f, e := strconv.ParseFloat(tk.Value, 64)

	if e != nil {
		panic(err("parseFloat", tk, "Could not parse number"))
	}

	return numberLiteral(f)
}

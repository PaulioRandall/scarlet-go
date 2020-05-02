package alpha

import (
	"strconv"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// result represents a value within the executing program, either the value of
// a variable or an intermediate within a statement.
type result interface {
	get() interface{}

	String() string
}

type voidLiteral struct{}

func (_ voidLiteral) get() interface{} {
	return nil
}

func (_ voidLiteral) String() string {
	return "(VOID) _"
}

type boolLiteral bool

func (b boolLiteral) get() interface{} {
	return bool(b)
}

func (b boolLiteral) String() string {
	return "(BOOL) " + strconv.FormatBool(bool(b))
}

type numberLiteral float64

func (n numberLiteral) get() interface{} {
	return float64(n)
}

func (n numberLiteral) ToInt() int64 {
	return int64(float64(n))
}

func (n numberLiteral) String() string {
	return "(NUMBER) " + strconv.FormatFloat(float64(n), 'f', -1, 64)
}

type stringLiteral string

func (s stringLiteral) get() interface{} {
	return string(s)
}

func (s stringLiteral) String() string {
	return "(STRING) " + string(s)
}

type templateLiteral string

func (t templateLiteral) get() interface{} {
	return string(t)
}

func (t templateLiteral) String() string {
	return "(TEMPLATE) " + string(t)
}

type listLiteral []result

func (l listLiteral) get() interface{} {
	return []result(l)
}

func (l listLiteral) String() string {
	s := "(LIST) " + "{"
	for i, item := range []result(l) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + "}"
}

type tuple []result

func (t tuple) get() interface{} {
	return []result(t)
}

func (t tuple) String() string {
	s := "(TUPLE) " + "("
	for i, item := range []result(t) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + ")"
}

type functionLiteral st.FuncDef

func (f functionLiteral) get() interface{} {
	return st.FuncDef(f)
}

func (f functionLiteral) String() string {

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

func valueOf(tk token.Token) result {

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

func parseFloat(tk token.Token) numberLiteral {
	f, e := strconv.ParseFloat(tk.Value, 64)

	if e != nil {
		panic(err("parseFloat", tk, "Could not parse number"))
	}

	return numberLiteral(f)
}

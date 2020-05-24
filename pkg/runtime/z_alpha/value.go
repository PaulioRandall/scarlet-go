package z_alpha

import (
	"strconv"

	"github.com/shopspring/decimal"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
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

type numberLiteral decimal.Decimal

func (n numberLiteral) get() interface{} {
	return decimal.Decimal(n)
}

func (n numberLiteral) ToInt() int64 {
	return decimal.Decimal(n).IntPart()
}

func (n numberLiteral) String() string {
	return "(NUMBER) " + decimal.Decimal(n).String()
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

type functionLiteral FuncDef

func (f functionLiteral) get() interface{} {
	return FuncDef(f)
}

func (f functionLiteral) String() string {

	s := "(FUNCTION) F("

	if f.Inputs != nil {
		for i, item := range f.Inputs {
			if i != 0 {
				s += ", "
			}

			s += item.Value()
		}
	}

	if f.Outputs != nil {
		for i, item := range f.Outputs {
			if i != 0 || f.Inputs != nil {
				s += ", "
			}

			s += "^" + item.Value()
		}
	}

	return s + ")"
}

func valueOf(tk Token) result {

	switch tk.Morpheme() {
	case VOID:
		return voidLiteral{}

	case BOOL:
		return boolLiteral(tk.Value() == `TRUE`)

	case NUMBER:
		return parseFloat(tk)

	case STRING:
		return stringLiteral(tk.Value())

	case TEMPLATE:
		return templateLiteral(tk.Value())
	}

	panic(err("valueOf", tk, "Invalid morpheme %s", tk.Morpheme().String()))
}

func parseFloat(tk Token) numberLiteral {
	d, e := decimal.NewFromString(tk.Value())

	if e != nil {
		panic(err("parseFloat", tk, "Could not parse number"))
	}

	return numberLiteral(d)
}

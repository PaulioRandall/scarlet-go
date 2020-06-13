package cutangle

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
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
	return ""
}

type boolLiteral bool

func (b boolLiteral) get() interface{} {
	return bool(b)
}

func (b boolLiteral) String() string {
	return strconv.FormatBool(bool(b))
}

type numberLiteral decimal.Decimal

func (n numberLiteral) get() interface{} {
	return decimal.Decimal(n)
}

func (n numberLiteral) ToInt() int64 {
	return decimal.Decimal(n).IntPart()
}

func (n numberLiteral) String() string {
	return decimal.Decimal(n).String()
}

type stringLiteral string

func (s stringLiteral) get() interface{} {
	return string(s)
}

func (s stringLiteral) String() string {
	return string(s)
}

type listLiteral []result

func (l listLiteral) get() interface{} {
	return []result(l)
}

func (l listLiteral) String() string {
	s := "{"
	for i, item := range []result(l) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + "}"
}

type tuple []result

func newTuple(rs ...result) tuple {
	return tuple(rs)
}

func (t tuple) get() interface{} {
	return []result(t)
}

func (t tuple) String() string {
	s := "("
	for i, item := range []result(t) {
		if i != 0 {
			s += ","
		}
		s += item.String()
	}
	return s + ")"
}

type funcLiteral FuncDef

func (f funcLiteral) get() interface{} {
	return FuncDef(f)
}

func (f funcLiteral) String() string {

	s := "F("

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

			s += "^" + item.Token().Value()
		}
	}

	return s + ")"
}

type exprFuncLiteral ExprFuncDef

func (e exprFuncLiteral) get() interface{} {
	return ExprFuncDef(e)
}

func (e exprFuncLiteral) String() string {

	s := "E("

	if e.Inputs != nil {
		for i, item := range e.Inputs {
			if i != 0 {
				s += ", "
			}

			s += item.Value()
		}
	}

	return s + ")"
}

func valueOf(tk Token) result {

	switch tk.Type() {
	case TK_VOID:
		return voidLiteral{}

	case TK_BOOL:
		return boolLiteral(tk.Value() == `TRUE`)

	case TK_NUMBER:
		return parseFloat(tk)

	case TK_STRING:
		return stringLiteral(tk.Value())
	}

	err.Panic(
		fmt.Sprintf("SANITY CHECK! Invalid morpheme %s", tk.Type().String()),
		err.At(tk),
	)
	return nil
}

func parseFloat(tk Token) numberLiteral {
	d, e := decimal.NewFromString(tk.Value())

	if e != nil {
		err.Panic("Unparsable number", err.At(tk))
	}

	return numberLiteral(d)
}

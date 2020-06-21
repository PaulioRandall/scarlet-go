package result

import (
	"fmt"
	//"strconv"
	//"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	//. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
	//. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type Void struct{}

func (Void) String() string {
	return "_"
}

type Result struct {
	fmt.Stringer
	typ ResultType
	val interface{}
}

func (r Result) String() string {
	return fmt.Sprintf("%v", r.val)
}

func (r Result) Type() ResultType {
	return r.typ
}

func (r Result) Is(typ ResultType) bool {
	return r.typ == typ
}

func (r Result) Void() (Void, bool) {
	return Void{}, r.Is(RT_VOID)
}

/*
type boolLiteral struct {
	bool
}

func (b boolLiteral) String() string {
	return strconv.FormatBool(b.bool)
}

type numberLiteral decimal.Decimal

func (n numberLiteral) ToInt() int64 {
	return decimal.Decimal(n).IntPart()
}

func (n numberLiteral) String() string {
	return decimal.Decimal(n).String()
}

type stringLiteral string

func (s stringLiteral) String() string {
	return string(s)
}

type listLiteral []result

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
*/

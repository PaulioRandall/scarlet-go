package result

import (
	"fmt"
	//"strconv"
	//"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
	//. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/number"
)

type Void struct{}

func (Void) String() string {
	return "_"
}

type Result struct {
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

func (r Result) Bool() (bool, bool) {

	if r.Is(RT_BOOL) {
		return r.val.(bool), true
	}

	return false, false
}

func (r Result) Num() (number.Number, bool) {

	if r.Is(RT_NUMBER) {
		return r.val.(number.Number), true
	}

	return nil, false
}

func (r Result) Str() (string, bool) {

	if r.Is(RT_STRING) {
		return r.val.(string), true
	}

	return "", false
}

func (r Result) List() ([]Result, bool) {

	if r.Is(RT_LIST) {
		return r.val.([]Result), true
	}

	return nil, false
}

func (r Result) Map() (map[Result]Result, bool) {

	if r.Is(RT_MAP) {
		return r.val.(map[Result]Result), true
	}

	return nil, false
}

func (r Result) Func() (statement.Function, bool) {

	if r.Is(RT_FUNC_DEF) {
		return r.val.(statement.Function), true
	}

	return nil, false
}

func (r Result) ExprFunc() (statement.ExpressionFunction, bool) {

	if r.Is(RT_EXPR_FUNC_DEF) {
		return r.val.(statement.ExpressionFunction), true
	}

	return nil, false
}

func (r Result) Tuple() ([]Result, bool) {

	if r.Is(RT_TUPLE) {
		return r.val.([]Result), true
	}

	return nil, false
}

/*

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

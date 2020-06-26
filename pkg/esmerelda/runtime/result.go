package runtime

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/number"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type VoidResult struct{}

func (VoidResult) String() string {
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

func (r Result) IsNot(typ ResultType) bool {
	return r.typ != typ
}

func (r *Result) Negate() {

	if r.Is(RT_BOOL) {
		r.val = !(r.val.(bool))
		return
	}

	if r.Is(RT_NUMBER) {
		r.val.(number.Number).Neg()
		return
	}

	panic("PROGRAMMERS ERROR! Can only negate booleans or numbers")
}

func (r Result) Equal(o Result) bool {

	if r.IsNot(o.Type()) {
		return false
	}

	switch r.typ {
	case RT_BOOL:
		return r.val.(bool) == o.val.(bool)

	case RT_NUMBER:
		return r.val.(number.Number).Equal(o.val.(number.Number))

	case RT_STRING:
		return r.val.(string) == o.val.(string)
	}

	return false
}

func (r Result) NotEqual(o Result) bool {
	return !r.Equal(o)
}

func (r Result) Void() (VoidResult, bool) {
	return VoidResult{}, r.Is(RT_VOID)
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

func (r Result) Func() (stats.FuncDef, bool) {

	if r.Is(RT_FUNC_DEF) {
		return r.val.(stats.FuncDef), true
	}

	return nil, false
}

func (r Result) ExprFunc() (stats.ExprFunc, bool) {

	if r.Is(RT_EXPR_FUNC_DEF) {
		return r.val.(stats.ExprFunc), true
	}

	return nil, false
}

func (r Result) Tuple() ([]Result, bool) {

	if r.Is(RT_TUPLE) {
		return r.val.([]Result), true
	}

	return nil, false
}

func ResultOf(tk token.Token) Result {

	switch tk.Type() {
	case token.TK_VOID:
		return Result{
			typ: RT_VOID,
			val: VoidResult{},
		}

	case token.TK_BOOL:
		return Result{
			typ: RT_BOOL,
			val: tk.Value() == "true",
		}

	case token.TK_NUMBER:
		return Result{
			typ: RT_NUMBER,
			val: number.New(tk.Value()),
		}

	case token.TK_STRING:
		s := tk.Value()

		return Result{
			typ: RT_STRING,
			val: s[1 : len(s)-1],
		}
	}

	line, sCol := tk.Begin()
	_, eCol := tk.End()

	msg := fmt.Sprintf("Unknown token type '%s', line %d [%d:%d]",
		tk.Type().String(), line+1, sCol, eCol)

	panic("PROGRAMMERS ERROR! " + msg)
}

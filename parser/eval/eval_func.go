package eval

import (
	"github.com/PaulioRandall/scarlet-go/parser/ctx"
)

// NewForFuncCall creates an Expr that invokes a Scarlet function when called.
func NewForFuncCall(getFunc Expr, params []Expr) Expr {

	getInput := evalParams(params)

	return func(parent ctx.Context) (v ctx.Value, e EvalErr) {

		f, e := getFuncProcedure(parent, getFunc)
		if e != nil {
			return
		}

		args, e := getFuncArgs(parent, getInput)
		if e != nil {
			return
		}

		c := parent.Schism()
		v, perr := f(c, args)
		if perr != nil {
			e = NewEvalErr(perr, -1, "TODO")
		}

		return
	}
}

func getFuncArgs(c ctx.Context, getInput Expr) (args []ctx.Value, e EvalErr) {

	asValue, e := getInput(c)
	if e != nil {
		return
	}

	args, err := asValue.ToList()
	if err != nil {
		e = NewEvalErr(err, -1, "TODO")
		return
	}

	return
}

func getFuncProcedure(c ctx.Context, getFunc Expr) (f ctx.Procedure, e EvalErr) {

	asValue, e := getFunc(c)
	if e != nil {
		return
	}

	f, err := asValue.ToFunc()
	if err != nil {
		e = NewEvalErr(err, -1, "TODO")
		return
	}

	return
}

package alpha

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/shopspring/decimal"
)

type spell func(ctx *alphaContext, call SpellCall, args []arg) result
type arg struct {
	tk Token
	r  result
}

func evalSpellCall(ctx *alphaContext, call SpellCall) result {

	sp := findSpell(call.ID)
	var args []arg

	for _, v := range call.Inputs {

		a := arg{
			r:  evalExpression(ctx, v),
			tk: v.Token(),
		}

		args = append(args, a)
	}

	return sp(ctx, call, args)
}

func findSpell(id Token) spell {

	switch strings.ToLower(id.Value()) {
	case `print`:
		return spellPrint
	case `println`:
		return spellPrintln
	case `inc`:
		return spellIncrement
	case `dec`:
		return spellDecrement
	}

	err.Panic("Not a known spell", err.At(id))
	return nil
}

func argToNumber(a arg) (decimal.Decimal, string) {
	n, ok := a.r.(numberLiteral)

	if !ok {
		return decimal.Decimal{}, "Not a number"
	}

	return decimal.Decimal(n), ""
}

func spellPrint(ctx *alphaContext, call SpellCall, args []arg) result {

	for _, v := range args {
		print(v.r.String())
	}

	return voidLiteral{}
}

func spellPrintln(ctx *alphaContext, call SpellCall, args []arg) result {

	for _, v := range args {
		print(v.r.String())
	}

	println()
	return voidLiteral{}
}

func spellIncrement(ctx *alphaContext, call SpellCall, args []arg) result {

	newErr := func(msg string) result {
		return newTuple(
			voidLiteral{},
			stringLiteral(msg),
		)
	}

	if len(args) != 1 {
		return newErr("Wrong number of arguments")
	}

	a := args[0]

	if a.tk.Morpheme() != IDENTIFIER {
		return newErr("Not an identifier")
	}

	d, e := argToNumber(a)
	if e != "" {
		return newErr(e)
	}

	one := decimal.NewFromInt(1)
	d = d.Add(one)
	n := numberLiteral(d)

	ctx.Set(a.tk, n)
	return newTuple(n, stringLiteral(""))
}

func spellDecrement(ctx *alphaContext, call SpellCall, args []arg) result {

	newErr := func(msg string) result {
		return newTuple(
			voidLiteral{},
			stringLiteral(msg),
		)
	}

	if len(args) != 1 {
		return newErr("Wrong number of arguments")
	}

	a := args[0]

	if a.tk.Morpheme() != IDENTIFIER {
		return newErr("Not an identifier")
	}

	d, e := argToNumber(a)
	if e != "" {
		return newErr(e)
	}

	one := decimal.NewFromInt(1)
	d = d.Sub(one)
	n := numberLiteral(d)

	ctx.Set(a.tk, n)
	return newTuple(n, stringLiteral(""))
}

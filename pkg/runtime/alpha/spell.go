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
	case `p`:
		return spellPrint
	case `pl`:
		return spellPrintln
	case `inc`:
		return spellIncrement
	case `dec`:
		return spellDecrement
	}

	err.Panic("Not a known spell", err.At(id))
	return nil
}

func checkArgsLen(call SpellCall, args []arg, expSize int) {
	if len(args) != expSize {
		err.Panic("Wrong number of arguments", err.At(call.Token()))
	}
}

func checkIsIdentifier(a arg) {
	if a.tk.Morpheme() != IDENTIFIER {
		err.Panic("Not an identifier", err.At(a.tk))
	}
}

func argToNumber(a arg) decimal.Decimal {
	n, ok := a.r.(numberLiteral)

	if !ok {
		err.Panic("Not a number", err.At(a.tk))
	}

	return decimal.Decimal(n)
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

	checkArgsLen(call, args, 1)
	a := args[0]
	checkIsIdentifier(a)

	d := argToNumber(a)
	one := decimal.NewFromInt(1)
	d = d.Add(one)
	n := numberLiteral(d)

	ctx.Set(a.tk, n)
	return n
}

func spellDecrement(ctx *alphaContext, call SpellCall, args []arg) result {

	checkArgsLen(call, args, 1)
	a := args[0]
	checkIsIdentifier(a)

	d := argToNumber(a)
	one := decimal.NewFromInt(1)
	d = d.Sub(one)
	n := numberLiteral(d)

	ctx.Set(a.tk, n)
	return n
}

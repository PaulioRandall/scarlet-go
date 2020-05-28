package alpha

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalSpellCall(ctx *alphaContext, call SpellCall) result {
	sp := findSpell(call.ID)
	args := evalExpressions(ctx, call.Inputs)
	return sp(ctx, args)
}

type spell func(ctx *alphaContext, args []result) result

func findSpell(id Token) spell {

	switch strings.ToUpper(id.Value()) {
	case `P`:
		return spellPrint
	case `PL`:
		return spellPrintln
	}

	err.Panic("Not a known spell", err.At(id))
	return nil
}

func spellPrint(ctx *alphaContext, args []result) result {

	for _, v := range args {
		print(v.String())
	}

	return voidLiteral{}
}

func spellPrintln(ctx *alphaContext, args []result) result {

	for _, v := range args {
		print(v.String())
	}

	println()
	return voidLiteral{}
}

package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/shopspring/decimal"
)

func exeBlock(ctx *alphaContext, b Block) {
	exeStatements(ctx, b.Stats)
}

func exeStatements(ctx *alphaContext, ss []Statement) {
	for _, s := range ss {
		exeStatement(ctx, s)
	}
}

func exeStatement(ctx *alphaContext, s Statement) {
	switch v := s.(type) {
	case Increment:
		exeIncrement(ctx, v)

	case Assignment:
		exeAssignment(ctx, v)

	case Match:
		exeMatch(ctx, v)

	case Guard:
		exeGuard(ctx, v)

	case Loop:
		exeLoop(ctx, v)

	case ForEach:
		exeForEach(ctx, v)

	case Expression:
		_ = evalExpression(ctx, v)

	default:
		err.Panic("Unknown statement", err.At(s.Token()))
	}
}

func exeIncrement(ctx *alphaContext, inc Increment) {

	n := evalNumber(ctx, inc.ID)
	one := decimal.NewFromInt(1)

	if inc.Direction.Morpheme() == INCREMENT {
		n = n.Add(one)
	} else {
		n = n.Sub(one)
	}

	num := numberLiteral(n)
	ctx.Set(inc.ID.Token(), num)
}

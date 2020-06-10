package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
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
	case Assignment:
		exeAssignment(ctx, v)

	case When:
		exeWhen(ctx, v)

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

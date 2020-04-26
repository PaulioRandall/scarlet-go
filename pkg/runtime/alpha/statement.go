package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func ExeBlock(ctx *alphaContext, b st.Block) {
	ExeStatements(ctx, b.Stats)
}

func ExeStatements(ctx *alphaContext, ss []st.Statement) {
	for _, s := range ss {
		ExeStatement(ctx, s)
	}
}

func ExeStatement(ctx *alphaContext, s st.Statement) {
	switch v := s.(type) {
	case st.Assignment:
		ExeAssignment(ctx, v)

	case st.Match:
		ExeMatch(ctx, v)

	case st.Guard:
		ExeGuard(ctx, v)

	case st.Expression:
		_ = EvalExpression(ctx, v)

	default:
		panic(err("ExeStatement", s.Token(), "Unknown statement type"))
	}
}

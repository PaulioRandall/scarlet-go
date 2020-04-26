package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func exeBlock(ctx *alphaContext, b st.Block) {
	exeStatements(ctx, b.Stats)
}

func exeStatements(ctx *alphaContext, ss []st.Statement) {
	for _, s := range ss {
		exeStatement(ctx, s)
	}
}

func exeStatement(ctx *alphaContext, s st.Statement) {
	switch v := s.(type) {
	case st.Assignment:
		exeAssignment(ctx, v)

	case st.Match:
		exeMatch(ctx, v)

	case st.Guard:
		exeGuard(ctx, v)

	case st.Expression:
		_ = evalExpression(ctx, v)

	default:
		panic(err("exeStatement", s.Token(), "Unknown statement type"))
	}
}

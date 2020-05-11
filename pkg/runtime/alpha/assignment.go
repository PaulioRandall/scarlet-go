package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func exeAssignment(ctx *alphaContext, a st.Assignment) {

	vs := evalExpressions(ctx, a.Exprs)
	checkAssignments(a.Targets, vs, a.Assign)

	for i, id := range a.Targets {
		if a.Fixed {
			ctx.SetFixed(id, vs[i])
		} else {
			ctx.Set(id, vs[i])
		}
	}
}

func checkAssignments(ids []token.Token, vals []result, operator token.Token) {

	a, b := len(ids), len(vals)

	if a > b {
		panic(err("ExeStatement", operator,
			"Missing expression results to populate variables... have %d, want %d",
			a, b,
		))
	}

	if a < b {
		panic(err("ExeStatement", operator,
			"Too many expression results to populate variables... have %d, want %d",
			a, b,
		))
	}
}

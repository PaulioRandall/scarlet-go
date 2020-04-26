package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func exeAssignment(ctx *alphaContext, a st.Assignment) {

	values := evalExpressions(ctx, a.Exprs)
	checkAssignments(a.IDs, values, a.Assign)

	for i, id := range a.IDs {
		if a.Fixed {
			ctx.SetFixed(id, values[i])
		} else {
			ctx.Set(id, values[i])
		}
	}
}

func checkAssignments(ids []token.Token, vals []Value, operator token.Token) {

	a, b := len(ids), len(vals)

	if a > b {
		panic(err("ExeStatement", operator,
			"Missing expression values to populate variables... have %d, want %d",
			a, b,
		))
	}

	if a < b {
		panic(err("ExeStatement", operator,
			"Too many expression values to populate variables... have %d, want %d",
			a, b,
		))
	}
}

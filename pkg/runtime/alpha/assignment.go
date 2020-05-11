package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func exeAssignment(ctx *alphaContext, a st.Assignment) {

	vs := evalExpressions(ctx, a.Exprs)
	checkAssignTargets(a.Targets, vs, a.Assign)

	for i, at := range a.Targets {
		if at.Index == nil {
			assignVar(ctx, at.ID, a.Fixed, vs[i])
		} else {
			assignListItem(ctx, at.ID, at.Index, vs[i])
		}
	}
}

func assignVar(ctx *alphaContext, id token.Token, fixed bool, v result) {
	if fixed {
		ctx.SetFixed(id, v)
	} else {
		ctx.Set(id, v)
	}
}

func assignListItem(ctx *alphaContext, id token.Token, index st.Expression, v result) {
	panic(err("assignListItem", id, "Not yet implemented"))
}

func checkAssignTargets(ats []st.AssignTarget, vals []result, operator token.Token) {

	a, b := len(ats), len(vals)

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

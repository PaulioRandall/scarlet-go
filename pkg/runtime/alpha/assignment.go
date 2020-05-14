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

	n := evalNumber(ctx, index)
	i := int64(n)

	listVal := ctx.GetNonFixed(id)
	if listVal == nil {
		panic(err("assignListItem", id, "List variable does not exist"))
	}

	list, ok := listVal.(listLiteral)
	if !ok {
		panic(err("assignListItem", id, "Variable is not a list"))
	}

	items := []result(list)
	size := int64(len(items))

	if i < 0 || i >= size {
		panic(err("assignListItem", index.Token(),
			"Index out of range, accessing %s[%d] from %s[0:%d]",
			index.Token().Value, i, index.Token().Value, size))
	}

	items[i] = v
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

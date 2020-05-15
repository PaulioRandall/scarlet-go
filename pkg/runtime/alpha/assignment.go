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

	list := getListLiteral(ctx, id)
	items := []result(list)

	i := int64(evalNumber(ctx, index))
	checkIndexInRange(items, i, index)

	items = updateListItems(ctx, items, i, v)
	list = listLiteral(items)

	ctx.Set(id, list)
}

func getListLiteral(ctx *alphaContext, id token.Token) listLiteral {

	listVal := ctx.GetNonFixed(id)
	if listVal == nil {
		panic(err("assignListItem", id, "List variable does not exist"))
	}

	list, ok := listVal.(listLiteral)
	if !ok {
		panic(err("assignListItem", id, "Variable is not a list"))
	}

	return list
}

func checkIndexInRange(items []result, i int64, index st.Expression) {

	size := int64(len(items))

	if i < 0 || i >= size {
		panic(err("assignListItem", index.Token(),
			"Index out of range, accessing %s[%d] from %s[0:%d]",
			index.Token().Value, i, index.Token().Value, size))
	}
}

func updateListItems(ctx *alphaContext, items []result, i int64, v result) []result {

	switch _, ok := v.(voidLiteral); {
	case !ok:
		items[i] = v

	case i == 0:
		items = items[1:]

	case i > int64(len(items)):
		items = items[0 : len(items)-1]

	default:
		items = append(items[:i], items[i+1:]...)
	}

	return items
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

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
	size := int64(len(items))

	if void, ok := v.(voidLiteral); ok {
		i := evalIndexExpr(ctx, size, index, true)
		checkListIndexRange(i, size, id, index.Token())
		items = deleteListItem(ctx, items, i, void)

	} else {
		i := evalIndexExpr(ctx, size, index, false)
		items = updateListItems(ctx, items, i, v)
	}

	list = listLiteral(items)
	ctx.Set(id, list)
}

func deleteListItem(ctx *alphaContext, items []result, i int64, v result) []result {
	switch {
	case i == 0:
		items = items[1:]

	case i > int64(len(items)):
		items = items[0 : len(items)-1]

	default:
		items = append(items[:i], items[i+1:]...)
	}

	return items
}

func getListLiteral(ctx *alphaContext, id token.Token) listLiteral {

	listVal := ctx.GetLocal(id.Value)
	if listVal == nil {
		panic(err("assignListItem", id, "List variable is fixed or does not exist"))
	}

	list, ok := listVal.(listLiteral)
	if !ok {
		panic(err("assignListItem", id, "Variable is not a list"))
	}

	return list
}

func updateListItems(ctx *alphaContext, items []result, i int64, v result) []result {

	switch i {
	case -1:
		items = append([]result{v}, items...)

	case int64(len(items)):
		items = append(items, v)

	default:
		items[i] = v
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

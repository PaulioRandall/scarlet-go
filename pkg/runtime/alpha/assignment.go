package alpha

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func exeAssignment(ctx *alphaContext, a Assignment) {

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

func assignVar(ctx *alphaContext, id Token, fixed bool, v result) {
	if fixed {
		ctx.SetFixed(id, v)
	} else {
		ctx.Set(id, v)
	}
}

func assignListItem(ctx *alphaContext, id Token, index Expression, v result) {

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

func getListLiteral(ctx *alphaContext, id Token) listLiteral {

	listVal := ctx.GetLocal(id.Value())
	if listVal == nil {
		err.Panic("List variable is fixed or does not exist", err.At(id))
	}

	list, ok := listVal.(listLiteral)
	if !ok {
		err.Panic("Variable is not a list", err.At(id))
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

func checkAssignTargets(ats []AssignTarget, vals []result, operator Token) {

	a, b := len(ats), len(vals)

	if a > b {
		err.Panic(
			fmt.Sprintf("Too many identifiers on left side... have %d, want %d", a, b),
			err.At(operator),
		)
	}

	if a < b {
		err.Panic(
			fmt.Sprintf("Too many expressions on right side... have %d, want %d", a, b),
			err.At(operator),
		)
	}
}

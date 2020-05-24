package alpha

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalList(ctx *alphaContext, list List) result {
	return listLiteral(evalExpressions(ctx, list.Exprs))
}

func evalListAccess(ctx *alphaContext, la ListAccess) result {

	list := getList(ctx, la.ID)
	items := []result(list)

	size := int64(len(items))
	i := evalIndexExpr(ctx, size, la.Index, true)

	checkListIndexRange(i, size, la.ID.Token(), la.Index.Token())
	return items[i]
}

func getList(ctx *alphaContext, id Identifier) listLiteral {

	v := evalIdentifier(ctx, id)

	if list, ok := v.(listLiteral); ok {
		return list
	}

	panic(err("EvalListAccess", id.Token(), "Can't get item of a non-list"))
}

func evalIndexExpr(ctx *alphaContext, listSize int64, expr Expression, inclusiveRef bool) int64 {

	if ref, ok := expr.(ListItemRef); ok {
		return resolveListRef(listSize, ref.Token(), inclusiveRef)
	}

	n := evalExpression(ctx, expr)

	if i, ok := n.(numberLiteral); ok {
		return i.ToInt()
	}

	panic(err("EvalListAccess", expr.Token(), "Expected number as result"))
}

func resolveListRef(listSize int64, ref Token, inclusiveRef bool) int64 {

	var min, max int64

	if inclusiveRef {
		min, max = 0, listSize-1
	} else {
		min, max = -1, listSize
	}

	switch ref.Morpheme() {
	case LIST_START:
		return int64(min)

	case LIST_END:
		return int64(max)
	}

	panic(err("getListIndex", ref, "Unknown list reference type"))
}

func checkListIndexRange(i, size int64, id, index Token) {

	if i < 0 {
		panic(err("EvalListAccess", id, "Index out of range, %d < 0", i))
	}

	if i >= size {
		panic(err("EvalListAccess", index, "Index out of range, len(list) < %d", i))
	}
}

package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalList(ctx *alphaContext, list st.List) result {
	return listLiteral(evalExpressions(ctx, list.Exprs))
}

func evalListAccess(ctx *alphaContext, la st.ListAccess) result {

	list := getList(ctx, la.ID)
	items := []result(list)

	size := int64(len(items))
	i := evalIndexExpr(ctx, size, la.Index, true)

	checkListIndexRange(i, size, la.ID.Token(), la.Index.Token())
	return items[i]
}

func getList(ctx *alphaContext, id st.Identifier) listLiteral {

	v := evalIdentifier(ctx, id)

	if list, ok := v.(listLiteral); ok {
		return list
	}

	panic(err("EvalListAccess", id.Token(), "Can't get item of a non-list"))
}

func evalIndexExpr(ctx *alphaContext, listSize int64, expr st.Expression, inclusiveRef bool) int64 {

	if ref, ok := expr.(st.ListItemRef); ok {
		return resolveListRef(listSize, ref.Token(), inclusiveRef)
	}

	n := evalExpression(ctx, expr)

	if i, ok := n.(numberLiteral); ok {
		return i.ToInt()
	}

	panic(err("EvalListAccess", expr.Token(), "Expected number as result"))
}

func resolveListRef(listSize int64, ref token.Token, inclusiveRef bool) int64 {

	var min, max int64

	if inclusiveRef {
		min, max = 0, listSize-1
	} else {
		min, max = -1, listSize
	}

	switch ref.Type {
	case token.PREPEND:
		return int64(min)

	case token.APPEND:
		return int64(max)
	}

	panic(err("getListIndex", ref, "Unknown list reference type"))
}

func resolveListSetterRef(listSize int64, ref token.Token) int64 {

	switch ref.Type {
	case token.PREPEND:
		return int64(0)

	case token.APPEND:
		return int64(listSize - 1)
	}

	panic(err("getListIndex", ref, "Unknown list reference type"))
}

func checkListIndexRange(i, size int64, id, index token.Token) {

	if i < 0 {
		panic(err("EvalListAccess", id, "Index out of range, %d < 0", i))
	}

	if i >= size {
		panic(err("EvalListAccess", index, "Index out of range, len(list) < %d", i))
	}
}

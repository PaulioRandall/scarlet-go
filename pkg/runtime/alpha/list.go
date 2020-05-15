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
	i := getListIndex(ctx, size, la)

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

func getListIndex(ctx *alphaContext, listSize int64, la st.ListAccess) int64 {

	if ref, ok := la.Index.(st.ListItemRef); ok {
		return resolveListItemRef(listSize, ref.Token())
	}

	n := evalExpression(ctx, la.Index)

	if i, ok := n.(numberLiteral); ok {
		return i.ToInt()
	}

	panic(err("EvalListAccess", la.Index.Token(), "Expected number as result"))
}

func resolveListItemRef(listSize int64, ref token.Token) int64 {

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

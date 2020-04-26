package alpha

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func EvalList(ctx *alphaContext, list st.List) Value {
	return List(EvalExpressions(ctx, list.Exprs))
}

func EvalListAccess(ctx *alphaContext, la st.ListAccess) Value {

	v := EvalIdentifier(ctx, la.ID)
	list, ok := v.(List)

	if !ok {
		panic(err("EvalListAccess", la.ID.Token(), "Can't get item of a non-list"))
	}

	n := EvalExpression(ctx, la.Index)
	index, ok := n.(Number)

	if !ok {
		panic(err("EvalListAccess", la.Index.Token(), "Expected number as result"))
	}

	i := index.ToInt()
	if i < 0 {
		panic(err("EvalListAccess", la.ID.Token(), "Index must be greater than zero"))
	}

	items := []Value(list)
	if i >= int64(len(items)) {
		panic(err("EvalListAccess", la.Index.Token(), "Index out of range"))
	}

	return items[i]
}

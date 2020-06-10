package alpha

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
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
	list, ok := v.(listLiteral)

	if !ok {
		err.Panic("Not a list", err.At(id.Token()))
	}

	return list
}

func evalIndexExpr(ctx *alphaContext, listSize int64, expr Expression, inclusiveRef bool) int64 {

	if ref, ok := expr.(ListItemRef); ok {
		return resolveListRef(listSize, ref.Token(), inclusiveRef)
	}

	n := evalExpression(ctx, expr)
	i, ok := n.(numberLiteral)

	if !ok {
		err.Panic("Need a number", err.At(expr.Token()))
	}

	return i.ToInt()
}

func resolveListRef(listSize int64, ref Token, inclusiveRef bool) int64 {

	var min, max int64

	if inclusiveRef {
		min, max = 0, listSize-1
	} else {
		min, max = -1, listSize
	}

	switch ref.Type() {
	case TK_LIST_START:
		return int64(min)

	case TK_LIST_END:
		return int64(max)
	}

	err.Panic("Unknown list reference type", err.At(ref))
	return 0
}

func checkListIndexRange(i, size int64, id, index Token) {

	if i < 0 {
		m := fmt.Sprintf("Index out of range, (index) %d < 0 (min)", i)
		err.Panic(m, err.At(id))
	}

	if i >= size {
		m := fmt.Sprintf("Index out of range, (index) %d >= %d (size)", i, size)
		err.Panic(m, err.At(index))
	}
}

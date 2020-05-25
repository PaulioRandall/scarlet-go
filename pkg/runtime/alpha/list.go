package alpha

import (
	"fmt"

	errr "github.com/PaulioRandall/scarlet-go/pkg/err"
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

	errr.Panic("Not a list", errr.At(id.Token()))
	return nil
}

func evalIndexExpr(ctx *alphaContext, listSize int64, expr Expression, inclusiveRef bool) int64 {

	if ref, ok := expr.(ListItemRef); ok {
		return resolveListRef(listSize, ref.Token(), inclusiveRef)
	}

	n := evalExpression(ctx, expr)

	if i, ok := n.(numberLiteral); ok {
		return i.ToInt()
	}

	errr.Panic("Need a number", errr.At(expr.Token()))
	return 0
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

	errr.Panic("Unknown list reference type", errr.At(ref))
	return 0
}

func checkListIndexRange(i, size int64, id, index Token) {

	if i < 0 {
		m := fmt.Sprintf("Index out of range, (index) %d < 0 (min)", i)
		errr.Panic(m, errr.At(id))
	}

	if i >= size {
		m := fmt.Sprintf("Index out of range, (index) %d >= %d (size)", i, size)
		errr.Panic(m, errr.At(index))
	}
}

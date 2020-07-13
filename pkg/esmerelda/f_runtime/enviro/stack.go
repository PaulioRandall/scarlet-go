package enviro

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/types"
)

type Stack struct {
	size int
	top  *node
}

type node struct {
	data types.Value
	next *node
}

func (stk *Stack) Push(data types.Value) {

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *Stack) Pop() types.Value {

	if stk.size == 0 {
		perror.Panic("Nothing to pop, check stack first")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

func (stk *Stack) Descend(f func(data types.Value)) {
	for n := stk.top; n != nil; n = n.next {
		f(n.data)
	}
}

func (stk *Stack) ToArray() []types.Value {

	vs := []types.Value{}

	stk.Descend(func(v types.Value) {
		vs = append(vs, v)
	})

	return vs
}

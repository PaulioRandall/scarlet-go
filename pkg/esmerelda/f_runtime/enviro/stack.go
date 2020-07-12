package enviro

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
)

type Stack struct {
	size int
	top  *node
}

type node struct {
	data result.Result
	next *node
}

func (stk *Stack) Push(data result.Result) {

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *Stack) Pop() result.Result {

	if stk.size == 0 {
		perror.Panic("Nothing to pop, check stack first")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

func (stk *Stack) Descend(f func(data result.Result)) {
	for n := stk.top; n != nil; n = n.next {
		f(n.data)
	}
}

func (stk *Stack) ToArray() []result.Result {

	rs := []result.Result{}

	stk.Descend(func(r result.Result) {
		rs = append(rs, r)
	})

	return rs
}

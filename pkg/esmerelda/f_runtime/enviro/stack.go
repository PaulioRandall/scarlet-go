package enviro

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
)

type Stack struct {
	size int
	top  *node
}

type node struct {
	data Result
	next *node
}

func (stk *Stack) Push(data Result) {

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *Stack) Pop() Result {

	if stk.size == 0 {
		perror.Panic("Nothing to pop, check stack first")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

func (stk *Stack) Descend(f func(data Result)) {
	for n := stk.top; n != nil; n = n.next {
		f(n.data)
	}
}

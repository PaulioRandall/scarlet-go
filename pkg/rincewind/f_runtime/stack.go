package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

type stack struct {
	size int
	top  *node
}

type node struct {
	data result
	next *node
}

func (stk *stack) push(data result) {

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *stack) pop() result {

	if stk.size == 0 {
		perror.Panic("Nothing to pop, check stack first")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

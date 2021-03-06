package value

type Stack struct {
	size int
	top  *node
}

type node struct {
	data Value
	next *node
}

func (stk *Stack) Size() int {
	return stk.size
}

func (stk *Stack) Empty() bool {
	return stk.size == 0
}

func (stk *Stack) Top() Value {
	if stk.size == 0 {
		panic("Nothing on top, check stack first")
	}
	return stk.top.data
}

func (stk *Stack) Push(data Value) {
	stk.top = &node{
		data: data,
		next: stk.top,
	}
	stk.size++
}

func (stk *Stack) Pop() Value {
	if stk.size == 0 {
		panic("Nothing to pop, check stack first")
	}
	data := stk.top.data
	stk.top = stk.top.next
	stk.size--
	return data
}

func (stk *Stack) Descend(f func(data Value)) {
	for n := stk.top; n != nil; n = n.next {
		f(n.data)
	}
}

func (stk *Stack) ToSlice() []Value {
	vs := []Value{}
	for n := stk.top; n != nil; n = n.next {
		vs = append(vs, n.data)
	}
	return vs
}

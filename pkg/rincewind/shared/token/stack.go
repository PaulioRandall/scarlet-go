package token

type Stack struct {
	top  *node
	size int
}

type node struct {
	data Token
	next *node
}

func (stk *Stack) Empty() bool {
	return stk.top == nil
}

func (stk *Stack) Peek() Token {

	if stk.top == nil {
		panic("token.Stack.Peek: stack is empty")
	}

	return stk.top.data
}

func (stk *Stack) Push(data Token) {

	if data == nil {
		panic("token.Stack.Push: nil data is not allowed")
	}

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *Stack) Pop() Token {

	if stk.size == 0 {
		panic("token.Stack.Pop: stack is empty")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

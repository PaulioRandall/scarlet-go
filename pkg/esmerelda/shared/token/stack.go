package token

type Stack struct {
	top  *stkNode
	size int
}

type stkNode struct {
	data Token
	next *stkNode
}

func (stk *Stack) Top() Token {

	if stk.top == nil {
		return nil
	}

	return stk.top.data
}

func (stk *Stack) Push(data Token) bool {

	if data == nil {
		return false
	}

	stk.top = &stkNode{
		data: data,
		next: stk.top,
	}

	stk.size++
	return true
}

func (stk *Stack) Pop() Token {

	if stk.size == 0 {
		return nil
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

func (stk *Stack) Descend(f func(data Token)) {
	for n := stk.top; n != nil; n = n.next {
		f(n.data)
	}
}
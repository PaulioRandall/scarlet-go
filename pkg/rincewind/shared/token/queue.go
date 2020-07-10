package token

type Queue struct {
	front *queNode
	back  *queNode
}

type queNode struct {
	data Token
	next *queNode
}

func (q *Queue) Empty() bool {
	return q.front == nil
}

func (q *Queue) Put(data Token) {

	n := &queNode{
		data: data,
	}

	if q.front == nil {
		q.front = n
	} else {
		q.back.next = n
	}

	q.back = n
}

func (q *Queue) Take() Token {

	if q.front == nil {
		return nil
	}

	data := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.back = nil
	}

	return data
}

func (q *Queue) Descend(f func(Token)) {
	for n := q.front; n != nil; n = n.next {
		f(n.data)
	}
}

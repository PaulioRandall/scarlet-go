package check

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type queue struct {
	front *node
	back  *node
}

type node struct {
	data Token
	next *node
}

func (q *queue) empty() bool {
	return q.front == nil
}

func (q *queue) put(tk Token) {

	n := &node{
		data: tk,
	}

	if q.front == nil {
		q.front, q.back = n, n
		return
	}

	q.back.next = n
	q.back = n
}

func (q *queue) take() Token {

	if q.front == nil {
		return nil
	}

	tk := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.back = nil
	}

	return tk
}

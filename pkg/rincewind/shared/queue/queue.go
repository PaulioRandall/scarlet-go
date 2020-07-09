package queue

type Queue struct {
	front *node
	back  *node
}

type node struct {
	data interface{}
	next *node
}

func (q *Queue) Empty() bool {
	return q.front == nil
}

func (q *Queue) Put(data interface{}) {

	n := &node{
		data: data,
	}

	if q.front == nil {
		q.front, q.back = n, n
		return
	}

	q.back.next = n
	q.back = n
}

func (q *Queue) Take() interface{} {

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

func (q *Queue) Descend(f func(interface{})) {
	for n := q.front; n != nil; n = n.next {
		f(n.data)
	}
}

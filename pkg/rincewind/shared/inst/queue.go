package inst

type Queue struct {
	front *instNode
	back  *instNode
}

type instNode struct {
	data Instruction
	next *instNode
}

func (q *Queue) Empty() bool {
	return q.front == nil
}

func (q *Queue) Put(data Instruction) {

	n := &instNode{
		data: data,
	}

	if q.front == nil {
		q.front = n
	} else {
		q.back.next = n
	}

	q.back = n
}

func (q *Queue) Take() Instruction {

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

func (q *Queue) Descend(f func(Instruction)) {
	for n := q.front; n != nil; n = n.next {
		f(n.data)
	}
}

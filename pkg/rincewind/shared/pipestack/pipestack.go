package pipestack

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/perror"
)

type DataPipe interface {
	Next() interface{}
}

type Matcher interface {
	Match(data, other interface{}) bool
	Expect(data, other interface{}) error
}

type PipeStack struct {
	ds   DataPipe
	mtc  Matcher
	buff interface{}
	top  *node
	size int
}

type node struct {
	data interface{}
	next *node
}

func NewPipeStack(ds DataPipe, mtc Matcher) *PipeStack {

	stk := &PipeStack{
		ds:  ds,
		mtc: mtc,
	}

	stk.buff = stk.ds.Next()
	return stk
}

func (stk *PipeStack) Empty() bool {
	return stk.EmptyBuff() && stk.EmptyStack()
}

func (stk *PipeStack) EmptyBuff() bool {
	return stk.buff == nil
}

func (stk *PipeStack) EmptyStack() bool {
	return stk.size == 0
}

func (stk *PipeStack) Next() interface{} {

	if stk.Empty() {
		perror.Panic("No data remaining, check first")
	}

	data := stk.buff
	stk.buff = stk.ds.Next()
	return data
}

func (stk *PipeStack) PeekBuff() interface{} {
	return stk.buff
}

func (stk *PipeStack) PeekTop() interface{} {

	if stk.size == 0 {
		return nil
	}

	return stk.top.data
}

func (stk *PipeStack) push(data interface{}) {

	if data == nil {
		perror.Panic("PipeStack does not allow nil data")
	}

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *PipeStack) Pop() interface{} {

	if stk.size == 0 {
		perror.Panic("Nothing to pop, check stack first")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

func (stk *PipeStack) MatchBuff(other interface{}) bool {
	return stk.mtc.Match(stk.buff, other)
}

func (stk *PipeStack) AcceptPush(other interface{}) bool {

	if stk.MatchBuff(other) {
		stk.push(stk.Next())
		return true
	}

	return false
}

func (stk *PipeStack) ExpectPush(other interface{}) error {

	e := stk.mtc.Expect(stk.buff, other)
	if e != nil {
		return e
	}

	stk.push(stk.Next())
	return nil
}

func (stk *PipeStack) MatchTop(other interface{}) bool {
	return stk.mtc.Match(stk.PeekTop(), other)
}

func (stk *PipeStack) AcceptPop(other interface{}) interface{} {

	if stk.MatchTop(other) {
		return stk.Pop()
	}

	return nil
}

func (stk *PipeStack) ExpectPop(other interface{}) (interface{}, error) {

	if stk.EmptyStack() {
		return nil, stk.mtc.Expect(nil, other)
	}

	e := stk.mtc.Expect(stk.top.data, other)
	if e != nil {
		return nil, e
	}

	return stk.Pop(), nil
}

func (stk *PipeStack) DescendStack(f func(data interface{})) {
	for n := stk.top; n != nil; n = n.next {
		f(n.data)
	}
}

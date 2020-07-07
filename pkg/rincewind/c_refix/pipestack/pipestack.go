package pipestack

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

type DataPipe interface {
	Next() interface{}
}

type Matcher interface {
	Match(data, other interface{}) bool
	Expect(data, other interface{}) error
}

type Stack struct {
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

func New(ds DataPipe, mtc Matcher) *Stack {

	stk := &Stack{
		ds:  ds,
		mtc: mtc,
	}

	stk.buff = stk.ds.Next()
	return stk
}

func (stk *Stack) Empty() bool {
	return stk.buff == nil && stk.size == 0
}

func (stk *Stack) EmptyStk() bool {
	return stk.size == 0
}

func (stk *Stack) Next() interface{} {

	if stk.Empty() {
		perror.Panic("No data remaining, check first")
	}

	data := stk.buff
	stk.buff = stk.ds.Next()
	return data
}

func (stk *Stack) PeekNext() interface{} {
	return stk.buff
}

func (stk *Stack) PeekStk() interface{} {

	if stk.size == 0 {
		return nil
	}

	return stk.top.data
}

func (stk *Stack) push(data interface{}) {

	if data == nil {
		perror.Panic("Iterator Stack does not allow nil data")
	}

	stk.top = &node{
		data: data,
		next: stk.top,
	}

	stk.size++
}

func (stk *Stack) Pop() interface{} {

	if stk.size == 0 {
		perror.Panic("Nothing to pop, check stack first")
	}

	data := stk.top.data
	stk.top = stk.top.next
	stk.size--

	return data
}

func (stk *Stack) MatchNext(other interface{}) bool {
	return stk.mtc.Match(stk.buff, other)
}

func (stk *Stack) AcceptPush(other interface{}) bool {

	if stk.MatchNext(other) {
		stk.push(stk.Next())
		return true
	}

	return false
}

func (stk *Stack) ExpectPush(other interface{}) error {

	e := stk.mtc.Expect(stk.buff, other)
	if e != nil {
		return e
	}

	stk.push(stk.Next())
	return nil
}

func (stk *Stack) MatchStk(other interface{}) bool {
	return stk.mtc.Match(stk.PeekStk(), other)
}

func (stk *Stack) AcceptPop(other interface{}) interface{} {

	if stk.MatchStk(other) {
		return stk.Pop()
	}

	return nil
}

func (stk *Stack) ExpectPop(other interface{}) (interface{}, error) {

	if stk.EmptyStk() {
		return nil, stk.mtc.Expect(nil, other)
	}

	e := stk.mtc.Expect(stk.top.data, other)
	if e != nil {
		return nil, e
	}

	return stk.Pop(), nil
}

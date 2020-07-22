package shunter

/*
import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type shunter struct {
	in   *Stack
	yard *Stack
	out  *Queue
}

func (shu *shunter) empty() bool {
	return shu.in.Empty() && shu.yard.Empty()
}

func (shu *shunter) more() bool {
	return shu.in.More() || shu.yard.More()
}

func (shu *shunter) acceptIntoYard(props ...prop.Prop) bool {

	if shu.in.Empty() {
		return false
	}

	if shu.in.Top().Is(props...) {
		shu.yard.Push(shu.in.Pop())
		return true
	}

	return false
}

func (shu *shunter) acceptIntoOut(props ...prop.Prop) bool {

	if shu.in.Empty() {
		return false
	}

	if shu.in.Top().Is(props...) {
		shu.out.Put(shu.in.Pop())
		return true
	}

	return false
}

func (shu *shunter) acceptYardToOut(props ...prop.Prop) {
	if shu.in.Top().Is(props...) {
		shu.out.Put(shu.in.Pop())
	}
}
*/

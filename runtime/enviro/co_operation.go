package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func popNumOperands(env *Environment) (left, right types.Num, ok bool) {

	right, ok = env.Pop().(types.Num)
	if !ok {
		env.Fail(perror.New("Expected number on right side of operation"))
		return
	}

	left, ok = env.Pop().(types.Num)
	if !ok {
		env.Fail(perror.New("Expected number on left side of operation"))
	}

	return
}

func coAdd(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	left.Add(right.Number)
	env.Push(left)
}

func coSub(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	left.Sub(right.Number)
	env.Push(left)
}

func coMul(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	left.Mul(right.Number)
	env.Push(left)
}

func coDiv(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	left.Div(right.Number)
	env.Push(left)
}

func coRem(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	left.Mod(right.Number)
	env.Push(left)
}

func popBoolOperands(env *Environment) (left, right types.Bool, ok bool) {

	right, ok = env.Pop().(types.Bool)
	if !ok {
		env.Fail(perror.New("Expected bool on right side of operation"))
		return
	}

	left, ok = env.Pop().(types.Bool)
	if !ok {
		env.Fail(perror.New("Expected bool on left side of operation"))
	}

	return
}

func coAnd(env *Environment, in inst.Instruction) {

	left, right, ok := popBoolOperands(env)
	if !ok {
		return
	}

	env.Push(left.And(right))
}

func coOr(env *Environment, in inst.Instruction) {

	left, right, ok := popBoolOperands(env)
	if !ok {
		return
	}

	env.Push(left.Or(right))
}

func coLess(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	answer := left.Less(right.Number)
	env.Push(types.Bool(answer))
}

func coMore(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	answer := left.More(right.Number)
	env.Push(types.Bool(answer))
}

func coLessOrEqual(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	answer := left.LessOrEqual(right.Number)
	env.Push(types.Bool(answer))
}

func coMoreOrEqual(env *Environment, in inst.Instruction) {

	left, right, ok := popNumOperands(env)
	if !ok {
		return
	}

	answer := left.MoreOrEqual(right.Number)
	env.Push(types.Bool(answer))
}

package processor

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

type runtimeEnv struct {
	// Program
	started bool
	counter int
	ins     []inst.Inst
	// Runtime
	value.Stack
	ids      map[value.Ident]value.Value
	book     spell.Book
	exitFlag bool
	exitCode int
	err      error
}

func (rt *runtimeEnv) Next() (inst.Inst, bool) {

	if rt.started {
		rt.counter++
	} else {
		rt.started = true
	}

	if rt.counter >= len(rt.ins) {
		return inst.Inst{}, false
	}

	return rt.ins[rt.counter], true
}

func (rt *runtimeEnv) Fetch(id value.Ident) value.Value {
	if v, ok := rt.ids[id]; ok {
		return v
	}
	rt.Fail(1, errors.New("Identifier "+string(id)+" not found in scope"))
	return nil
}

func (rt *runtimeEnv) FetchPush(id value.Ident) {
	if v, ok := rt.ids[id]; ok {
		rt.Push(v)
		return
	}
	rt.Fail(1, errors.New("Identifier "+string(id)+" not found in scope"))
}

func (rt *runtimeEnv) Spellbook() spell.Book {
	return rt.book
}

func (rt *runtimeEnv) Bind(id value.Ident, v value.Value) {
	rt.ids[id] = v
}

func (rt *runtimeEnv) Fail(code int, e error) {
	rt.exitCode = code
	rt.err = e
	rt.exitFlag = true
}

func (rt *runtimeEnv) Exit(code int) {
	rt.exitCode = code
	rt.exitFlag = true
}

func (rt *runtimeEnv) GetErr() error {
	return rt.err
}

func (rt *runtimeEnv) GetExitCode() int {
	return rt.exitCode
}

func (rt *runtimeEnv) GetExitFlag() bool {
	return rt.exitFlag
}

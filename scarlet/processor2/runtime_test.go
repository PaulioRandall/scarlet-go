package processor2

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

type testRuntime struct {
	scope    map[value.Ident]value.Value
	book     spell.Book
	exitFlag bool
	exitCode int
	err      error
}

func (env *testRuntime) Spellbook() spell.Book {
	return env.book
}

func (env *testRuntime) Bind(id value.Ident, v value.Value) {
	env.scope[id] = v
}

func (env *testRuntime) Fetch(id value.Ident) value.Value {
	if v, ok := env.scope[id]; ok {
		return v
	}
	env.Fail(1, errors.New("Identifier "+string(id)+" not found in scope"))
	return nil
}

func (env *testRuntime) Fail(code int, e error) {
	env.exitCode = code
	env.err = e
	env.exitFlag = true
}

func (env *testRuntime) Exit(code int) {
	env.exitCode = code
	env.exitFlag = true
}

func (env *testRuntime) GetErr() error {
	return env.err
}

func (env *testRuntime) GetExitFlag() bool {
	return env.exitFlag
}

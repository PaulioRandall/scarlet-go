package runtime

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"

	"github.com/PaulioRandall/scarlet-go/cmd/spells"
)

// RuntimeEnv represents a Runtime and an Environment for a specific program.
type RuntimeEnv struct {
	// processor.Runtime
	scope     spell.Scope
	spellbook spell.Book
	exitFlag  bool
	exitCode  int
	err       error
}

// New creates and returns a new RuntimeEnv for a specific program.
func New() *RuntimeEnv {
	return &RuntimeEnv{
		scope:     spell.Scope{},
		spellbook: spells.NewBook(),
	}
}

// Spellbook implements processor.Runtime.Spellbook.
func (env *RuntimeEnv) Spellbook() spell.Book {
	r := make(spell.Book, len(env.spellbook))
	for k, v := range env.spellbook {
		r[k] = v
	}
	return r
}

// Scope implements processor.Runtime.Scope.
func (env *RuntimeEnv) Scope() spell.Scope {
	r := make(spell.Scope, len(env.scope))
	for k, v := range env.scope {
		r[k] = v
	}
	return r
}

// Bind implements processor.Runtime.Bind.
func (env *RuntimeEnv) Bind(id value.Ident, v value.Value) {
	env.scope[id] = v
}

// Unbind implements processor.Runtime.Unbind.
func (env *RuntimeEnv) Unbind(id value.Ident) {
	delete(env.scope, id)
}

// Fetch implements processor.Runtime.Fetch.
func (env *RuntimeEnv) Exists(id value.Ident) value.Bool {
	_, ok := env.scope[id]
	return value.Bool(ok)
}

// Fetch implements processor.Runtime.Fetch.
func (env *RuntimeEnv) Fetch(id value.Ident) value.Value {
	if v, ok := env.scope[id]; ok {
		return v
	}
	env.Fail(1, errors.New("Identifier "+string(id)+" not found in scope"))
	return nil
}

// Fail implements processor.Runtime.Fail.
func (env *RuntimeEnv) Fail(exitCode int, e error) {
	env.exitCode = exitCode
	env.err = e
	env.exitFlag = true
}

// Exit implements processor.Runtime.Exit.
func (env *RuntimeEnv) Exit(exitCode int) {
	env.exitCode = exitCode
	env.exitFlag = true
}

// GetErr implements processor.Runtime.Exit.
func (env *RuntimeEnv) GetErr() error {
	return env.err
}

// GetExitCode implements processor.Runtime.GetExitCode.
func (env *RuntimeEnv) GetExitCode() int {
	return env.exitCode
}

// GetExitFlag implements processor.Runtime.GetExitFlag.
func (env *RuntimeEnv) GetExitFlag() bool {
	return env.exitFlag
}

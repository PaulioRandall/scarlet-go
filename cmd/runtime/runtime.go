package runtime

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"

	"github.com/PaulioRandall/scarlet-go/cmd/spells"
)

// IdMap represents a mapping of declared identifiers with their current values.
type IdMap map[value.Ident]value.Value

// RuntimeEnv represents a Runtime and an Environment for a specific program.
type RuntimeEnv struct {
	// processor.Program
	Counter int
	Program []inst.Inst
	Size    int
	// processor.Runtime
	value.Stack
	Scope     IdMap
	spellbook spell.Book
	exitFlag  bool
	exitCode  int
	err       error
}

// New creates and returns a new RuntimeEnv for a specific program.
func New(program []inst.Inst) *RuntimeEnv {
	return &RuntimeEnv{
		Counter:   -1,
		Program:   program,
		Size:      len(program),
		Scope:     IdMap{},
		spellbook: spells.NewBook(),
	}
}

// Next implements processor.Runtime.Next.
func (env *RuntimeEnv) Next() (inst.Inst, bool) {
	env.Counter++
	if env.Counter >= env.Size {
		return inst.Inst{}, false
	}
	return env.Program[env.Counter], true
}

// Fetch implements processor.Runtime.Fetch.
func (env *RuntimeEnv) Fetch(id value.Ident) value.Value {
	if v, ok := env.Scope[id]; ok {
		return v
	}
	env.Fail(1, errors.New("Identifier "+string(id)+" not found in scope"))
	return nil
}

// Fetch implements processor.Runtime.FetchPush.
func (env *RuntimeEnv) FetchPush(id value.Ident) {
	if v, ok := env.Scope[id]; ok {
		env.Push(v)
		return
	}
	env.Fail(1, errors.New("Identifier "+string(id)+" not found in scope"))
}

// Spellbook implements processor.Runtime.Spellbook.
func (env *RuntimeEnv) Spellbook() spell.Book {
	return env.spellbook
}

// Bind implements processor.Runtime.Bind.
func (env *RuntimeEnv) Bind(id value.Ident, v value.Value) {
	env.Scope[id] = v
}

// Unbind implements processor.Runtime.Unbind.
func (env *RuntimeEnv) Unbind(id value.Ident) {
	delete(env.Scope, id)
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

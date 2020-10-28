package runtime

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

// IdMap represents a mapping of declared identifiers with their current values.
type IdMap map[value.Ident]value.Value

// Environment represents a Runtime Environment for a specific program.
type Environment struct {
	value.Stack
	Counter int
	Program []inst.Inst
	Size    int
	Scope   IdMap
}

// New creates and returns a new Environment for a specific program.
func New(program []inst.Inst) *Environment {
	return &Environment{
		Counter: -1,
		Program: program,
		Size:    len(program),
		Scope:   IdMap{},
	}
}

// Next implements processor.Runtime.Next.
func (env *Environment) Next() (inst.Inst, bool) {
	env.Counter++
	if env.Counter >= env.Size {
		return inst.Inst{}, false
	}
	return env.Program[env.Counter], true
}

// Fetch implements processor.Runtime.Fetch.
func (env *Environment) Fetch(id value.Ident) (value.Value, error) {
	if v, ok := env.Scope[id]; ok {
		return v, nil
	}
	return nil, errors.New("Identifier " + string(id) + " not found in scope")
}

// Bind implements processor.Runtime.Bind.
func (env *Environment) Bind(id value.Ident, v value.Value) error {
	env.Scope[id] = v
	return nil
}

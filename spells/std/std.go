package std

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func Default() {
	InscribeAll(func(name string, spell spellbook.Spell) {
		e := spellbook.Inscribe(""+name, spell)
		if e != nil {
			panic(e)
		}
	})
}

func InscribeAll(inscribe spellbook.Inscriber) {
	inscribe("exit", Exit{})
	inscribe("print", Print{})
	inscribe("println", Println{})
	inscribe("set", Set{})
	inscribe("del", Del{})
}

type Exit struct{}

func (Exit) Summary() string {
	return `@Exit(exitcode)
	Exit terminates the current script with a specific exit code.`
}

func (sp Exit) Docs() string {
	return sp.Summary() + `
		@Exit(0)
		@Exit(1)`
}

func (Exit) Invoke(env spellbook.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("@Exit requires one argument"))
		return
	}

	if c, ok := args[0].(types.Num); ok {
		env.Exit(int(c.Integer()))
		return
	}

	env.Fail(errors.New("@Exit requires its argument be a number"))
}

var printSpellDocs = `@Print(value...)
@Println(value...)
	Prints all arguments to standard output in the order provided, if @Println
	is used then a linefeed is appended after.

	# Outputs: "Hello, Scarlet!"
	@Print("Hello, Scarlet!")
	@Println("Hello, Scarlet!")

	# Outputs: "a*b = c"
	@Print(a, "*", b, " = ", c)
	@Println(a, "*", b, " = ", c)`

type Print struct{}

func (Print) Summary() string {
	return `@Print(value...)
	Prints all arguments to standard output in the order provided`
}

func (Print) Docs() string {
	return printSpellDocs
}

func (Print) Invoke(_ spellbook.Enviro, args []types.Value) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

type Println struct{}

func (Println) Summary() string {
	return `@Println(value...)
	Prints all arguments to standard output in the order provided then appends
	a linefeed.`
}

func (Println) Docs() string {
	return printSpellDocs
}

func (Println) Invoke(_ spellbook.Enviro, args []types.Value) {
	Print{}.Invoke(nil, args)
	fmt.Println()
}

type Set struct{}

func (Set) Summary() string {
	return `@Set("identifier", value)
	Sets the value of variable represented by the first argument as the second
	argument.`
}

func (sp Set) Docs() string {
	return sp.Summary() + `
	
	# x := 1
	@Set("x", 1)
	
	# name := "Scarlet"
	@Set("name", "Scarlet")`
}

func (Set) Invoke(env spellbook.Enviro, args []types.Value) {

	if len(args) != 2 {
		env.Fail(errors.New("@Set requires two arguments"))
		return
	}

	idStr, ok := args[0].(types.Str)
	id := string(idStr)

	if !ok || !isIdentifier(id) {
		env.Fail(errors.New("@Set requires the first argument be an identifier string"))
		return
	}

	env.Bind(id, args[1])
}

type Del struct{}

func (Del) Summary() string {
	return `@Del("identifier")
	Deletes the variable represented by the first argument`
}

func (sp Del) Docs() string {
	return sp.Summary() + `
	
	# Deletes variable 'x'
	@Del("x")
	
	# Deletes varaibel 'name'
	@Set("name")`
}

func (Del) Invoke(env spellbook.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("@Del requires one argument"))
		return
	}

	id, ok := args[0].(types.Str)
	if !ok {
		env.Fail(errors.New("@Del requires its argument be an identifier string"))
		return
	}

	env.Unbind(string(id))
}

func isIdentifier(id string) bool {

	for i, ru := range id {

		if i == 0 {
			if !unicode.IsLetter(ru) {
				return false
			}

			continue
		}

		if !unicode.IsLetter(ru) || ru != '_' {
			return false
		}
	}

	return true
}

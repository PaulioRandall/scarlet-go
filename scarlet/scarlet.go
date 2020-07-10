package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/a_scan"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/b_sanitise"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/c_check"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/d_shunt"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/e_compile"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/f_runtime"
)

func main() { // Run it with `./godo run`

	exitCode, e := esme("test.scarlet")
	fmt.Println()

	if e != nil {
		fmt.Printf("[ERROR] %+v\n", e)
	}

	fmt.Printf("Exit code: %d\n", exitCode)

	todo()
	os.Exit(exitCode)
}

func esme(file string) (int, error) {

	const ERROR_CODE = 1

	s, e := ioutil.ReadFile(file)
	if e != nil {
		return ERROR_CODE, e
	}

	tks, e := scan.ScanAll(string(s))
	if e != nil {
		return ERROR_CODE, e
	}

	tks, e = sanitise.SanitiseAll(tks)
	if e != nil {
		return ERROR_CODE, e
	}

	tks, e = check.CheckAll(tks)
	if e != nil {
		return ERROR_CODE, e
	}

	tks, e = shunt.ShuntAll(tks)
	if e != nil {
		return ERROR_CODE, e
	}

	ins, e := compile.CompileAll(tks)
	if e != nil {
		return ERROR_CODE, e
	}

	rt := runtime.New(ins)
	_, e = rt.Start()
	return rt.ExitCode(), e
}

func todo() {
	println()
	println()
	println("TODO:")
	println("[Next] Test f_runtime pkg")
	println("[Next] Check an identifier is valid when using @set")
	println("[Next] Put spells in their own pkg & create spell register")
	println()
	println("[Think] About how to abstract test utilities")
	println()
	println("[Plan]")
	println("- a_scan:     scans in tokens including redundant ones")
	println("- b_sanitise: removes redundant tokens")
	println("- c_check:    checks the token sequence follows language rules")
	println("- d_shunt:    converts from infix to postfix notation")
	println("- e_compile:  converts the tokens into instructions")
	println("- f_runtime:  executes an instruction list")
	println("- ...")
}

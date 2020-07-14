package program

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

func run(args []string) error {
	return runEsme()
}

func runEsme() error { // Run it with `./godo run`

	exitCode, e := esme("test.scarlet")
	fmt.Println()

	if e != nil {
		fmt.Printf("[ERROR] %+v\n", e)
	}

	fmt.Printf("Exit code: %d\n", exitCode)

	os.Exit(exitCode)
	return nil
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
	rt.Start()
	return rt.Env().ExitCode, rt.Env().Err
}

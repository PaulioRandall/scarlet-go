package main

import (
	"fmt"
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/parser"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/runtime"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/scanner"
	//"github.com/pkg/errors"

	a_scan "github.com/PaulioRandall/scarlet-go/pkg/rincewind/a_scan"
	b_sanitise "github.com/PaulioRandall/scarlet-go/pkg/rincewind/b_sanitise"
	c_check "github.com/PaulioRandall/scarlet-go/pkg/rincewind/c_check"
	d_shunt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/d_shunt"
	e_compile "github.com/PaulioRandall/scarlet-go/pkg/rincewind/e_compile"
	f_runtime "github.com/PaulioRandall/scarlet-go/pkg/rincewind/f_runtime"
)

func main() { // Run it with `./godo run`

	e := rince("test.scarlet")
	if e != nil {
		fmt.Printf("%+v\n", e)
		return
	}

	//esme()

	println()
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

type symItr struct {
	scanner.SymItr
	symbols []rune
	size    int
	i       int
}

func rince(file string) error {

	s, e := ioutil.ReadFile(file)
	if e != nil {
		return e
	}

	tks, e := a_scan.ScanAll(string(s))
	if e != nil {
		return e
	}

	tks, e = b_sanitise.SanitiseAll(tks)
	if e != nil {
		return e
	}

	tks, e = c_check.CheckAll(tks)
	if e != nil {
		return e
	}

	tks, e = d_shunt.ShuntAll(tks)
	if e != nil {
		return e
	}

	ins, e := e_compile.CompileAll(tks)
	if e != nil {
		return e
	}

	rt := f_runtime.New(ins)
	_, e = rt.Start()
	return e
}

func esme() {

	f := "test.scarlet"
	s, e := ioutil.ReadFile(f)
	if e != nil {
		panic(e)
	}

	tks, e := scan(string(s))
	if e != nil {
		fmt.Printf("%+v\n", e)
		return
	}

	sts, e := parse(tks)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return
	}

	ctx, e := run(sts)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return
	}

	println()
	println(ctx.String())
}

func (itr *symItr) Next() (rune, bool) {

	if itr.i >= itr.size {
		return rune(0), false
	}

	ru := itr.symbols[itr.i]
	itr.i++
	return ru, true
}

func scan(s string) ([]token.Token, error) {

	println("# Scanning...")

	itr := &symItr{
		symbols: []rune(s),
		size:    len(s),
		i:       0,
	}

	var (
		tks = []token.Token{}
		tk  token.Token
		f   scanner.ScanFunc
		e   error
	)

	for f = scanner.New(itr); f != nil; {

		tk, f, e = f()
		if e != nil {
			return nil, e
		}

		tks = append(tks, tk)
	}

	return tks, nil
}

type tokenStream struct {
	parser.TokenStream
	tokens []token.Token
	size   int
	i      int
}

func (s *tokenStream) Next() token.Token {

	if s.i >= s.size {
		return nil
	}

	tk := s.tokens[s.i]
	s.i++
	return tk
}

func parse(tks []token.Token) ([]stats.Expr, error) {

	println("# Parsing...")

	ts := &tokenStream{
		tokens: tks,
		size:   len(tks),
		i:      0,
	}

	var (
		exprs = []stats.Expr{}
		expr  stats.Expr
		f     parser.ParseFunc
		e     error
	)

	for f = parser.New(ts); f != nil; {

		expr, f, e = f()
		if e != nil {
			return nil, e
		}

		exprs = append(exprs, expr)
	}

	return exprs, e
}

func run(sts []stats.Expr) (*runtime.Context, error) {
	println("# Executing...")
	return runtime.Run(sts)
}

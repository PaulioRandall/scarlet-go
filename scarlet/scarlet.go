package main

import (
	"fmt"
	"io/ioutil"
	//	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/stats"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/parser"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/runtime"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/scanner"
	//"github.com/pkg/errors"
)

func main() { // Run it with `./godo run`

	esme()

	println()
	println("[Next] Start work on c_inst")
	println("[Plan]")
	println("- a_scan:  scans in tokens including redundant ones")
	println("- b_group: groups tokens into statements")
	println("- c_inst:  sub-divides statements into instructions")
	println("- d_check: checks sequences of instructions follow language rules")
	println("- e_amass: amalgamates instructions into a single list")
	println("- f_exec:  executes an instruction list")
	println("- ...")
}

type symItr struct {
	scanner.SymItr
	symbols []rune
	size    int
	i       int
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

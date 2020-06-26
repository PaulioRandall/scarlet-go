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

type symItr struct {
	scanner.SymItr
	symbols []rune
	size    int
	i       int
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

func (s *tokenStream) peek(i int) token.Token {

	if i >= s.size {
		return nil
	}

	return s.tokens[i]
}

func (s *tokenStream) Next() token.Token {

	tk := s.peek(s.i)
	if tk == nil {
		return nil
	}

	s.i++
	return tk
}

func (s *tokenStream) Peek() token.Token {
	return s.peek(s.i)
}

func (s *tokenStream) PeekBeyond() token.Token {
	return s.peek(s.i + 1)
}

func parse(tks []token.Token) ([]stats.Expr, error) {

	println("# Parsing...")

	stream := &tokenStream{
		tokens: tks,
		size:   len(tks),
		i:      0,
	}

	return parser.ParseStatements(stream)
}

func run(sts []stats.Expr) (*runtime.Context, error) {
	println("# Executing...")
	return runtime.Run(sts)
}

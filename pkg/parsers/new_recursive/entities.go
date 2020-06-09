package recursive

import (
	"fmt"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type builder struct {
	strings.Builder
}

func (b *builder) add(indent int, s string) {

	for _, ru := range s {
		b.WriteRune(ru)

		if ru == '\n' {
			for i := 0; i < indent; i++ {
				b.WriteRune('\t')
			}
		}
	}
}

func (b *builder) addToken(indent int, tk Token) {
	b.add(indent, ToString(tk))
}

func (b *builder) newline() {
	b.add(0, "\n")
}

func (b *builder) String() string {
	return b.Builder.String()
}

func (b *builder) print() {
	fmt.Println(b.String())
	fmt.Println()
}

func startPos(tk Token) (line int, col int) {
	return tk.Line(), tk.Col()
}

func endPos(tk Token) (line int, col int) {
	return tk.Line(), tk.Col() + len(tk.Value())
}

type voidExpr struct {
	tk Token
}

func (voidExpr) Kind() Kind {
	return ST_VOID
}

func (v voidExpr) Tk() Token {
	return v.tk
}

func (v voidExpr) Begin() (int, int) {
	return startPos(v.tk)
}

func (v voidExpr) End() (int, int) {
	return endPos(v.tk)
}

func (v voidExpr) String() string {
	return VoidString(v)
}

type identifierExpr struct {
	tk Token
}

func (identifierExpr) Kind() Kind {
	return ST_VOID
}

func (id identifierExpr) Tk() Token {
	return id.tk
}

func (id identifierExpr) Begin() (int, int) {
	return startPos(id.tk)
}

func (id identifierExpr) End() (int, int) {
	return endPos(id.tk)
}

func (id identifierExpr) String() string {
	return IdentifierString(id)
}

type literalExpr struct {
	tk Token
}

func (literalExpr) Kind() Kind {
	return ST_LITERAL
}

func (l literalExpr) Tk() Token {
	return l.tk
}

func (l literalExpr) Begin() (int, int) {
	return startPos(l.tk)
}

func (l literalExpr) End() (int, int) {
	return endPos(l.tk)
}

func (l literalExpr) String() string {
	return LiteralString(l)
}

type listAccessorExpr struct {
	id    Expression
	index Expression
}

func (listAccessorExpr) Kind() Kind {
	return ST_LIST_ACCESSOR
}

func (l listAccessorExpr) ID() Expression {
	return l.id
}

func (l listAccessorExpr) Index() Expression {
	return l.id
}

func (l listAccessorExpr) Begin() (int, int) {
	return l.id.Begin()
}

func (l listAccessorExpr) End() (int, int) {
	return l.index.End()
}

func (l listAccessorExpr) String() string {
	return ListAccessorString(l)
}

type listConstructorExpr struct {
	open, close Token
	items       []Expression
}

func (listConstructorExpr) Kind() Kind {
	return ST_LIST_CONSTRUCTOR
}

func (l listConstructorExpr) Open() Token {
	return l.open
}

func (l listConstructorExpr) Close() Token {
	return l.close
}

func (l listConstructorExpr) Items() []Expression {
	return l.items
}

func (l listConstructorExpr) Begin() (int, int) {
	return startPos(l.open)
}

func (l listConstructorExpr) End() (int, int) {
	return endPos(l.close)
}

func (l listConstructorExpr) String() string {
	return ListConstructorString(l)
}

type negationExpr struct {
	expr Expression
}

func (negationExpr) Kind() Kind {
	return ST_NEGATION
}

func (n negationExpr) Expr() Expression {
	return n.expr
}

func (n negationExpr) Begin() (int, int) {
	return n.expr.Begin()
}

func (n negationExpr) End() (int, int) {
	return n.expr.End()
}

func (n negationExpr) String() string {
	return NegationString(n)
}

type assignmentStat struct {
	target Expression
	source Expression
}

func (assignmentStat) Kind() Kind {
	return ST_ASSIGNMENT
}

func (a assignmentStat) Target() Expression {
	return a.target
}

func (a assignmentStat) Source() Expression {
	return a.source
}

func (a assignmentStat) Begin() (int, int) {
	return a.target.Begin()
}

func (a assignmentStat) End() (int, int) {
	return a.source.End()
}

func (a assignmentStat) String() string {

	b := builder{}

	b.add(0, "[Assignment] ")

	b.newline()
	b.add(1, "Target: ")
	b.newline()
	b.add(1, a.target.String())

	b.newline()
	b.add(1, "Source: ")
	b.newline()
	b.add(2, a.source.String())

	return b.String()
}

type Block struct {
	start, end Token
	Stats      []Statement
}

func (bk Block) Begin() (int, int) {
	return startPos(bk.start)
}

func (bk Block) End() (int, int) {
	return startPos(bk.end)
}

func (bk Block) String() string {

	b := builder{}

	for _, a := range bk.Stats {
		b.add(0, a.String())
	}

	return b.String()
}

type Parameters struct {
	open, close Token
	Inputs      []Token
	Outputs     []Token
}

func (p Parameters) Begin() (int, int) {
	return startPos(p.open)
}

func (p Parameters) End() (int, int) {
	return endPos(p.close)
}

func (p Parameters) String() string {

	b := builder{}

	b.add(0, "[Parameters] ")

	if len(p.Inputs) > 0 {
		b.newline()
		b.add(1, "Inputs: ")

		for _, in := range p.Inputs {
			b.newline()
			b.addToken(2, in)
		}
	}

	if len(p.Inputs) > 0 {
		b.newline()
		b.add(1, "Outputs: ")

		for _, out := range p.Outputs {
			b.newline()
			b.addToken(2, out)
		}
	}

	return b.String()
}

type Function struct {
	key    Token
	Params Parameters
	Body   Block
}

func (f Function) Begin() (int, int) {
	return startPos(f.key)
}

func (f Function) End() (int, int) {
	return f.Body.End()
}

func (f Function) String() string {

	b := builder{}

	b.add(0, "[Function] ")

	b.newline()
	b.add(1, f.Params.String())

	b.newline()
	b.add(1, f.Body.String())

	return b.String()
}

type Operation struct {
	Operator    Token
	Left, Right Expression
}

func (o Operation) Precedence() int {
	return o.Operator.Morpheme().Precedence()
}

func (o Operation) Begin() (int, int) {
	return o.Left.Begin()
}

func (o Operation) End() (int, int) {
	return o.Right.End()
}

func (o Operation) String() string {

	b := builder{}

	b.add(0, "[Operation] ")
	b.addToken(0, o.Operator)

	b.newline()
	b.add(1, "Left: ")
	b.newline()
	b.add(2, o.Left.String())

	b.newline()
	b.add(1, "Right: ")
	b.newline()
	b.add(2, o.Right.String())

	return b.String()
}

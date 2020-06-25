package stats

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type exitStat struct {
	tk   Token
	code Expr
}

func (exitStat) Kind() Kind {
	return ST_EXIT
}

func (e exitStat) Tk() Token {
	return e.tk
}

func (e exitStat) Code() Expr {
	return e.code
}

func (e exitStat) Begin() (int, int) {
	return e.tk.Begin()
}

func (e exitStat) End() (int, int) {
	return e.code.End()
}

func (e exitStat) String() string {
	return ExitString(e)
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
	return v.tk.Begin()
}

func (v voidExpr) End() (int, int) {
	return v.tk.End()
}

func (v voidExpr) String() string {
	return VoidString(v)
}

type identifierExpr struct {
	tk Token
}

func (identifierExpr) Kind() Kind {
	return ST_IDENTIFIER
}

func (id identifierExpr) Tk() Token {
	return id.tk
}

func (id identifierExpr) Begin() (int, int) {
	return id.tk.Begin()
}

func (id identifierExpr) End() (int, int) {
	return id.tk.End()
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
	return l.tk.Begin()
}

func (l literalExpr) End() (int, int) {
	return l.tk.End()
}

func (l literalExpr) String() string {
	return LiteralString(l)
}

type containerItemExpr struct {
	container Expr
	key       Expr
}

func (containerItemExpr) Kind() Kind {
	return ST_CONTAINER_ITEM
}

func (c containerItemExpr) Container() Expr {
	return c.container
}

func (c containerItemExpr) Key() Expr {
	return c.key
}

func (c containerItemExpr) Begin() (int, int) {
	return c.container.Begin()
}

func (c containerItemExpr) End() (int, int) {
	return c.key.End()
}

func (c containerItemExpr) String() string {
	return ContainerItemString(c)
}

type negationExpr struct {
	expr Expr
}

func (negationExpr) Kind() Kind {
	return ST_NEGATION
}

func (n negationExpr) Expr() Expr {
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

type operationExpr struct {
	operator    Token
	left, right Expr
}

func (operationExpr) Kind() Kind {
	return ST_OPERATION
}

func (o operationExpr) Operator() Token {
	return o.operator
}

func (o operationExpr) Left() Expr {
	return o.left
}

func (o operationExpr) Right() Expr {
	return o.right
}

func (o operationExpr) Begin() (int, int) {
	return o.left.Begin()
}

func (o operationExpr) End() (int, int) {
	return o.right.End()
}

func (o operationExpr) String() string {
	return OperationString(o)
}

type assignStat struct {
	final  bool
	target Expr
	source Expr
}

func (assignStat) Kind() Kind {
	return ST_ASSIGN
}

func (a assignStat) Const() bool {
	return a.final
}

func (a assignStat) Target() Expr {
	return a.target
}

func (a assignStat) Source() Expr {
	return a.source
}

func (a assignStat) Begin() (int, int) {
	return a.target.Begin()
}

func (a assignStat) End() (int, int) {
	return a.source.End()
}

func (a assignStat) String() string {
	return AssignString(a)
}

type assignBlockStat struct {
	final   bool
	targets []Expr
	sources []Expr
	count   int
}

func (assignBlockStat) Kind() Kind {
	return ST_ASSIGN_BLOCK
}

func (a assignBlockStat) Const() bool {
	return a.final
}

func (a assignBlockStat) Targets() []Expr {
	return a.targets
}

func (a assignBlockStat) Sources() []Expr {
	return a.sources
}

func (a assignBlockStat) Count() int {
	return a.count
}

func (a assignBlockStat) Begin() (int, int) {
	return a.targets[0].Begin()
}

func (a assignBlockStat) End() (int, int) {
	return a.sources[a.count-1].End()
}

func (a assignBlockStat) String() string {
	return AssignBlockString(a)
}

type blockExpr struct {
	open, close Token
	stats       []Expr
}

func (blockExpr) Kind() Kind {
	return ST_BLOCK
}

func (bk blockExpr) Stats() []Expr {
	return bk.stats
}

func (bk blockExpr) Begin() (int, int) {
	return bk.open.Begin()
}

func (bk blockExpr) End() (int, int) {
	return bk.close.End()
}

func (bk blockExpr) String() string {
	return BlockString(bk)
}

type undelimBlockExpr struct {
	stats []Expr
}

func (undelimBlockExpr) Kind() Kind {
	return ST_BLOCK
}

func (bk undelimBlockExpr) Stats() []Expr {
	return bk.stats
}

func (bk undelimBlockExpr) Begin() (int, int) {
	return bk.stats[0].Begin()
}

func (bk undelimBlockExpr) End() (int, int) {
	i := len(bk.stats) - 1
	return bk.stats[i].End()
}

func (bk undelimBlockExpr) String() string {
	return BlockString(bk)
}

type exprFuncExpr struct {
	key    Token
	inputs []Token
	expr   Expr
}

func (exprFuncExpr) Kind() Kind {
	return ST_EXPR_FUNC
}

func (e exprFuncExpr) Inputs() []Token {
	return e.inputs
}

func (e exprFuncExpr) Expr() Expr {
	return e.expr
}

func (e exprFuncExpr) Begin() (int, int) {
	return e.key.Begin()
}

func (e exprFuncExpr) End() (int, int) {
	return e.expr.End()
}

func (e exprFuncExpr) String() string {
	return ExprFuncString(e)
}

type funcDefExpr struct {
	key     Token
	inputs  []Token
	outputs []Token
	body    Expr
}

func (funcDefExpr) Kind() Kind {
	return ST_FUNC_DEF
}

func (f funcDefExpr) Inputs() []Token {
	return f.inputs
}

func (f funcDefExpr) Outputs() []Token {
	return f.outputs
}

func (f funcDefExpr) Body() Expr {
	return f.body
}

func (f funcDefExpr) Begin() (int, int) {
	return f.key.Begin()
}

func (f funcDefExpr) End() (int, int) {
	return f.body.End()
}

func (f funcDefExpr) String() string {
	return FuncDefString(f)
}

type funcCallExpr struct {
	close    Token
	function Expr
	args     []Expr
}

func (funcCallExpr) Kind() Kind {
	return ST_FUNC_CALL
}

func (f funcCallExpr) Function() Expr {
	return f.function
}

func (f funcCallExpr) Arguments() []Expr {
	return f.args
}

func (f funcCallExpr) Begin() (int, int) {
	return f.function.Begin()
}

func (f funcCallExpr) End() (int, int) {
	return f.close.End()
}

func (f funcCallExpr) String() string {
	return FuncCallString(f)
}

type watchStat struct {
	key  Token
	ids  []Token
	body Block
}

func (watchStat) Kind() Kind {
	return ST_WATCH
}

func (w watchStat) Identifiers() []Token {
	return w.ids
}

func (w watchStat) Body() Block {
	return w.body
}

func (w watchStat) Begin() (int, int) {
	return w.key.Begin()
}

func (w watchStat) End() (int, int) {
	return w.body.End()
}

func (w watchStat) String() string {
	return WatchString(w)
}

type guardStat struct {
	open      Token
	condition Expr
	body      Block
}

func (guardStat) Kind() Kind {
	return ST_GUARD
}

func (g guardStat) Condition() Expr {
	return g.condition
}

func (g guardStat) Body() Block {
	return g.body
}

func (g guardStat) Begin() (int, int) {
	return g.open.Begin()
}

func (g guardStat) End() (int, int) {
	return g.body.End()
}

func (g guardStat) String() string {
	return GuardString(g)
}

type whenCaseStat struct {
	object Expr
	body   Block
}

func (whenCaseStat) Kind() Kind {
	return ST_WHEN_CASE
}

func (wc whenCaseStat) Condition() Expr {
	return wc.object
}

func (wc whenCaseStat) Body() Block {
	return wc.body
}

func (wc whenCaseStat) Begin() (int, int) {
	return wc.object.Begin()
}

func (wc whenCaseStat) End() (int, int) {
	return wc.body.End()
}

func (wc whenCaseStat) String() string {
	return WhenCaseString(wc)
}

type whenStat struct {
	key, close Token
	init       Assign
	cases      []WhenCase
}

func (whenStat) Kind() Kind {
	return ST_WHEN
}

func (w whenStat) Initialiser() Assign {
	return w.init
}

func (w whenStat) Cases() []WhenCase {
	return w.cases
}

func (w whenStat) Begin() (int, int) {
	return w.key.Begin()
}

func (w whenStat) End() (int, int) {
	return w.close.End()
}

func (w whenStat) String() string {
	return WhenString(w)
}

type loopStat struct {
	key   Token
	init  Assign
	guard Guard
}

func (loopStat) Kind() Kind {
	return ST_LOOP
}

func (l loopStat) Initialiser() Assign {
	return l.init
}

func (l loopStat) Guard() Guard {
	return l.guard
}

func (l loopStat) Begin() (int, int) {
	return l.key.Begin()
}

func (l loopStat) End() (int, int) {
	return l.guard.End()
}

func (l loopStat) String() string {
	return LoopString(l)
}

type spellCallExpr struct {
	spell, close Token
	args         []Expr
}

func (spellCallExpr) Kind() Kind {
	return ST_SPELL_CALL
}

func (s spellCallExpr) Spell() Token {
	return s.spell
}

func (s spellCallExpr) Args() []Expr {
	return s.args
}

func (s spellCallExpr) Begin() (int, int) {
	return s.spell.Begin()
}

func (s spellCallExpr) End() (int, int) {
	return s.close.End()
}

func (s spellCallExpr) String() string {
	return SpellCallString(s)
}

type existsExpr struct {
	close   Token
	subject Expr
}

func (existsExpr) Kind() Kind {
	return ST_EXISTS
}

func (e existsExpr) Subject() Expr {
	return e.subject
}

func (e existsExpr) Begin() (int, int) {
	return e.subject.Begin()
}

func (e existsExpr) End() (int, int) {
	return e.close.End()
}

func (e existsExpr) String() string {
	return ExistsString(e)
}

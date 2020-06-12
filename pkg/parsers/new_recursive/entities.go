package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

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

type collectionAccessorExpr struct {
	collection Expression
	key        Expression
}

func (collectionAccessorExpr) Kind() Kind {
	return ST_COLLECTION_ACCESSOR
}

func (c collectionAccessorExpr) Collection() Expression {
	return c.collection
}

func (c collectionAccessorExpr) Key() Expression {
	return c.key
}

func (c collectionAccessorExpr) Begin() (int, int) {
	return c.collection.Begin()
}

func (c collectionAccessorExpr) End() (int, int) {
	return c.key.End()
}

func (c collectionAccessorExpr) String() string {
	return CollectionAccessorString(c)
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

type operationExpr struct {
	operator    Token
	left, right Expression
}

func (operationExpr) Kind() Kind {
	return ST_OPERATION
}

func (o operationExpr) Operator() Token {
	return o.operator
}

func (o operationExpr) Left() Expression {
	return o.left
}

func (o operationExpr) Right() Expression {
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
	return AssignmentString(a)
}

type assignmentBlockStat struct {
	assignments []Assignment
}

func (assignmentBlockStat) Kind() Kind {
	return ST_ASSIGNMENT_BLOCK
}

func (a assignmentBlockStat) Assignments() []Assignment {
	return a.assignments
}

func (a assignmentBlockStat) Begin() (int, int) {
	return a.assignments[0].Begin()
}

func (a assignmentBlockStat) End() (int, int) {
	i := len(a.assignments) - 1
	return a.assignments[i].End()
}

func (a assignmentBlockStat) String() string {
	return AssignmentBlockString(a)
}

type blockExpr struct {
	open, close Token
	stats       []Expression
}

func (blockExpr) Kind() Kind {
	return ST_BLOCK
}

func (bk blockExpr) Stats() []Expression {
	return bk.stats
}

func (bk blockExpr) Begin() (int, int) {
	return startPos(bk.open)
}

func (bk blockExpr) End() (int, int) {
	return endPos(bk.close)
}

func (bk blockExpr) String() string {
	return BlockString(bk)
}

type unDelimiteredBlockExpr struct {
	stats []Expression
}

func (unDelimiteredBlockExpr) Kind() Kind {
	return ST_BLOCK
}

func (bk unDelimiteredBlockExpr) Stats() []Expression {
	return bk.stats
}

func (bk unDelimiteredBlockExpr) Begin() (int, int) {
	return bk.stats[0].Begin()
}

func (bk unDelimiteredBlockExpr) End() (int, int) {
	i := len(bk.stats) - 1
	return bk.stats[i].End()
}

func (bk unDelimiteredBlockExpr) String() string {
	return BlockString(bk)
}

type expressionFunctionExpr struct {
	key    Token
	inputs []Token
	expr   Expression
}

func (expressionFunctionExpr) Kind() Kind {
	return ST_EXPRESSION_FUNCTION
}

func (e expressionFunctionExpr) Inputs() []Token {
	return e.inputs
}

func (e expressionFunctionExpr) Expr() Expression {
	return e.expr
}

func (e expressionFunctionExpr) Begin() (int, int) {
	return startPos(e.key)
}

func (e expressionFunctionExpr) End() (int, int) {
	return e.expr.End()
}

func (e expressionFunctionExpr) String() string {
	return ExpressionFunctionString(e)
}

type parametersDef struct {
	open, close Token
	inputs      []Token
	outputs     []Token
}

func (parametersDef) Kind() Kind {
	return ST_PARAMETERS
}

func (p parametersDef) Inputs() []Token {
	return p.inputs
}

func (p parametersDef) Outputs() []Token {
	return p.outputs
}

func (p parametersDef) Begin() (int, int) {
	return startPos(p.open)
}

func (p parametersDef) End() (int, int) {
	return endPos(p.close)
}

func (p parametersDef) String() string {
	return ParametersString(p)
}

type functionExpr struct {
	key    Token
	params Parameters
	body   Block
}

func (functionExpr) Kind() Kind {
	return ST_FUNCTION
}

func (f functionExpr) Params() Parameters {
	return f.params
}

func (f functionExpr) Body() Block {
	return f.body
}

func (f functionExpr) Begin() (int, int) {
	return startPos(f.key)
}

func (f functionExpr) End() (int, int) {
	return f.body.End()
}

func (f functionExpr) String() string {
	return FunctionString(f)
}

type functionCallExpr struct {
	close    Token
	function Expression
	args     []Expression
}

func (functionCallExpr) Kind() Kind {
	return ST_FUNCTION_CALL
}

func (f functionCallExpr) Function() Expression {
	return f.function
}

func (f functionCallExpr) Arguments() []Expression {
	return f.args
}

func (f functionCallExpr) Begin() (int, int) {
	return f.function.Begin()
}

func (f functionCallExpr) End() (int, int) {
	return endPos(f.close)
}

func (f functionCallExpr) String() string {
	return FunctionCallString(f)
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
	return startPos(w.key)
}

func (w watchStat) End() (int, int) {
	return w.body.End()
}

func (w watchStat) String() string {
	return WatchString(w)
}

type guardStat struct {
	open      Token
	condition Expression
	body      Block
}

func (guardStat) Kind() Kind {
	return ST_GUARD
}

func (g guardStat) Condition() Expression {
	return g.condition
}

func (g guardStat) Body() Block {
	return g.body
}

func (g guardStat) Begin() (int, int) {
	return startPos(g.open)
}

func (g guardStat) End() (int, int) {
	return g.body.End()
}

func (g guardStat) String() string {
	return GuardString(g)
}

type whenCaseStat struct {
	object Expression
	body   Block
}

func (whenCaseStat) Kind() Kind {
	return ST_WHEN_CASE
}

func (wc whenCaseStat) Condition() Expression {
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
	init       Assignment
	cases      []WhenCase
}

func (whenStat) Kind() Kind {
	return ST_WHEN
}

func (w whenStat) Initialiser() Assignment {
	return w.init
}

func (w whenStat) Cases() []WhenCase {
	return w.cases
}

func (w whenStat) Begin() (int, int) {
	return startPos(w.key)
}

func (w whenStat) End() (int, int) {
	return endPos(w.close)
}

func (w whenStat) String() string {
	return WhenString(w)
}

type loopStat struct {
	key   Token
	init  Assignment
	guard Guard
}

func (loopStat) Kind() Kind {
	return ST_LOOP
}

func (l loopStat) Initialiser() Assignment {
	return l.init
}

func (l loopStat) Guard() Guard {
	return l.guard
}

func (l loopStat) Begin() (int, int) {
	return startPos(l.key)
}

func (l loopStat) End() (int, int) {
	return l.guard.End()
}

func (l loopStat) String() string {
	return LoopString(l)
}

type spellCallExpr struct {
	spell, close Token
	args         []Expression
}

func (spellCallExpr) Kind() Kind {
	return ST_SPELL_CALL
}

func (s spellCallExpr) Spell() Token {
	return s.spell
}

func (s spellCallExpr) Arguments() []Expression {
	return s.args
}

func (s spellCallExpr) Begin() (int, int) {
	return startPos(s.spell)
}

func (s spellCallExpr) End() (int, int) {
	return endPos(s.close)
}

func (s spellCallExpr) String() string {
	return SpellCallString(s)
}

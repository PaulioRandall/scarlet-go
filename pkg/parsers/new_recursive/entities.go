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

type matchCaseStat struct {
	object Expression
	body   Block
}

func (matchCaseStat) Kind() Kind {
	return ST_MATCH_CASE
}

func (mc matchCaseStat) Condition() Expression {
	return mc.object
}

func (mc matchCaseStat) Body() Block {
	return mc.body
}

func (mc matchCaseStat) Begin() (int, int) {
	return mc.object.Begin()
}

func (mc matchCaseStat) End() (int, int) {
	return mc.body.End()
}

func (mc matchCaseStat) String() string {
	return MatchCaseString(mc)
}

type matchStat struct {
	key, close Token
	input      Expression
	cases      []MatchCase
}

func (matchStat) Kind() Kind {
	return ST_MATCH
}

func (m matchStat) Input() Expression {
	return m.input
}

func (m matchStat) Cases() []MatchCase {
	return m.cases
}

func (m matchStat) Begin() (int, int) {
	return startPos(m.key)
}

func (m matchStat) End() (int, int) {
	return endPos(m.close)
}

func (m matchStat) String() string {
	return MatchString(m)
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

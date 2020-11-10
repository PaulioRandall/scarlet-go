package tree

type (
	// Position represents a point within a text source.
	Position interface {
		Offset() int // bytes
		Line() int   // index
		ColByte() int
		ColRune() int
	}

	// Range represents a snippet between two points within a text source.
	Range interface {
		Begin() Position
		End() Position
	}

	// Node represents a node in a syntax tree.
	Node interface {
		Pos() Range
		node()
	}

	// Assignee is a Node that represents something that can have value bound
	// to it, i.e. an identifier.
	Assignee interface {
		Node
		assignee()
	}

	// Expr (Expression) is a Node that represents a traditional programmers
	// expression, i.e. a statement that always returns a single result.
	Expr interface {
		Node
		expr()
	}

	// MultiExpr (Multi return Expression) is a Node that represents a
	// an expression returning multiple values such as spell calls and functions.
	MultiExpr interface {
		Node
		multiExpr()
	}

	// Literal is a Node that represents a literal value such as a bool, a number
	// or a string.
	Literal interface {
		Node
		Expr
		literal()
	}

	// Stat (Statement) is a Node representing a traditional programmers
	// statement.
	Stat interface {
		Node
		stat()
	}

	// Guard is a Node representing a guarded statement or block.
	Guard interface {
		Node
		Condition() Expr
		guard()
	}
)

type (
	// Ident Node is an Expr representing an identifier.
	Ident struct {
		Range Range
		Val   string // Identifier name as defined in source
	}

	// AnonIdent Node is an Expr representing an anonymous identifier such as
	// will be used for ignoring a function or spell result.
	AnonIdent struct {
		Range Range
	}

	// BoolLit Node is an Expr representing a literal boolean.
	BoolLit struct {
		Range Range
		Val   bool
	}

	// NumLit Node is an Expr representing a literal number.
	NumLit struct {
		Range Range
		Val   float64
	}

	// StrLit Node is an Expr representing a literal string.
	StrLit struct {
		Range Range
		Val   string
	}

	// SingleAssign Node is a Stat representing a single assignment.
	SingleAssign struct {
		Range Range
		Left  Assignee
		Infix Range
		Right Expr
	}

	// AsymAssign Node is a Stat representing an assignment with multiple
	// target identifiers but only one expression, a function or spell call.
	AsymAssign struct {
		Range Range
		Left  []Assignee // Ordered left to right
		Infix Range
		Right Expr
	}

	// MultiAssign Node is a Stat representing a multiple assignment.
	MultiAssign struct {
		Range Range
		Left  []Assignee // Ordered left to right
		Infix Range
		Right []Expr // Ordered left to right
	}

	// UnaryExpr Node is an Expr representing a unary operation.
	UnaryExpr struct {
		Range Range
		Term  Expr
		Op    Operator
	}

	// BinaryExpr Node is an Expr representing an operation with two operands.
	BinaryExpr struct {
		Range Range
		Left  Expr
		Op    Operator
		Right Expr
	}

	// SpellCall Node is a MultiExpr representing a spell invocation.
	SpellCall struct {
		Range Range
		Name  string
		Args  []Expr
	}

	// Block Node is a Stat representing a block of statements.
	Block struct {
		Range Range
		Stmts []Node
	}

	// GuardedStmt Node is a Stat representing a conditional statement.
	GuardedStmt struct {
		Range Range
		Cond  Expr
		Stmt  Stat
	}

	// GuardedBlock Node is a Stat representing a conditional block.
	GuardedBlock struct {
		Range Range
		Cond  Expr
		Body  Block
	}

	// When is a Node representing a guarded block of guards where only the
	// statement of the first  matching case is executed.
	When struct {
		Range Range
		Cases []Guard
	}
)

func (n Ident) Pos() Range        { return n.Range }
func (n AnonIdent) Pos() Range    { return n.Range }
func (n BoolLit) Pos() Range      { return n.Range }
func (n NumLit) Pos() Range       { return n.Range }
func (n StrLit) Pos() Range       { return n.Range }
func (n SingleAssign) Pos() Range { return n.Range }
func (n AsymAssign) Pos() Range   { return n.Range }
func (n MultiAssign) Pos() Range  { return n.Range }
func (n UnaryExpr) Pos() Range    { return n.Range }
func (n BinaryExpr) Pos() Range   { return n.Range }
func (n SpellCall) Pos() Range    { return n.Range }
func (n Block) Pos() Range        { return n.Range }
func (n GuardedStmt) Pos() Range  { return n.Range }
func (n GuardedBlock) Pos() Range { return n.Range }
func (n When) Pos() Range         { return n.Range }

func (n GuardedStmt) Condition() Expr  { return n.Cond }
func (n GuardedBlock) Condition() Expr { return n.Cond }

func (n Ident) node()        {}
func (n AnonIdent) node()    {}
func (n BoolLit) node()      {}
func (n NumLit) node()       {}
func (n StrLit) node()       {}
func (n SingleAssign) node() {}
func (n AsymAssign) node()   {}
func (n MultiAssign) node()  {}
func (n UnaryExpr) node()    {}
func (n BinaryExpr) node()   {}
func (n SpellCall) node()    {}
func (n Block) node()        {}
func (n GuardedStmt) node()  {}
func (n GuardedBlock) node() {}
func (n When) node()         {}

func (n Ident) assignee()     {}
func (n AnonIdent) assignee() {}

func (n Ident) expr()      {}
func (n AnonIdent) expr()  {}
func (n BoolLit) expr()    {}
func (n NumLit) expr()     {}
func (n StrLit) expr()     {}
func (n UnaryExpr) expr()  {}
func (n BinaryExpr) expr() {}
func (n SpellCall) expr()  {}

func (n SpellCall) multiExpr() {}

func (n BoolLit) literal() {}
func (n NumLit) literal()  {}
func (n StrLit) literal()  {}

func (n SingleAssign) stat() {}
func (n AsymAssign) stat()   {}
func (n MultiAssign) stat()  {}
func (n SpellCall) stat()    {}
func (n Block) stat()        {}
func (n GuardedStmt) stat()  {}
func (n GuardedBlock) stat() {}
func (n When) stat()         {}

func (n GuardedStmt) guard()  {}
func (n GuardedBlock) guard() {}

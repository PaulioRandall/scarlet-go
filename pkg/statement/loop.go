package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Loop struct {
	Open      Token
	IndexId   Token
	InitIndex Expression
	Guard     Guard
}

func (l Loop) Token() Token {
	return l.Open
}

func (l Loop) String(i int) string {

	var s str

	s.indent(i).
		append("[Loop] ").
		appendTk(l.Open)

	s.newline().
		indent(i + 1).
		append("Index:")

	s.newline().
		indent(i + 2).
		appendTk(l.IndexId)

	s.newline().
		indent(i + 1).
		append("IndexInit:")

	s.newline().
		append(l.InitIndex.String(i + 2))

	s.newline().
		indent(i + 1).
		append("Guard:")

	s.newline().
		append(l.Guard.String(i + 2))

	return s.String()
}

type ForEach struct {
	Open    Token
	IndexId Token
	ValueId Token
	MoreId  Token
	List    Expression
	Block   Block
}

func (f ForEach) Token() Token {
	return f.Open
}

func (f ForEach) String(i int) string {

	var s str

	s.indent(i).
		append("[Loop] ").
		appendTk(f.Open)

	s.newline().
		indent(i + 1).
		append("Index: ").
		appendTk(f.IndexId)

	s.newline().
		indent(i + 1).
		append("Value: ").
		appendTk(f.ValueId)

	s.newline().
		indent(i + 1).
		append("More: ").
		appendTk(f.MoreId)

	s.newline().
		indent(i + 1).
		append("List: ")

	s.newline().
		append(f.List.String(i + 2))

	s.newline().
		indent(i + 1).
		append("Block:")

	s.newline().
		append(f.Block.String(i + 2))

	return s.String()
}

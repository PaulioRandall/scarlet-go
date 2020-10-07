package parser

type context struct {
	LexemeIterator
	parent *context
}

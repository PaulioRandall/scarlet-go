package formatter

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func comment(itr lexeme.View) bool {
	return itr.Curr().Tok == lexeme.COMMENT
}

func alignComments(itr *lexeme.Iterator) {
	for itr.JumpToNext(comment) {
		if itr.Curr().Col != 0 {
			alignCommentGroup(itr)
		}
	}
}

func alignCommentGroup(itr *lexeme.Iterator) {

	prev := itr.Curr()
	n := 0

	for itr.JumpToNext(comment) &&
		itr.Curr().Col != 0 &&
		itr.Curr().Line == prev.Line+1 {

		n++
		prev = itr.Curr()
	}

	itr.JumpToPrev(comment)
	if n == 0 {
		return
	}

	maxCol := findMaxCommentCol(itr, n)
	levelComments(itr, n, maxCol)
}

func findMaxCommentCol(itr *lexeme.Iterator, groupSize int) (maxCol int) {
	for i := 0; i <= groupSize; i++ {

		if itr.Curr().Col > maxCol {
			maxCol = itr.Curr().Col
		}

		if i != groupSize {
			itr.JumpToPrev(comment)
		}
	}
	return
}

func levelComments(itr *lexeme.Iterator, groupSize, maxCol int) {

	for i := 0; i <= groupSize; i++ {

		col := itr.Curr().Col
		if col < maxCol {
			itr.Prepend(&lexeme.Lexeme{
				Tok: lexeme.WHITESPACE,
				Raw: strings.Repeat(" ", maxCol-col),
			})
		}

		if i != groupSize {
			itr.JumpToNext(comment)
		}
	}
}

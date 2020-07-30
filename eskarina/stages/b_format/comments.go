package format

/*
import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func alignComments(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())

	for itr.Next() {

	}

	return itr.ToContainer()
}

func alignComment(itr Iterator, col int) {

	if itr.Curr()

}

// Align all comments on preceeding lines, only if their col index is 1 or more
func alignPrevComments(itr Iterator, col int) {

	offset := 0

	jumpToPrevLine := func() bool {

		for itr.Prev() {
			offset++

			if itr.Curr().Tok == lexeme.NEWLINE {
				itr.Prev()
				offset++
				return true
			}
		}

		return false
	}

	for jumpToPrevLine() {

		if itr.Curr().Col > 0 && itr.Curr().Tok == lexeme.COMMENT {
			prependCommentSpace(itr, col)
			itr.Prev()
			offset++
		}
	}

	for i := offset; i > 0; i-- {
		itr.Next()
	}
}

func prependCommentSpace(itr Iterator, col int) {

	size := col - itr.Curr().Col

	itr.Prepend(&lexeme.Lexeme{
		Tok:  lexeme.WHITESPACE,
		Raw:  strings.Repeat(" ", size),
		Line: itr.Curr().Line,
		Col:  itr.Curr().Col,
	})
}
*/

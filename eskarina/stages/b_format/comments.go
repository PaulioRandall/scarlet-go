package format

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type line struct {
	con     *lexeme.Container
	comment bool
}

func alignComments(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())

	lines := splitLines(itr)
	markCommentLines(lines)

	for start, ok := 0, true; ok; {
		var begin, end int
		begin, end, ok = nextCommentGroup(lines, start)

		if ok {
			alignCommentLexemes(lines, begin, end)
			start = end
		}
	}

	// 4. Join lines back together

	return itr.ToContainer()
}

func splitLines(itr Iterator) []line {

	var r []line

	for itr.Next() {

		if itr.Curr().Tok == lexeme.NEWLINE {
			itr.Next()

			l := line{
				con: itr.Split(),
			}

			r = append(r, l)
		}
	}

	return r
}

func markCommentLines(lines []line) {
	for _, l := range lines {
		itr := l.con.ToIterator()

		for itr.Next() {
			if itr.Before() != nil && itr.Curr().Tok == lexeme.COMMENT {
				l.comment = true
			}
		}

		l.con = itr.ToContainer()
	}
}

func nextCommentGroup(lines []line, start int) (begin, end int, found bool) {
	// TODO
	return
}

func alignCommentLexemes(lines []line, begin, end int) {
	// 		Find the comment with the greatest col index
	//    Insert whitespace before comments in each line to match the greatest
}

package format

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type line struct {
	con     *lexeme.Container
	comment int
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

	return joinLines(lines)
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

		for i := 0; itr.Next(); i++ {
			if itr.Curr().Tok == lexeme.COMMENT {
				l.comment = i
				break
			}
		}

		l.con = itr.ToContainer()
	}
}

func nextCommentGroup(lines []line, start int) (begin, end int, found bool) {

	// Find the group beginning line
	for i := start; i < len(lines); i++ {
		if lines[i].comment > 0 {
			found = true
			begin = i
			break
		}
	}

	if !found {
		return
	}

	// Find the group ending line
	for i := begin + 1; i < len(lines); i++ {
		if lines[i].comment < 1 {
			end = i
			break
		}
	}

	// If no end found, then all remain lines have comments
	// SANITY CHECK, should never heappen?
	if end == 0 {
		end = len(lines)
	}

	return
}

func alignCommentLexemes(lines []line, begin, end int) {
	// 		Find the comment with the greatest col index
	//    Insert whitespace before comments in each line to match the greatest
}

func joinLines(lines []line) *lexeme.Container {
	// TODO
	return nil
}

package spellbook

import (
	"sort"
	"strconv"
	"strings"

	"github.com/PaulioRandall/scarlet-go/manual"
)

func init() {
	manual.Register("spells", spellSummaries)
}

type SpellDoc struct {
	Pattern  string
	Summary  string
	Examples []string
}

func FmtSpellDoc(sp SpellDoc) string {

	sb := strings.Builder{}
	lineBreak := func() {
		sb.WriteRune('\n')
		sb.WriteRune('\n')
	}

	sb.WriteString(sp.Pattern)
	lineBreak()

	s := indentLines(sp.Summary)
	sb.WriteString(s)
	lineBreak()

	for i, ex := range sp.Examples {
		if i != 0 {
			lineBreak()
		}

		sb.WriteString("Example ")
		n := strconv.Itoa(i + 1)
		sb.WriteString(n)
		sb.WriteString(":\n")

		ex = indentLines(ex)
		sb.WriteString(ex)
	}

	return sb.String()
}

func spellSummaries() string {

	names := SpellNames()
	sort.Strings(names)

	sb := strings.Builder{}

	for i, v := range names {

		sp := LookUp(v)

		if i != 0 {
			sb.WriteString("\n\n")
		}

		s := fmtSpellSummary(sp.Docs())
		sb.WriteString(s)
	}

	return sb.String()
}

func fmtSpellSummary(sp SpellDoc) string {

	sb := strings.Builder{}

	sb.WriteString(sp.Pattern)
	sb.WriteRune('\n')

	s := indentLines(sp.Summary)
	sb.WriteString(s)

	return sb.String()
}

func indentLines(s string) string {

	sb := strings.Builder{}

	newline := true
	for _, ru := range s {
		if newline {
			newline = false
			sb.WriteRune('\t')
		}

		sb.WriteRune(ru)
		newline = ru == '\n'
	}

	return sb.String()
}

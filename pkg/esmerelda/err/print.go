package err

import (
	"fmt"
	"io"
)

type ScriptFile interface {
	GetLine(index int) (string, bool)
}

func Print(w io.Writer, script ScriptFile, e Err) {

	const (
		LEADING_LINES  = 5
		TRAILING_LINES = 4
	)

	fPrintln(w, "+++Error, redo from start+++")
	fPrintln(w, "[ERROR] %s", e.Error())
}

func calcLeftMargin(s Snippet) (n int) {

	maxLineNum, _ := s.End()

	for maxLineNum != 0 {
		maxLineNum /= 10
		n++
	}
	return
}

/*
func loadScript(f string) []string {

	bs, e := ioutil.ReadFile(f)
	if e != nil {
		panic(e)
	}

	s := string(bs)
	strings.ReplaceAll(s, "\r", "")
	return strings.Split(s, "\n")
}
*/
func fPrintln(w io.Writer, s string, args ...interface{}) {
	fmt.Fprintf(w, s, args...)
	fmt.Fprintln(w)
}

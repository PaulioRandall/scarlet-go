package inst

import (
	"fmt"
	"io"
	"strings"
)

func Print(w io.StringWriter, ins []Instruction) error {
	return printInstructions(w, ins)
}

func printInstructions(w io.StringWriter, ins []Instruction) error {

	if len(ins) == 0 {
		return nil
	}

	codeMax := findPrintPaddings(ins)

	for _, in := range ins {

		code := padBack(codeMax+1, in.Code.String()+",")
		data := dataToString(in.Data)

		e := writeLine(w, code, " ", data)
		if e != nil {
			return e
		}
	}

	return nil
}

func findPrintPaddings(ins []Instruction) (code int) {

	for _, in := range ins {
		codeLen := len(in.Code.String())
		if codeLen > code {
			code = codeLen
		}
	}

	return
}

func padBack(min int, s string) string {
	pad := strings.Repeat(" ", min-len(s))
	return s + pad
}

func dataToString(data interface{}) string {

	switch v := data.(type) {
	case nil:
		return ""

	case string:
		return fmt.Sprintf("%q", v)

	case fmt.Stringer:
		return v.String()
	}

	return fmt.Sprintf("%v", data)
}

func writeLine(w io.StringWriter, strs ...string) error {

	for _, s := range strs {
		_, e := w.WriteString(s)
		if e != nil {
			return e
		}
	}

	_, e := w.WriteString("\n")
	return e
}

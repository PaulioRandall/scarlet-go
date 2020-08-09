package inst

import (
	//"fmt"
	"io"
	//"strconv"
	//"strings"
)

func Print(w io.StringWriter, ins []Instruction) error {
	return printInstructions(w, ins)
}

func printInstructions(w io.StringWriter, ins []Instruction) error {

	if len(ins) == 0 {
		return nil
	}

	for _, in := range ins {

		if _, e := w.WriteString(in.String()); e != nil {
			return e
		}

		if _, e := w.WriteString("\n"); e != nil {
			return e
		}
	}

	return nil
}

package lexeme

import (
	"fmt"
	"strings"
)

func Descend(first *Lexeme, f func(*Lexeme)) {
	for lex := first; lex != nil; lex = lex.Next {
		f(lex)
	}
}

func DiffPrint(left, right *Lexeme) {

	pad := func(s string, n int) string {

		if len(s) >= n {
			return s
		}

		p := n - len(s)
		return s + strings.Repeat(" ", p)
	}

	const padding = 38

	fmt.Print("\n  ")
	fmt.Print(pad("Left", padding))
	fmt.Println("right")

	for left != nil || right != nil {

		var lStr, rStr string

		if left != nil {
			lStr = pad(left.String(), padding)
			left = left.Next
		}

		if right != nil {
			rStr = pad(right.String(), padding)
			right = right.Next
		}

		if lStr != rStr {
			fmt.Print("* ")
		} else {
			fmt.Print("  ")
		}

		fmt.Print(lStr)
		fmt.Println(rStr)
	}

	fmt.Println()
}

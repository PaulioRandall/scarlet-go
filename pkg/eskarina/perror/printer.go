package perror

import (
	"fmt"
	"strings"
)

type Node interface {
	NextNode() Node
	String() string
}

func DiffPrint(left, right Node) {

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
			left = left.NextNode()
		}

		if right != nil {
			rStr = pad(right.String(), padding)
			right = right.NextNode()
		}

		if lStr != rStr {
			fmt.Print("- ")
		} else {
			fmt.Print("+ ")
		}

		fmt.Print(lStr)
		fmt.Println(rStr)
	}

	fmt.Println()
}

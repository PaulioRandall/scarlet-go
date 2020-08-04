package docs

import (
	"fmt"
)

func printVariablesOverview() {

	s := `Variables are symbols used to represent a value which can be changed
through assignment. A value can be simple or complex entities from numbers and
strings of characters to lists and even functions. Variables are extremely
useful as they can hold values we cannot possibly know at the time of coding,
e.g. your search text when using an online search engine.

Variable names:

	Variable names must start with a letter and then any unbroken series of
	letters and underscores providing it is not a keyword.

	Good: 'x'
	Good: 'playerName'
	Good: 'enemy_health'

	Bad:  '_x'               First character must be a letter
	Bad:  'player name'      Spaces are not allowed
	Bad:  'x_123'            Numbers are not allowed

Example usage:

	Curently, the only way to assign or reassign a variable is using the @Set
	spell. This will change in future.

		@Set("variable_name", value)

	Consider the first @Set spell below which takes two arguments. The spell
	sets the variable 'x' to represent the value '42'. We can now use the
	variable where ever we want to use the number '42' such as in the
	@Println spell.	The @Println spell will display the text "6 * 7 = 42"
	in the output console.

		@Set("x", 42)            Sets the variable 'x' to '42'
		@Println("6 * 7 = ", x)  Displays "6 * 7 = 42" in the output terminal

	The code is run from top to bottom so	variables can be reassigned before
	they are used. 

		@Set("x", true)          Sets the variable 'x' to 'true'
		@Set("x", ":)")          Sets the variable 'x' to '":)"'
		@Set("x", 21)            Sets the variable 'x' to '21'
		@Println("3 * 7 = ", x)  Displays "3 * 7 = 21" in the output terminal

Future changes:

	One of the next features will be native assignments as a replacement for
	the @Set spell. This is iconic programming code, easier to type and read,
	and a prerequisite for other language enhancements.

	x := true
	x := 1
	x := "Scarlet"

	Multiple returns will also be allowed to make functions and spells more
	versatile. This means multiple variables can be assigned values on the
	same line.

	x, y := 6, 7
	n, e := @ParseNum("42")

	Once these changes are in place the @Set spell will be removed.
`

	fmt.Println(s)
}

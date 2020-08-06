package types

import (
	//	"fmt"
	"github.com/PaulioRandall/scarlet-go/manual"
)

func init() {
	manual.Register("type", typesOverview)
	manual.Register("types", typesOverview)
	manual.Register("vars", variablesOverview)
	manual.Register("variable", variablesOverview)
	manual.Register("variables", variablesOverview)
}

func typesOverview() string {
	return `
The format and interaction rules of each value in Scarlet is defined
by its data type. There are three intrinsic types but new ones can be added
in a similar fashion to spells.  

	"As soon as you saw people as things to be measured, they didn’t measure up."
		- 'Night Watch' by Terry Pratchett

Usage:

	Specifing types is not required. When a value is assigned to a variable,
	the data type is automatically inferred. However, this means users are
	required to know a variables current data type as different operations
	require values of different types. For example:

		x := 2 * 3              Valid:   multiplying two numbers
		x := 2 * "hello"        Invalid: multiplying a string is nonsensical

Intrinsic data types:

	'Bool'
		Holds one of two possible values, 'true' or 'false'.

	'Num'
		Holds an arbitrary length floating point number. The standard numeric
		operations can be perform on two numbers such as addition and
		multiplication. When an operation or spell requires an integer, the
		integer part of the number passed will be used, i.e. no rounding will
		occur.

	'Str'
		Holds a sequance of UTF-8 characters. Scarlet is very high level and
		does not intrinsically deal with byte data so string manipulation is
		done purely in UTF-8.

Future types:

	I hope to add additional default types such as 'list', 'map', and 'file'
	but these	won't be intrinsic to Scarlet; they can be modified or replaced
	at leisure by any inquisitive programmer.

	'List'
		And its accompanying spells allow a list of values to be stored in an
		ordered manner and operated on through random access or sequentially.

	'Map'
		And its accompanying spells create and store a mapping between
		two values. One will represent the key and the other the mapped value.
		Maps will probably not be ordered but spells might be provided by
		default to return an ordered set of keys.

	'File'
		Will likely only be accessible through special spells. These spells
		will accept a filename along with a function that accepts a 'File'
		variable. Upon invocation the file will be opened and the function
		called with the 'File' as a value which can be used to perform IO.
		Upon function exit or error the file is automatically closed before
		the spell finishes.`
}

func variablesOverview() string {
	return `
Variables are symbols used to represent a value which can be changed
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

	Once these changes are in place the @Set spell will be removed.`
}

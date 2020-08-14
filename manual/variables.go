package manual

func init() {
	Register("var", variablesOverview)
	Register("vars", variablesOverview)
	Register("variable", variablesOverview)
	Register("variables", variablesOverview)
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

	Good:
		'x'
		'playerName'
		'enemy_health'

	Bad:
		'_x'               First character must be a letter
		'player name'      Spaces are not allowed
		'x_123'            Numbers are not allowed

Example usage:

	Assignments are performed in classic imperative fashion allowing for
	multiple assignments using a single assignment operator.
	
		alive        := true
		playerName   := "bob"
		x, y, answer := 6, 7, 42

	Consider the first assignment below which sets the variable 'x' to
	represent the value '42'. We can now use the variable where ever we
	want to use the number '42' such as in the @Println spell. The @Println
	spell will display the text "6 * 7 = 42" in the output console.

		x := 42                  Sets the variable 'x' to '42'
		@Println("6 * 7 = ", x)  Displays "6 * 7 = 42" in the output terminal

	The code is run from top to bottom so	variables can be reassigned before
	they are used. 

		x := true                Sets the variable 'x' to 'true'
		x := ":)"                Sets the variable 'x' to '":)"'
		x := 21                  Sets the variable 'x' to '21'
		@Println("3 * 7 = ", x)  Displays "3 * 7 = 21" in the output terminal

Future changes:

	Some of the next features will be native variable deletions and void
	assignment targets to replace the @Del spell. Void assignment targets
	can be used to ignore the result of an expression.

	_ := 1
	x := _
`
}

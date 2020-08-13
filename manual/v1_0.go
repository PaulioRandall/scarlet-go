package manual

func init() {
	Register("v1", v1_0)
	Register("v1.0", v1_0)
}

func v1_0() string {
	return `
Features introduced in version 1.

Comments

	Comments start with the pound symbol '#' and run to the end of the line.
	They can be placed after some code on the same line to provide clear
	context.

	# The answer to life,
	# the universe,
	# and everything.
		
	@Println(42) # Show your working: 6 * 7

Variables, Intrinsic Types, and Assignments

	bool := true         # Boolean values can either be 'true' or 'false'
	num  := 123.456      # All numbers are arbitrary length floating point
	str  := "Scarlet"    # All strings are UTF-8 character sequences

Spells

	Spells are inbuilt functions and are always begin with the at symbol '@'.

	@Print("Scarlet")        # Prints arguments to output stream
	@Println("6 * 7 = ", 42) # Same as @Print but concludes with a linefeed
	@Exit(0)                 # Exits the script with the specified exit code
`
}

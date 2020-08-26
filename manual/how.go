package manual

func init() {
	Register("how", how)
}

func how() string {
	return `
Comments

	Comments are nonsensical remarks that accompany code but are ignored by
	the compiler. Comments start with a '#' hash symbol (or pound for some
	Americans) and ends at the end of the line.

		# Write whatever you want here
		# Put two or more comment lines together for longer descriptions

Variables & Types
	
	Variables are symbols used to represent a value which can be changed
	through assignment. A value can be simple or complex entities from numbers
	and strings of characters to lists and even functions. Variables are
	extremely useful as they can hold values we cannot possibly know at the
	time of coding. Variable names must start with a letter and then any
	unbroken series of letters and underscores providing it is not a keyword.

	Good: 'x'
		    'playerName'
		    'enemy_health'

	Bad: '_x'                 First character must be a letter
		   'player name'        Spaces are not allowed
		   'player123'          Numbers are not allowed

	The format and interaction rules of each value in Scarlet are defined by
	their data type. There are three intrinsic types but new ones can be added
	fairly easily. Specifing types is not required but programmers are
	required to know a variables value data type as different operations limit
	the types that can be used as operands.

	Assignments to variables are performed in classic imperative fashion
	allowing for multiple assignments using a single assignment operator.
	The type information of a value stays solely with the value so variables
	can be reassigned values of a different type.
	
		alive        := true
		playerName   := "bob"
		answer       := 6 * 7
		answer       := "42"

		x := 2 * 3              Valid:   multiplying two numbers
		x := 2 * "hello"        Invalid: multiplying a string is nonsensical

Operations

	Intrinsic data types
		
		Bool: Holds one of two possible values, 'true' or 'false'.
		Num: Holds an arbitrary length floating point number. The standard
		     numeric operations can be perform on two numbers such as addition
		     and multiplication. When an operation or spell requires an integer,
		     the integer part of the number passed will be used, i.e. no
		     rounding will occur.
		Str: Holds a sequance of UTF-8 characters. Scarlet is very high level
		     and does not intrinsically deal with byte data so string
		     manipulation is done purely in UTF-8.

	These are the operations available with their precedence, a higher number
	means greater precedence and those of equal precedence are prioritised by
	first come first computed.
	
	(6) Num  *   Num
	    Num  /   Num
	    Num  %   Num
	(5) Num  +   Num
	    Num  -   Num
	(4) Num  <   Num
	    Num  >   Num
	    Num  <=  Num
	    Num  >=  Num
	(3) Any  ==  Any
	    Any  !=  Any
	(2) Bool && Bool
	(1) Bool || Bool

Guards

	Guards are used to provide conditional code execution such as printing
	a number only if a specific condition is meet.

		# Code within the curly brackets is only executed if 'x' is greater
		# than zero.
		[x > 0] {
			... # Some conditional code
		}

Loops

	Loops (while) are guards that are repeated until the guard condition
	is false.

		# A simple example that will only loop once
		exit := false
		loop [exit] {
			exit := true
		}

		# This example loops 5 times printing the number held by 'i' on each
		# iteration.
		i := 1
		loop [i < 6] {
			@Print(i, " ")
		  i := i - 1
		}

Spells

	Spells are the central concept on which Scarlet was built. A less
	glamorous name would be 'inbuilt functions', but where's the fun in that.

	"Any sufficiently advanced technology is indistinguishable from magic"
		- [One of] 'Clarke's three laws' by Arthur C Clarke

	Spells are always prefixed with an at sign '@' followed by their name
	and accept arguments in the same manner as iconic C-style functions.
	Unlike variable names, spell names may contain dots '.' to mimic
	namespaces. This can make spells more readable, better convey their
	usage, and are easier to mass search-and-replace. A registered spell
	name may have as many namespace segments as the coder likes but they
	should strive to create names that are short and meaningful.

	Usage: @spell_name([argument...])

		@Println("6 * 7 = ", 6 * 7)
		@Exit(0)`
}

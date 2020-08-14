package manual

func init() {
	Register("type", typesOverview)
	Register("types", typesOverview)
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

Data types:

	'Bool' (Intrinsic)
		Holds one of two possible values, 'true' or 'false'.

	'Num' (Intrinsic)
		Holds an arbitrary length floating point number. The standard numeric
		operations can be perform on two numbers such as addition and
		multiplication. When an operation or spell requires an integer, the
		integer part of the number passed will be used, i.e. no rounding will
		occur.

	'Str' (Intrinsic)
		Holds a sequance of UTF-8 characters. Scarlet is very high level and
		does not intrinsically deal with byte data so string manipulation is
		done purely in UTF-8.

Operations:

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

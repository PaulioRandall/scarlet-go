package manual

func init() {
	Register("future", future)
}

func future() string {
	return `
Assignments

	Some of the next features will be native variable deletions and void
	assignment targets to replace the @Del spell. Void assignment targets
	can be used to ignore the result of an expression.

		_ := 1
		x := _

Types

	I hope to add additional default types such as 'list', 'map', and 'file'
	but these	won't be intrinsic to Scarlet; they can be modified or replaced
	at leisure by any inquisitive programmer.

		'List' and its accompanying spells allow a list of values to be stored
		in an ordered manner and operated on through random access or
		sequentially.

		'Map' and its accompanying spells create and store a mapping between
		two values. One will represent the key and the other the mapped value.
		Maps will probably not be ordered but spells might be provided by
		default to return an ordered set of keys.

		'File' will likely only be accessible through special spells. These
		spells will accept a filename along with a function that accepts a
		'File' variable. Upon invocation the file will be opened and the
		function called with the 'File' as a value which can be used to
		perform IO. Upon function exit or error the file is automatically
		closed before the spell terminates.

Guards

	Guards will probably allow inline expressions since ease and conciseness
	are good properties of scripting tool. However, the inline expression
	must appear on the same line as the guard condition.

		[x == 0] @Println("x is 0")

Loops

	Loops will likely get an optional initialiser assignment before the
	loop condition where a variable can be initialised and only accessible
	within the loop.

		loop i := 0 [i < 5] {
			i := i + 1
		}

Incrementors & Decrmenetors

	Number increment and decrement operations are very probable, however,
	they will only be allowed as statements. This is because usage within
	expressions produces difficult to read code and subtle errors that are
	hard to debug.

		i++
		i--

Spells

	Spells will be able to return multiple values soon. The results being
	assignable to variables.

		x := @Len(s)

	I hope to add blocks as parameters to enable code to be run with the
	callers scope and variables. Here is an example spell with block parameter:

		@If(true, {
			x := 1
		})

	In future I hope to add some very common but completely removable and
	modifable spells to get users started. However, many of these depend on
	language features not yet implemented:

		x := @Args()                    Get the program arguments
		x := @Exists("variable_name")   Does a variable exist?
		x := @Len(value)                Find the length of lengthed value
		x := @Str(value)                Stringify a value of any type
		@Panic(exitCode, message)       Exit the script with an error message
		e := @Catch({                   Catch any panics and return as an error
			...
		})

		@str.                           'Str' spells
		@list.                          'List' type & spells
		@map.                           'Map' type & spells
		@fmt.                           'Template' type & spells
		@io.                            Basic input and output spells`
}

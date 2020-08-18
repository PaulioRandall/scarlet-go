package manual

func init() {
	Register("version", versionDocs)
}

func versionDocs() string {
	return `
Version 1 (Upcoming)

	- Comments
	- Variables
	- Types
		- Intrinsic Types
	- Assignments
	- Expressions
		- Arithmetic
		- Logical
		- Relational
	- Guards
	- Loops (While)
	- Spells
		- @Exit(exitcode)
		- @Print(values...)
		- @Println(values...)
	- Manual`
}

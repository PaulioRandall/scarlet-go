package manual

func init() {
	Register("chapters", chapterDocs)
}

func chapterDocs() string {
	return `
Chapter 1

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
		- @FmtScroll(scroll_file)
	- Manual`
}

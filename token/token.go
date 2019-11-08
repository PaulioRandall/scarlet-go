package token

// Token represents a grammer token within a source file.
type Token struct {
	value string
	kind  Kind
	line  int
	start int
	end   int
}

// ScanToken is a recursive descent function that returns the next token
// followed by the callable tail function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (Token, ScanToken)

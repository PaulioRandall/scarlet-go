package token

// Token represents a grammer token within a source file.
type Token struct {
	Value string // The value of the token within the source code
	Kind  Kind   // The type of the token
	Line  int    // The line in the source code of the token
	Start int    // The index of the first character in the source code
	End   int    // The index after the last character within the source code
}

// ScanToken is a recursive descent function that returns the next token
// followed by the callable tail function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (Token, ScanToken)

package formatter

import (
	"os"
	//		"github.com/PaulioRandall/scarlet-go/eskarina/stages/a_scanner"
)

func FormatAll(filename string) error {

	f, e := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
	if e != nil {
		return e
	}

	// TODO: Catch any panics and return them as an error
	defer f.Close()

	return format(f)
}

func format(file *os.File) error {

	head, e := scanLines(file)
	if e != nil {
		return e
	}

	head = trimLines(head)
	head = removeUselessSpace(head)
	head = removeUselessLines(head)
	head = formaliseLineEndings(head)
	head = insertSpaces(head)
	head = insertIndentation(head)
	head = updateTokenPositions(head)
	head = alignComments(head)

	return updateFile(file, head)
}

func scanLines(file *os.File) (*line, error) {
	// Scan each line including the line ending
	// Put the line into the scanner
	// Return a new *line
	return nil, nil
}

func trimLines(head *line) *line {
	//  Remove leading and trailing whitespace from the line
	return head
}

func removeUselessSpace(head *line) *line {
	//  Remove whitespace where it's not needed
	return head
}

func removeUselessLines(head *line) *line {
	// Remove successive empty lines
	return head
}

func formaliseLineEndings(head *line) *line {
	// Identify the line ending of the first line
	//	Update all line endings that do not conform with the first
	return head
}

func insertSpaces(head *line) *line {
	// Insert whitespace in each line
	//  After a comma etc
	return head
}

func insertIndentation(head *line) *line {
	//  Add nesting indentation to the begining of each line
	return head
}

func updateTokenPositions(head *line) *line {
	// Update all line and column numbers
	return head
}

func alignComments(head *line) *line {
	// Identify code chunks (chunks of lines between empty lines)
	// Align all comments within a chunk
	return nil
}

func updateFile(file *os.File, head *line) error {
	// Stringify each line
	// Compare it to the relevant line in the file
	// If ANY line does not match
	//	Rewrite the whole file
	return nil
}

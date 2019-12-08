package parser

// Context represents a specific scope within a script.
// E.g.
// - Root of the script file
// - Inside a function body `F`
// - Inside a match block `MATCH`
// - etc
type Context interface {
}

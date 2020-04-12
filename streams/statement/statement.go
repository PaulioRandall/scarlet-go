// statement package was created to separate the concern of grouping tokens into
// statements from the parsing of those statements. The API consumes a
// TokenStream to produce a StatementStream using the production rules; a
// stream of UnparsedStatements where each is made up of some tokens and
// possible some more UnparsedStatements.
//
// Key decisions: N/A
//
// This package does not parse the statements, just groups tokens together so
// a parser can more easily parse them.
package statement

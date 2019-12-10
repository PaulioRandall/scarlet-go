# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## Versioning

This project adheres to a form of [Calendar Versioning](https://calver.org/).

> "YYYY.MM.DD" { "-MOD" }

- `YYYY`: Year of the version.
- `MM`: Month of the version.
- `DD`: Day of the version.
- `-MOD`: Multiple optional modifiers that may specify lower level version numbers and tags.

#### Examples

> 2019.11.05
> MAYBE-2019.11.05
> MAYBE-2019.11.05-001

## Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [TODO]
### [Add]
- Function call parsing to parser.
- ID assignment parsing to parser.
- Spell invocation parsing to parser.
- Function assignment parsing to parser.
- Function definition parsing to parser.
- Parser that parses scanned tokens

## [Next Milestone]
### Added
- Evaluator, a wrapper for strimmer that removes tags from string literals.
- Strimmer, a wrapper for scanner that removes whitespace.
- Function call scanning to scanner.
- String literal scanning to scanner.
- ID assignment scanning to scanner.
- Spell invocation scanning to scanner.
- Function assignments scanning to scanner.
- Function definition scanning to scanner.
- Scanner that scans tokens from source code.
- WSN (EBNF) formal grammer.
- This changelog.

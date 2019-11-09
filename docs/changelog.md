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
> 2019.11.05-MAYBE
> 2019.11.05-001-MAYBE

#### Modifiers

- `-MAYBE`: is a candidate for the next stage, e.g. release candidate

## Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Next Milestone]
### Features
- Scan, parse, and run the spell `@Print(STR, BOOL)` which prints the input STR and optionally a new line if the BOOL is TRUE.
- Scan, parse, and run procedures with no input.
- Scan a scarlet source code file.
- WSN (EBNF) formal grammer.
- This changelog.

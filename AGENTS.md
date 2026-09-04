# AGENTS.md - Guidance for AI Agents

## Critical Rules

- **Follow ASD-STE100 writing guidelines** for your prose and documentation.
- **NEVER push directly to `master`** - Any change must be done through a
  patch with a new branch. Create a feature branch, commit your changes,
  and open a PR.
- **NEVER use `/` in branch names** - Use `-` or `_` as separators instead.
- **Keep PRs small and reviewable.** One branch and one PR per logical
  unit of work.
- **Add a CHANGELOG entry for every PR, in its own commit.** The changelog
  entry MUST always be a separate commit from the code it describes.
- **When cutting a tag/release, the CHANGELOG MUST be updated in the same
  PR.** Move the entry from `[Unreleased]` to the version heading.
- **Record the model in the commit message and the PR.** A trailing line
  such as `Model: <provider>/<model-id>`.
- **Do not reference plan or task numbers in commits, PRs, or the changelog.**
- **Verify the full CI pipeline is green before considering a branch done.**

## Security

- **NEVER add any new dependency or update a dependency** without the
  explicit consent of the prompt engineer.

## Key Architectural Patterns

### Lexer/Parser

The library parses YQL (Yet Another Query Language) strings into
Elasticsearch DSL queries. The architecture:

- `const.go` - Token types, condition types, match types, special runes
- `lexer.go` - The `Lexer` struct and `Parse()` entry point
- `states.go` - State machine functions (`lexText`, `lexTerm`, `lexQueryGroup`)
- `*_string.go` - Stringer-generated `String()` methods (build artifacts)

The parser supports:
- Field-value terms: `field:value`
- Grouping: `(term1 OR term2)`
- Negation: `-field:value`
- Range: `field>value`, `field<value`
- Match types: `~regex`, `=keyword`, `"phrase"`, bare terms
- Nested fields via `WithNestedPaths`
- Field aliases via `WithFieldMapping`

## Dependencies

- `github.com/olivere/elastic/v7` - Elasticsearch client (query builder)

## Development Commands

```bash
make help          # Show available targets
make vet           # Run go vet
make test          # Run tests
make lint          # Run golangci-lint
make vuln          # Run govulncheck
make check-format  # Check Go formatting
make check-changelog # Check changelog commit references
```
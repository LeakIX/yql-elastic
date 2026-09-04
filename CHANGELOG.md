# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Dependabot configuration for Go module and GitHub Actions updates ([8b68f39])
- zizmor linting of GitHub Actions workflows in CI ([8b68f39])

### Changed

- Replaced the internal CI configuration with a GitHub Actions workflow ([8b68f39])
- Fixed staticcheck findings in the lexer ([8b68f39])
- Pinned all GitHub Actions to commit hashes and scoped workflow
  permissions ([8b68f39])

## [0.1.1]

### Changed

- Go directive bumped from 1.14 to 1.26, gofmt the codebase, and add
  Makefile, GitHub Actions workflow, .golangci.yml, and AGENTS.md ([8b68f39])

### Fixed

- Use match phrase for quoted terms ([1add022])
- Default fields use match query ([5aa9bbf])
- Quote state reset when already in quote state ([04e6329])
- Unquoting of lexer values ([5661172])
- Nested support for default fields ([1b979bb])

## [0.1.0]

### Added

- Initial release. YQL lexer/parser to Elasticsearch DSL
  ([b9add0f])

<!-- Commit links -->
[b9add0f]: https://github.com/LeakIX/yql-elastic/commit/b9add0f
[1b979bb]: https://github.com/LeakIX/yql-elastic/commit/1b979bb
[5661172]: https://github.com/LeakIX/yql-elastic/commit/5661172
[04e6329]: https://github.com/LeakIX/yql-elastic/commit/04e6329
[5aa9bbf]: https://github.com/LeakIX/yql-elastic/commit/5aa9bbf
[1add022]: https://github.com/LeakIX/yql-elastic/commit/1add022
[8b68f39]: https://github.com/LeakIX/yql-elastic/commit/8b68f39
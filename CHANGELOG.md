# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [0.1.1]

### Added

- Field callback with error return ([8887a10], [def9ffb])

### Changed

- Go directive bumped from 1.14 to 1.26 ([0c42931])
- gofmt the entire codebase, remove unreachable code, extract
  `.keyword` suffix into a named constant ([29fa7cf])
- Added Makefile, .drone.yml, .golangci.yml, AGENTS.md

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
[8887a10]: https://github.com/LeakIX/yql-elastic/commit/8887a10
[def9ffb]: https://github.com/LeakIX/yql-elastic/commit/def9ffb
[29fa7cf]: https://github.com/LeakIX/yql-elastic/commit/29fa7cf
[0c42931]: https://github.com/LeakIX/yql-elastic/commit/0c42931
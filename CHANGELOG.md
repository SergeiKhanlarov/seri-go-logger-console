# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security

## [v0.1.0] - 2025-11-30

### Added
- Initial release of console provider for `seri-go-logger`
- `ConsoleFormatter` with ANSI color support for different log levels
- `consoleProvider` implementing `sglogger.LoggerProvider` interface
- Color-coded log output with timestamp and caller information
- Support for structured logging with key-value fields
- Smart caller detection that skips logrus and logger frames
- Configurable log levels with `ProviderConfig`
- `NewConsoleProvider` constructor with formatter support
- `NewConsoleFormatter` for creating default console formatter
- Comprehensive documentation and examples

### Features
- **Color Scheme**:
  - Debug: Cyan (`\033[36m`)
  - Info: Green (`\033[32m`)
  - Warn: Yellow (`\033[33m`)
  - Error/Fatal: Red (`\033[31m`)
  - Default: White (`\033[37m`)

- **Output Format**: `TIMESTAMP [COLORED_LEVEL] FILE(LINE) - MESSAGE [key=value]`
- **Level Filtering**: Efficient level-based message filtering
- **Resource Management**: No external resources to clean up

### Technical Details
- Built on top of `logrus` for robust logging backend
- Implements full `sglogger.LoggerProvider` interface
- Supports context-aware logging
- Thread-safe implementation

## [v0.1.1] - 2025-11-30

### Fixed
- downgraded required go version

## [v0.1.2] - 2025-11-30

### Fixed
- cannot refer to unexported field level in struct literal of type sgloggerconsole.ProviderConfig
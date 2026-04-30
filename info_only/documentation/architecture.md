# Architecture Overview

## Project Structure
```
adbjson/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command and global flags
│   ├── devices.go         # devices command
│   └── ...
├── internal/
│   ├── adb/               # ADB executor
│   ├── config/            # Configuration management
│   ├── errors/            # Structured error types
│   ├── formatter/         # Output formatters (JSON, YAML)
│   ├── logger/            # Structured logging
│   ├── model/             # Data models
│   ├── parser/            # ADB output parsers
│   └── sanitize/          # Input sanitization
├── collection/            # Individual command documentation
├── docs/                  # Architecture and best practices docs
└── platform-tools/        # Bundled ADB binaries
```

## Core Components

### 1. ADB Executor (`internal/adb/`)
Executes ADB commands and returns raw output.

### 2. Parsers (`internal/parser/`)
Parse raw ADB output into structured data models.
- All parsers implement the `Parser[T]` interface
- Include validation logic
- Version tracking for compatibility

### 3. Formatters (`internal/formatter/`)
Format structured data into various output formats.
- JSON (default)
- YAML
- Extensible for future formats

### 4. Logger (`internal/logger/`)
Structured JSON logging to stderr.
- Log levels: DEBUG, INFO, WARN, ERROR
- Contextual information
- Enabled via `--debug` flag or `ADBJSON_DEBUG` env var

### 5. Error Handling (`internal/errors/`)
Structured error types for consistent error handling.
- Categorized error types
- Rich context
- Error unwrapping support

### 6. Configuration (`internal/config/`)
Environment variable-based configuration.
- ADB path
- Log level
- Output format
- Timeout settings

### 7. Security (`internal/sanitize/`)
Input sanitization to prevent security vulnerabilities.
- Command injection prevention
- Path traversal protection
- Control character removal

## Data Flow

```
User Input → CLI Command → ADB Executor → Raw Output → Parser → Validation → Formatter → Output
```

## Best Practices Implemented

### ✅ Structured Logging
- JSON-formatted logs
- Multiple log levels
- Contextual information

### ✅ Parser Interface Standardization
- Generic `Parser[T]` interface
- Built-in validation
- Version tracking

### ✅ Error Handling Strategy
- Categorized error types
- Rich context
- Consistent error patterns

### ✅ Configuration Management
- Environment variables
- Default values
- Type-safe configuration

### ✅ Testing Framework
- Table-driven tests
- Parser tests
- Formatter tests
- Security tests

### ✅ Input Sanitization
- Command injection prevention
- Path traversal protection
- Control character removal

### ✅ Multiple Output Formats
- JSON (default)
- YAML
- Extensible architecture

## Design Principles

1. **Type Safety**: Strong typing throughout the codebase
2. **Extensibility**: Easy to add new commands, parsers, and formatters
3. **Testability**: Comprehensive test coverage
4. **Security**: Input sanitization and validation
5. **Observability**: Structured logging and error handling
6. **Performance**: Minimal overhead, efficient parsing

## Future Enhancements

- Schema validation for command outputs
- Plugin architecture for third-party commands
- Connection pooling for ADB
- Mock ADB server for testing
- Property-based testing
- Benchmark tests

# Parser Interface Standardization

## Overview
All ADB output parsers must implement the standardized `Parser[T]` interface to ensure consistency across the codebase.

## Parser Interface

```go
type Parser[T any] interface {
    // Parse converts raw ADB output into a structured type
    Parse(output string) (T, error)
    
    // Validate checks if the parsed result is valid
    Validate(result T) error
    
    // Name returns the parser name for identification
    Name() string
    
    // Version returns the parser version for compatibility tracking
    Version() string
}
```

## BaseParser
All parsers should embed `BaseParser` to get common functionality:

```go
type DevicesParser struct {
    *BaseParser
}

func NewDevicesParser() *DevicesParser {
    return &DevicesParser{
        BaseParser: NewBaseParser("devices", "1.0.0"),
    }
}
```

## Example Implementation

```go
package parser

import (
    "adbjson/internal/model"
    "strings"
)

type DevicesParser struct {
    *BaseParser
}

func NewDevicesParser() *DevicesParser {
    return &DevicesParser{
        BaseParser: NewBaseParser("devices", "1.0.0"),
    }
}

func (p *DevicesParser) Parse(output string) (*model.DevicesResponse, error) {
    // Parse logic here
    return result, nil
}

func (p *DevicesParser) Validate(result *model.DevicesResponse) error {
    // Validation logic here
    return nil
}
```

## Benefits
- **Type Safety**: Generic interface ensures type correctness
- **Versioning**: Track parser versions for compatibility
- **Validation**: Built-in validation for all parsers
- **Testing**: Standard interface makes testing easier
- **Extensibility**: Easy to add new parsers following the pattern

## Parser Versioning
- Use semantic versioning (MAJOR.MINOR.PATCH)
- Increment MAJOR for breaking changes
- Increment MINOR for new features
- Increment PATCH for bug fixes

## Validation Guidelines
- Return error for invalid data
- Use `NoOpValidate` for parsers that don't need validation
- Validate required fields
- Check data ranges and formats

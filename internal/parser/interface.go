package parser

import "context"

// Parser defines the standard interface for all ADB output parsers
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

// BaseParser provides common functionality for all parsers
type BaseParser struct {
	parserName    string
	parserVersion string
}

// NewBaseParser creates a new base parser
func NewBaseParser(name, version string) *BaseParser {
	return &BaseParser{
		parserName:    name,
		parserVersion: version,
	}
}

// Name returns the parser name
func (p *BaseParser) Name() string {
	return p.parserName
}

// Version returns the parser version
func (p *BaseParser) Version() string {
	return p.parserVersion
}

// NoOpValidate is a default validation that always passes
// Can be used by parsers that don't need custom validation
func NoOpValidate[T any](result T) error {
	return nil
}

// ParseWithContext allows parsing with context for cancellation
type ParserWithContext[T any] interface {
	Parser[T]
	ParseWithContext(ctx context.Context, output string) (T, error)
}

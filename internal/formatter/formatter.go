package formatter

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

// OutputFormat represents the output format type
type OutputFormat string

const (
	JSONFormat OutputFormat = "json"
	YAMLFormat OutputFormat = "yaml"
)

// Formatter interface for output formatting
type Formatter interface {
	Format(data interface{}, compact bool) ([]byte, error)
}

// JSONFormatter formats output as JSON
type JSONFormatter struct{}

// NewJSONFormatter creates a new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// Format formats data as JSON
func (f *JSONFormatter) Format(data interface{}, compact bool) ([]byte, error) {
	if compact {
		return json.Marshal(data)
	}
	return json.MarshalIndent(data, "", "  ")
}

// YAMLFormatter formats output as YAML
type YAMLFormatter struct{}

// NewYAMLFormatter creates a new YAML formatter
func NewYAMLFormatter() *YAMLFormatter {
	return &YAMLFormatter{}
}

// Format formats data as YAML
func (f *YAMLFormatter) Format(data interface{}, compact bool) ([]byte, error) {
	return yaml.Marshal(data)
}

// GetFormatter returns the appropriate formatter for the given format
func GetFormatter(format OutputFormat) Formatter {
	switch format {
	case YAMLFormat:
		return NewYAMLFormatter()
	case JSONFormat:
		fallthrough
	default:
		return NewJSONFormatter()
	}
}

// ParseFormat parses a string to OutputFormat
func ParseFormat(format string) OutputFormat {
	switch format {
	case "yaml", "yml":
		return YAMLFormat
	case "json":
		fallthrough
	default:
		return JSONFormat
	}
}

// FormatOutput formats data using the specified format
func FormatOutput(data interface{}, format OutputFormat, compact bool) ([]byte, error) {
	formatter := GetFormatter(format)
	return formatter.Format(data, compact)
}

// FormatOutputString formats data and returns as string
func FormatOutputString(data interface{}, format OutputFormat, compact bool) (string, error) {
	bytes, err := FormatOutput(data, format, compact)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// SupportedFormats returns list of supported formats
func SupportedFormats() []OutputFormat {
	return []OutputFormat{JSONFormat, YAMLFormat}
}

// IsValidFormat checks if a format is supported
func IsValidFormat(format string) bool {
	parsed := ParseFormat(format)
	for _, supported := range SupportedFormats() {
		if parsed == supported {
			return true
		}
	}
	return true // ParseFormat defaults to JSON, so always valid
}

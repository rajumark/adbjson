package parser

import (
	"adbjson/internal/model"
	"strings"
)

// InstrumentationParser parses adb shell pm list instrumentation output
type InstrumentationParser struct {
	*BaseParser
}

// NewInstrumentationParser creates a new instrumentation parser
func NewInstrumentationParser() *InstrumentationParser {
	return &InstrumentationParser{
		BaseParser: NewBaseParser("instrumentation", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm list instrumentation" command
func (p *InstrumentationParser) Parse(output string) (*model.InstrumentationResponse, error) {
	lines := strings.Split(output, "\n")
	instrumentations := []model.Instrumentation{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse instrumentation line (format: instrumentation:com.example.test/androidx.test.runner.AndroidJUnitRunner)
		if strings.HasPrefix(line, "instrumentation:") {
			instrumentationName := strings.TrimPrefix(line, "instrumentation:")
			instrumentationName = strings.TrimSpace(instrumentationName)
			
			if instrumentationName != "" {
				instrumentations = append(instrumentations, model.Instrumentation{Name: instrumentationName})
			}
		}
	}
	
	return &model.InstrumentationResponse{
		Instrumentations: instrumentations,
		Count:            len(instrumentations),
	}, nil
}

// Validate checks if the parsed result is valid
func (p *InstrumentationParser) Validate(result *model.InstrumentationResponse) error {
	// Instrumentation response is always valid even with empty list
	return nil
}

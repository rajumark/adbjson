package parser

import (
	"adbjson/internal/model"
	"strings"
)

// GetpropParser parses adb shell getprop output
type GetpropParser struct {
	*BaseParser
}

// NewGetpropParser creates a new getprop parser
func NewGetpropParser() *GetpropParser {
	return &GetpropParser{
		BaseParser: NewBaseParser("getprop", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell getprop" command
func (p *GetpropParser) Parse(output string) (*model.PropertiesResponse, error) {
	lines := strings.Split(output, "\n")
	properties := []model.Property{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse property line (format: [key]: [value])
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			
			// Remove brackets from key if present
			key = strings.Trim(key, "[]")
			
			if key != "" {
				properties = append(properties, model.Property{Key: key, Value: value})
			}
		}
	}
	
	return &model.PropertiesResponse{
		Properties: properties,
		Count:      len(properties),
	}, nil
}

// Validate checks if the parsed result is valid
func (p *GetpropParser) Validate(result *model.PropertiesResponse) error {
	// Properties response is always valid even with empty list
	return nil
}

package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DateParser parses adb shell date output
type DateParser struct {
	*BaseParser
}

// NewDateParser creates a new date parser
func NewDateParser() *DateParser {
	return &DateParser{
		BaseParser: NewBaseParser("date", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell date" command
func (p *DateParser) Parse(output string) (*model.DateResponse, error) {
	output = strings.TrimSpace(output)
	
	response := &model.DateResponse{
		DateTime: output,
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DateParser) Validate(result *model.DateResponse) error {
	// Check if datetime was parsed
	if result.DateTime == "" {
		return fmt.Errorf("datetime not found in output")
	}
	
	return nil
}

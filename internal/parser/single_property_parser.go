package parser

import (
	"adbjson/internal/model"
	"strings"
)

// SinglePropertyParser parses single property output from getprop
type SinglePropertyParser struct {
	*BaseParser
}

// NewSinglePropertyParser creates a new single property parser
func NewSinglePropertyParser() *SinglePropertyParser {
	return &SinglePropertyParser{
		BaseParser: NewBaseParser("single_property", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell getprop <property>" command
func (p *SinglePropertyParser) Parse(output string, property string) (*model.SinglePropertyResponse, error) {
	// Trim whitespace from output
	value := strings.TrimSpace(output)
	
	// If output is empty, property doesn't exist
	if value == "" {
		value = ""
	}
	
	response := &model.SinglePropertyResponse{
		Property: property,
		Value:    value,
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *SinglePropertyParser) Validate(result *model.SinglePropertyResponse) error {
	// Single property response is always valid
	return nil
}

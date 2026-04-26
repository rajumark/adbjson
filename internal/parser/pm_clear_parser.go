package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PmClearParser parses adb shell pm clear output
type PmClearParser struct {
	*BaseParser
}

// NewPmClearParser creates a new pm clear parser
func NewPmClearParser() *PmClearParser {
	return &PmClearParser{
		BaseParser: NewBaseParser("pm_clear", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm clear" command
func (p *PmClearParser) Parse(output string) (*model.ClearResponse, error) {
	output = strings.TrimSpace(output)
	
	clearResult := model.ClearResult{
		Success: false,
		Message: output,
	}
	
	// Check for success indicators
	if strings.Contains(output, "Success") {
		clearResult.Success = true
		clearResult.Message = "Package data cleared successfully"
	} else if strings.Contains(output, "Error") || strings.Contains(output, "Unknown package") {
		clearResult.Success = false
		clearResult.Message = "Failed to clear package data: " + output
	}
	
	return &model.ClearResponse{
		ClearResult: clearResult,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *PmClearParser) Validate(result *model.ClearResponse) error {
	// Clear response is always valid
	return nil
}

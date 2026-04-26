package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PmEnableParser parses adb shell pm enable output
type PmEnableParser struct {
	*BaseParser
}

// NewPmEnableParser creates a new pm enable parser
func NewPmEnableParser() *PmEnableParser {
	return &PmEnableParser{
		BaseParser: NewBaseParser("pm_enable", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm enable" command
func (p *PmEnableParser) Parse(output string) (*model.EnableResponse, error) {
	output = strings.TrimSpace(output)
	
	enableResult := model.EnableResult{
		Success: false,
		Message: output,
	}
	
	// Check for success indicators
	if strings.Contains(output, "Package enabled") || strings.Contains(output, "new state: enabled") {
		enableResult.Success = true
		enableResult.Message = "Package enabled successfully"
	} else if strings.Contains(output, "Error") || strings.Contains(output, "Unknown package") {
		enableResult.Success = false
		enableResult.Message = "Failed to enable package: " + output
	}
	
	return &model.EnableResponse{
		EnableResult: enableResult,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *PmEnableParser) Validate(result *model.EnableResponse) error {
	// Enable response is always valid
	return nil
}

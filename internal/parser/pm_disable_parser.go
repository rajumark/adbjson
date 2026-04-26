package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PmDisableParser parses adb shell pm disable output
type PmDisableParser struct {
	*BaseParser
}

// NewPmDisableParser creates a new pm disable parser
func NewPmDisableParser() *PmDisableParser {
	return &PmDisableParser{
		BaseParser: NewBaseParser("pm_disable", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm disable" command
func (p *PmDisableParser) Parse(output string) (*model.DisableResponse, error) {
	output = strings.TrimSpace(output)
	
	disableResult := model.DisableResult{
		Success: false,
		Message: output,
	}
	
	// Check for success indicators
	if strings.Contains(output, "Package disabled") || strings.Contains(output, "new state: disabled") {
		disableResult.Success = true
		disableResult.Message = "Package disabled successfully"
	} else if strings.Contains(output, "Error") || strings.Contains(output, "Unknown package") {
		disableResult.Success = false
		disableResult.Message = "Failed to disable package: " + output
	}
	
	return &model.DisableResponse{
		DisableResult: disableResult,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *PmDisableParser) Validate(result *model.DisableResponse) error {
	// Disable response is always valid
	return nil
}

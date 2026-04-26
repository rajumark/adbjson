package parser

import (
	"adbjson/internal/model"
	"strings"
)

// WmDensityResetParser parses adb shell wm density reset output
type WmDensityResetParser struct {
	*BaseParser
}

// NewWmDensityResetParser creates a new wm density reset parser
func NewWmDensityResetParser() *WmDensityResetParser {
	return &WmDensityResetParser{
		BaseParser: NewBaseParser("wm_density_reset", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell wm density reset" command
func (p *WmDensityResetParser) Parse(output string) (*model.WmDensityResetResponse, error) {
	// Default response - reset commands typically have no output on success
	response := &model.WmDensityResetResponse{
		Success: true,
		Message: strings.TrimSpace(output),
	}
	
	// Check for error messages in output
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for error messages
		if strings.Contains(line, "error") || strings.Contains(line, "failed") || strings.Contains(line, "cannot") {
			response.Success = false
			break
		}
		
		// Check for permission denied
		if strings.Contains(line, "Permission denied") {
			response.Success = false
			break
		}
		
		// Check for invalid operation
		if strings.Contains(line, "Invalid") || strings.Contains(line, "unknown") {
			response.Success = false
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *WmDensityResetParser) Validate(result *model.WmDensityResetResponse) error {
	// WM density reset response is always valid
	return nil
}

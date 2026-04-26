package parser

import (
	"adbjson/internal/model"
	"strings"
)

// WmSizeResetParser parses adb shell wm size reset output
type WmSizeResetParser struct {
	*BaseParser
}

// NewWmSizeResetParser creates a new wm size reset parser
func NewWmSizeResetParser() *WmSizeResetParser {
	return &WmSizeResetParser{
		BaseParser: NewBaseParser("wm_size_reset", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell wm size reset" command
func (p *WmSizeResetParser) Parse(output string) (*model.WmSizeResetResponse, error) {
	// Default response - reset commands typically have no output on success
	response := &model.WmSizeResetResponse{
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
func (p *WmSizeResetParser) Validate(result *model.WmSizeResetResponse) error {
	// WM size reset response is always valid
	return nil
}

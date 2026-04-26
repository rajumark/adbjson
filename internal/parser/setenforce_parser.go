package parser

import (
	"adbjson/internal/model"
	"strings"
)

// SetenforceParser parses adb shell setenforce output
type SetenforceParser struct {
	*BaseParser
}

// NewSetenforceParser creates a new setenforce parser
func NewSetenforceParser() *SetenforceParser {
	return &SetenforceParser{
		BaseParser: NewBaseParser("setenforce", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell setenforce" command
func (p *SetenforceParser) Parse(output string, mode string) (*model.SetenforceResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Default response
	response := &model.SetenforceResponse{
		Success: false,
		Mode:    mode,
		Message: strings.TrimSpace(output),
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for successful setenforce (usually no output on success)
		if line == "" {
			response.Success = true
			response.Message = "SELinux enforcing mode set successfully"
			break
		}
		
		// Check for permission denied
		if strings.Contains(line, "Permission denied") {
			response.Success = false
			response.Message = "Permission denied: requires root access"
			break
		}
		
		// Check for other errors
		if strings.Contains(line, "Couldn't set enforcing status") {
			response.Success = false
			break
		}
		
		// Check for invalid operation
		if strings.Contains(line, "Invalid operation") {
			response.Success = false
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *SetenforceParser) Validate(result *model.SetenforceResponse) error {
	// Setenforce response is always valid
	return nil
}

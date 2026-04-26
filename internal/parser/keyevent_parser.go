package parser

import (
	"adbjson/internal/model"
	"strings"
)

// KeyeventParser parses adb shell input keyevent output
type KeyeventParser struct {
	*BaseParser
}

// NewKeyeventParser creates a new keyevent parser
func NewKeyeventParser() *KeyeventParser {
	return &KeyeventParser{
		BaseParser: NewBaseParser("keyevent", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell input keyevent" command
func (p *KeyeventParser) Parse(output string, keycode string) (*model.KeyeventResponse, error) {
	// Default response - input commands typically have no output on success
	response := &model.KeyeventResponse{
		Success: true,
		Keycode: keycode,
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
		
		// Check for invalid keycode
		if strings.Contains(line, "Invalid key") || strings.Contains(line, "unknown") {
			response.Success = false
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *KeyeventParser) Validate(result *model.KeyeventResponse) error {
	// Keyevent response is always valid
	return nil
}

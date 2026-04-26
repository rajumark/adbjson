package parser

import (
	"adbjson/internal/model"
	"strings"
)

// UnrootParser parses adb unroot output
type UnrootParser struct {
	*BaseParser
}

// NewUnrootParser creates a new unroot parser
func NewUnrootParser() *UnrootParser {
	return &UnrootParser{
		BaseParser: NewBaseParser("unroot", "1.0.0"),
	}
}

// Parse parses the raw output from "adb unroot" command
func (p *UnrootParser) Parse(output string) (*model.UnrootResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Default response
	response := &model.UnrootResponse{
		Success: false,
		Message: strings.TrimSpace(output),
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for successful unroot
		if strings.Contains(line, "restarting adbd as non root") {
			response.Success = true
			break
		}
		
		// Check for already not running as root (this is actually success)
		if strings.Contains(line, "adbd not running as root") {
			response.Success = true
			break
		}
		
		// Check for failure messages
		if strings.Contains(line, "failed") || strings.Contains(line, "error") {
			response.Success = false
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *UnrootParser) Validate(result *model.UnrootResponse) error {
	// Unroot response is always valid
	return nil
}

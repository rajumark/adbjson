package parser

import (
	"adbjson/internal/model"
	"strings"
)

// RootParser parses adb root output
type RootParser struct {
	*BaseParser
}

// NewRootParser creates a new root parser
func NewRootParser() *RootParser {
	return &RootParser{
		BaseParser: NewBaseParser("root", "1.0.0"),
	}
}

// Parse parses the raw output from "adb root" command
func (p *RootParser) Parse(output string) (*model.RootResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Default response
	response := &model.RootResponse{
		Success: false,
		Message: strings.TrimSpace(output),
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check for successful root
		if strings.Contains(line, "restarting adbd as root") {
			response.Success = true
			break
		}
		
		// Check for failed root (production builds)
		if strings.Contains(line, "adbd cannot run as root in production builds") {
			response.Success = false
			break
		}
		
		// Check for other failure messages
		if strings.Contains(line, "failed") || strings.Contains(line, "error") {
			response.Success = false
			break
		}
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *RootParser) Validate(result *model.RootResponse) error {
	// Root response is always valid
	return nil
}

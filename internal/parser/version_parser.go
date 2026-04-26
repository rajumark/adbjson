package parser

import (
	"adbjson/internal/model"
	"strings"
)

// VersionParser parses adb version output
type VersionParser struct{}

// NewVersionParser creates a new version parser
func NewVersionParser() *VersionParser {
	return &VersionParser{}
}

// Parse parses the raw output from "adb version" command
func (p *VersionParser) Parse(output string) (*model.VersionResponse, error) {
	lines := strings.Split(output, "\n")
	
	response := &model.VersionResponse{
		Version: "unknown",
	}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Parse version line: "Android Debug Bridge version 1.0.41"
		if strings.Contains(line, "Android Debug Bridge version") {
			parts := strings.Split(line, "version")
			if len(parts) > 1 {
				response.Version = strings.TrimSpace(parts[1])
			}
		}
		
		// Parse revision if present
		if strings.Contains(line, "Revision") {
			parts := strings.Split(line, "Revision")
			if len(parts) > 1 {
				response.Revision = strings.TrimSpace(parts[1])
			}
		}
	}
	
	return response, nil
}

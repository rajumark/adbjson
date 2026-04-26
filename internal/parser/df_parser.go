package parser

import (
	"adbjson/internal/model"
	"strings"
)

// DfParser parses adb shell df output
type DfParser struct {
	*BaseParser
}

// NewDfParser creates a new df parser
func NewDfParser() *DfParser {
	return &DfParser{
		BaseParser: NewBaseParser("df", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell df" command
func (p *DfParser) Parse(output string) (*model.DfResponse, error) {
	lines := strings.Split(output, "\n")
	filesystems := []model.Filesystem{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Skip header line
		if strings.Contains(line, "Filesystem") && strings.Contains(line, "1K-blocks") {
			continue
		}
		
		// Parse filesystem line
		// Format: Filesystem 1K-blocks Used Available Use% Mounted on
		parts := strings.Fields(line)
		if len(parts) >= 6 {
			// Handle cases where mounted on path contains spaces
			var mountedOn string
			if len(parts) > 6 {
				// Join the remaining parts for mounted on path
				mountedOn = strings.Join(parts[5:], " ")
			} else {
				mountedOn = parts[5]
			}
			
			filesystem := model.Filesystem{
				Filesystem: parts[0],
				Blocks:     parts[1],
				Used:       parts[2],
				Available:  parts[3],
				UsePercent: parts[4],
				MountedOn:  mountedOn,
			}
			filesystems = append(filesystems, filesystem)
		}
	}
	
	response := &model.DfResponse{
		Filesystems: filesystems,
		Count:       len(filesystems),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DfParser) Validate(result *model.DfResponse) error {
	// Df response is always valid
	return nil
}

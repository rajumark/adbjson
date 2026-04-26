package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// LsRootParser parses adb shell ls / output
type LsRootParser struct {
	*BaseParser
}

// NewLsRootParser creates a new ls / parser
func NewLsRootParser() *LsRootParser {
	return &LsRootParser{
		BaseParser: NewBaseParser("ls /", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell ls /" command
func (p *LsRootParser) Parse(output string) (*model.LsRootResponse, error) {
	lines := strings.Split(output, "\n")
	items := []string{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Split by whitespace to get individual items
		parts := strings.Fields(line)
		for _, item := range parts {
			item = strings.TrimSpace(item)
			if item != "" {
				items = append(items, item)
			}
		}
	}
	
	response := &model.LsRootResponse{
		Items: items,
		Count: len(items),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *LsRootParser) Validate(result *model.LsRootResponse) error {
	// Check if items were parsed
	if len(result.Items) == 0 {
		return fmt.Errorf("no items found in root directory")
	}
	
	return nil
}

package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// LsProcParser parses adb shell ls /proc output
type LsProcParser struct {
	*BaseParser
}

// NewLsProcParser creates a new ls /proc parser
func NewLsProcParser() *LsProcParser {
	return &LsProcParser{
		BaseParser: NewBaseParser("ls /proc", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell ls /proc" command
func (p *LsProcParser) Parse(output string) (*model.LsProcResponse, error) {
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
	
	response := &model.LsProcResponse{
		Items: items,
		Count: len(items),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *LsProcParser) Validate(result *model.LsProcResponse) error {
	// Check if items were parsed
	if len(result.Items) == 0 {
		return fmt.Errorf("no items found in /proc directory")
	}
	
	return nil
}

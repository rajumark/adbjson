package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysActivityParser parses adb shell dumpsys activity output
type DumpsysActivityParser struct {
	*BaseParser
}

// NewDumpsysActivityParser creates a new dumpsys activity parser
func NewDumpsysActivityParser() *DumpsysActivityParser {
	return &DumpsysActivityParser{
		BaseParser: NewBaseParser("dumpsys activity", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys activity" command
func (p *DumpsysActivityParser) Parse(output string) (*model.DumpsysActivityResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysActivitySection{}
	
	var currentSection model.DumpsysActivitySection
	var contentLines []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check for section headers (lines in all caps with parentheses)
		if strings.Contains(line, "ACTIVITY MANAGER") && strings.Contains(line, "(") && strings.Contains(line, ")") {
			// Save previous section if exists
			if currentSection.Name != "" {
				currentSection.Content = strings.Join(contentLines, "\n")
				sections = append(sections, currentSection)
				contentLines = []string{}
			}
			
			// Extract section name
			parts := strings.Split(line, "(")
			if len(parts) >= 2 {
				sectionName := strings.TrimSpace(parts[0])
				currentSection = model.DumpsysActivitySection{
					Name:    sectionName,
					Content: "",
				}
			}
		} else if currentSection.Name != "" {
			// Add line to current section content
			contentLines = append(contentLines, line)
		}
	}
	
	// Save last section if exists
	if currentSection.Name != "" {
		currentSection.Content = strings.Join(contentLines, "\n")
		sections = append(sections, currentSection)
	}
	
	response := &model.DumpsysActivityResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysActivityParser) Validate(result *model.DumpsysActivityResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys activity output")
	}
	
	return nil
}

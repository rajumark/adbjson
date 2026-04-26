package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysLocationParser parses adb shell dumpsys location output
type DumpsysLocationParser struct {
	*BaseParser
}

// NewDumpsysLocationParser creates a new dumpsys location parser
func NewDumpsysLocationParser() *DumpsysLocationParser {
	return &DumpsysLocationParser{
		BaseParser: NewBaseParser("dumpsys location", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys location" command
func (p *DumpsysLocationParser) Parse(output string) (*model.DumpsysLocationResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysLocationSection{}
	
	var currentSection model.DumpsysLocationSection
	var contentLines []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check for section headers (lines ending with colon and not empty)
		if strings.HasSuffix(line, ":") && line != ":" {
			// Save previous section if exists
			if currentSection.Name != "" {
				currentSection.Content = strings.Join(contentLines, "\n")
				sections = append(sections, currentSection)
				contentLines = []string{}
			}
			
			// Extract section name (remove trailing colon)
			sectionName := strings.TrimSuffix(line, ":")
			currentSection = model.DumpsysLocationSection{
				Name:    sectionName,
				Content: "",
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
	
	response := &model.DumpsysLocationResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysLocationParser) Validate(result *model.DumpsysLocationResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys location output")
	}
	
	return nil
}

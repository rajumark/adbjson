package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysInputParser parses adb shell dumpsys input output
type DumpsysInputParser struct {
	*BaseParser
}

// NewDumpsysInputParser creates a new dumpsys input parser
func NewDumpsysInputParser() *DumpsysInputParser {
	return &DumpsysInputParser{
		BaseParser: NewBaseParser("dumpsys input", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys input" command
func (p *DumpsysInputParser) Parse(output string) (*model.DumpsysInputResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysInputSection{}
	
	var currentSection model.DumpsysInputSection
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
			currentSection = model.DumpsysInputSection{
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
	
	response := &model.DumpsysInputResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysInputParser) Validate(result *model.DumpsysInputResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys input output")
	}
	
	return nil
}

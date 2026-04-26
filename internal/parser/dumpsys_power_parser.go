package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysPowerParser parses adb shell dumpsys power output
type DumpsysPowerParser struct {
	*BaseParser
}

// NewDumpsysPowerParser creates a new dumpsys power parser
func NewDumpsysPowerParser() *DumpsysPowerParser {
	return &DumpsysPowerParser{
		BaseParser: NewBaseParser("dumpsys power", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys power" command
func (p *DumpsysPowerParser) Parse(output string) (*model.DumpsysPowerResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysPowerSection{}
	
	var currentSection model.DumpsysPowerSection
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
			currentSection = model.DumpsysPowerSection{
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
	
	response := &model.DumpsysPowerResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysPowerParser) Validate(result *model.DumpsysPowerResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys power output")
	}
	
	return nil
}

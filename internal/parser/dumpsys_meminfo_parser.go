package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysMeminfoParser parses adb shell dumpsys meminfo output
type DumpsysMeminfoParser struct {
	*BaseParser
}

// NewDumpsysMeminfoParser creates a new dumpsys meminfo parser
func NewDumpsysMeminfoParser() *DumpsysMeminfoParser {
	return &DumpsysMeminfoParser{
		BaseParser: NewBaseParser("dumpsys meminfo", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys meminfo" command
func (p *DumpsysMeminfoParser) Parse(output string) (*model.DumpsysMeminfoResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysMeminfoSection{}
	
	var currentSection model.DumpsysMeminfoSection
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
			currentSection = model.DumpsysMeminfoSection{
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
	
	response := &model.DumpsysMeminfoResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysMeminfoParser) Validate(result *model.DumpsysMeminfoResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys meminfo output")
	}
	
	return nil
}

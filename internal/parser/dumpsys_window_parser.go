package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysWindowParser parses adb shell dumpsys window output
type DumpsysWindowParser struct {
	*BaseParser
}

// NewDumpsysWindowParser creates a new dumpsys window parser
func NewDumpsysWindowParser() *DumpsysWindowParser {
	return &DumpsysWindowParser{
		BaseParser: NewBaseParser("dumpsys window", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys window" command
func (p *DumpsysWindowParser) Parse(output string) (*model.DumpsysWindowResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysWindowSection{}
	
	var currentSection model.DumpsysWindowSection
	var contentLines []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check for section headers (lines in all caps with parentheses)
		if strings.Contains(line, "WINDOW MANAGER") && strings.Contains(line, "(") && strings.Contains(line, ")") {
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
				currentSection = model.DumpsysWindowSection{
					Name:    sectionName,
					Content: "",
				}
			}
		} else if strings.HasSuffix(line, ":") && line != ":" && !strings.Contains(line, "WINDOW MANAGER") {
			// Check for other section headers (lines ending with colon)
			// Save previous section if exists
			if currentSection.Name != "" {
				currentSection.Content = strings.Join(contentLines, "\n")
				sections = append(sections, currentSection)
				contentLines = []string{}
			}
			
			// Extract section name (remove trailing colon)
			sectionName := strings.TrimSuffix(line, ":")
			currentSection = model.DumpsysWindowSection{
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
	
	response := &model.DumpsysWindowResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysWindowParser) Validate(result *model.DumpsysWindowResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys window output")
	}
	
	return nil
}

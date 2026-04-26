package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysConnectivityParser parses adb shell dumpsys connectivity output
type DumpsysConnectivityParser struct {
	*BaseParser
}

// NewDumpsysConnectivityParser creates a new dumpsys connectivity parser
func NewDumpsysConnectivityParser() *DumpsysConnectivityParser {
	return &DumpsysConnectivityParser{
		BaseParser: NewBaseParser("dumpsys connectivity", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys connectivity" command
func (p *DumpsysConnectivityParser) Parse(output string) (*model.DumpsysConnectivityResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysConnectivitySection{}
	
	var currentSection model.DumpsysConnectivitySection
	var contentLines []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check for section headers (lines that end with colon and are not empty)
		if strings.HasSuffix(line, ":") && line != ":" {
			// Save previous section if exists
			if currentSection.Name != "" {
				currentSection.Content = strings.Join(contentLines, "\n")
				sections = append(sections, currentSection)
				contentLines = []string{}
			}
			
			// Extract section name (remove trailing colon)
			sectionName := strings.TrimSuffix(line, ":")
			currentSection = model.DumpsysConnectivitySection{
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
	
	response := &model.DumpsysConnectivityResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysConnectivityParser) Validate(result *model.DumpsysConnectivityResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys connectivity output")
	}
	
	return nil
}

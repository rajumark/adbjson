package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysWifiParser parses adb shell dumpsys wifi output
type DumpsysWifiParser struct {
	*BaseParser
}

// NewDumpsysWifiParser creates a new dumpsys wifi parser
func NewDumpsysWifiParser() *DumpsysWifiParser {
	return &DumpsysWifiParser{
		BaseParser: NewBaseParser("dumpsys wifi", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys wifi" command
func (p *DumpsysWifiParser) Parse(output string) (*model.DumpsysWifiResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysWifiSection{}
	
	var currentSection model.DumpsysWifiSection
	var contentLines []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check for section headers (lines starting with "Dump of")
		if strings.HasPrefix(line, "Dump of ") {
			// Save previous section if exists
			if currentSection.Name != "" {
				currentSection.Content = strings.Join(contentLines, "\n")
				sections = append(sections, currentSection)
				contentLines = []string{}
			}
			
			// Extract section name
			sectionName := strings.TrimPrefix(line, "Dump of ")
			currentSection = model.DumpsysWifiSection{
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
	
	response := &model.DumpsysWifiResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysWifiParser) Validate(result *model.DumpsysWifiResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys wifi output")
	}
	
	return nil
}

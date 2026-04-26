package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysTelephonyRegistryParser parses adb shell dumpsys telephony.registry output
type DumpsysTelephonyRegistryParser struct {
	*BaseParser
}

// NewDumpsysTelephonyRegistryParser creates a new dumpsys telephony.registry parser
func NewDumpsysTelephonyRegistryParser() *DumpsysTelephonyRegistryParser {
	return &DumpsysTelephonyRegistryParser{
		BaseParser: NewBaseParser("dumpsys telephony.registry", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys telephony.registry" command
func (p *DumpsysTelephonyRegistryParser) Parse(output string) (*model.DumpsysTelephonyRegistryResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysTelephonyRegistrySection{}
	
	var currentSection model.DumpsysTelephonyRegistrySection
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
			currentSection = model.DumpsysTelephonyRegistrySection{
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
	
	response := &model.DumpsysTelephonyRegistryResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysTelephonyRegistryParser) Validate(result *model.DumpsysTelephonyRegistryResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys telephony.registry output")
	}
	
	return nil
}

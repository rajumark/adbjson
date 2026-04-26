package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// DumpsysCpuinfoParser parses adb shell dumpsys cpuinfo output
type DumpsysCpuinfoParser struct {
	*BaseParser
}

// NewDumpsysCpuinfoParser creates a new dumpsys cpuinfo parser
func NewDumpsysCpuinfoParser() *DumpsysCpuinfoParser {
	return &DumpsysCpuinfoParser{
		BaseParser: NewBaseParser("dumpsys cpuinfo", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell dumpsys cpuinfo" command
func (p *DumpsysCpuinfoParser) Parse(output string) (*model.DumpsysCpuinfoResponse, error) {
	lines := strings.Split(output, "\n")
	sections := []model.DumpsysCpuinfoSection{}
	
	var currentSection model.DumpsysCpuinfoSection
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
			currentSection = model.DumpsysCpuinfoSection{
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
	
	response := &model.DumpsysCpuinfoResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *DumpsysCpuinfoParser) Validate(result *model.DumpsysCpuinfoResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in dumpsys cpuinfo output")
	}
	
	return nil
}

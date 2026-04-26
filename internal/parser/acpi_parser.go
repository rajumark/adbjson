package parser

import (
	"adbjson/internal/model"
	"fmt"
	"strings"
)

// AcpiParser parses adb shell acpi output
type AcpiParser struct {
	*BaseParser
}

// NewAcpiParser creates a new acpi parser
func NewAcpiParser() *AcpiParser {
	return &AcpiParser{
		BaseParser: NewBaseParser("acpi", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell acpi" command
func (p *AcpiParser) Parse(output string) (*model.AcpiResponse, error) {
	lines := strings.Split(output, "\n")
	
	// Initialize sections
	coolingLines := []string{}
	thermalLines := []string{}
	errorLines := []string{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Categorize lines
		if strings.HasPrefix(line, "Cooling ") {
			coolingLines = append(coolingLines, line)
		} else if strings.HasPrefix(line, "Thermal ") {
			thermalLines = append(thermalLines, line)
		} else if strings.HasPrefix(line, "acpi:") {
			errorLines = append(errorLines, line)
		}
	}
	
	sections := []model.AcpiSection{}
	
	// Add cooling devices section if there are cooling lines
	if len(coolingLines) > 0 {
		sections = append(sections, model.AcpiSection{
			Name:    "Cooling Devices",
			Content: strings.Join(coolingLines, "\n"),
		})
	}
	
	// Add thermal information section if there are thermal lines
	if len(thermalLines) > 0 {
		sections = append(sections, model.AcpiSection{
			Name:    "Thermal Information",
			Content: strings.Join(thermalLines, "\n"),
		})
	}
	
	// Add permission errors section if there are error lines
	if len(errorLines) > 0 {
		sections = append(sections, model.AcpiSection{
			Name:    "Permission Errors",
			Content: strings.Join(errorLines, "\n"),
		})
	}
	
	response := &model.AcpiResponse{
		Sections: sections,
		Count:    len(sections),
	}
	
	return response, nil
}

// Validate checks if the parsed result is valid
func (p *AcpiParser) Validate(result *model.AcpiResponse) error {
	// Check if sections were parsed
	if len(result.Sections) == 0 {
		return fmt.Errorf("no sections found in acpi output")
	}
	
	return nil
}

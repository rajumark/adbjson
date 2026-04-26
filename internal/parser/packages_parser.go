package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PackagesParser parses adb shell pm list packages output
type PackagesParser struct {
	*BaseParser
}

// NewPackagesParser creates a new packages parser
func NewPackagesParser() *PackagesParser {
	return &PackagesParser{
		BaseParser: NewBaseParser("packages", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm list packages" command
func (p *PackagesParser) Parse(output string) (*model.PackagesResponse, error) {
	lines := strings.Split(output, "\n")
	packages := []model.Package{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Parse package line (format: package:com.example.app)
		if strings.HasPrefix(line, "package:") {
			packageName := strings.TrimPrefix(line, "package:")
			packageName = strings.TrimSpace(packageName)
			
			if packageName != "" {
				packages = append(packages, model.Package{Name: packageName})
			}
		}
	}
	
	return &model.PackagesResponse{
		Packages: packages,
		Count:    len(packages),
	}, nil
}

// Validate checks if the parsed result is valid
func (p *PackagesParser) Validate(result *model.PackagesResponse) error {
	// Packages response is always valid even with empty list
	return nil
}

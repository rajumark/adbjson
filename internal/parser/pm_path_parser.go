package parser

import (
	"adbjson/internal/model"
	"strings"
)

// PmPathParser parses adb shell pm path output
type PmPathParser struct {
	*BaseParser
}

// NewPmPathParser creates a new pm path parser
func NewPmPathParser() *PmPathParser {
	return &PmPathParser{
		BaseParser: NewBaseParser("pm_path", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell pm path" command
func (p *PmPathParser) Parse(output string) (*model.PackagePathResponse, error) {
	output = strings.TrimSpace(output)
	
	// Parse output format: package:/data/app/com.example/base.apk
	parts := strings.SplitN(output, ":", 2)
	
	packagePath := model.PackagePath{
		Package: "unknown",
		Path:    "unknown",
	}
	
	if len(parts) == 2 {
		packagePath.Package = strings.TrimSpace(parts[0])
		packagePath.Path = strings.TrimSpace(parts[1])
	}
	
	return &model.PackagePathResponse{
		PackagePath: packagePath,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *PmPathParser) Validate(result *model.PackagePathResponse) error {
	// Package path response is always valid
	return nil
}

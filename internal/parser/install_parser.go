package parser

import (
	"adbjson/internal/model"
	"strings"
)

// InstallParser parses adb install output
type InstallParser struct {
	*BaseParser
}

// NewInstallParser creates a new install parser
func NewInstallParser() *InstallParser {
	return &InstallParser{
		BaseParser: NewBaseParser("install", "1.0.0"),
	}
}

// Parse parses the raw output from "adb install" command
func (p *InstallParser) Parse(output string) (*model.InstallResponse, error) {
	output = strings.TrimSpace(output)
	
	installResult := model.InstallResult{
		Success: false,
		Message: output,
	}
	
	// Check for success indicators
	if strings.Contains(output, "Success") {
		installResult.Success = true
		installResult.Message = "Installation successful"
	} else if strings.Contains(output, "Failure") {
		installResult.Success = false
		installResult.Message = "Installation failed: " + output
	}
	
	return &model.InstallResponse{
		InstallResult: installResult,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *InstallParser) Validate(result *model.InstallResponse) error {
	// Install response is always valid
	return nil
}

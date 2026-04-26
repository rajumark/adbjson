package parser

import (
	"adbjson/internal/model"
	"strings"
)

// UninstallParser parses adb uninstall output
type UninstallParser struct {
	*BaseParser
}

// NewUninstallParser creates a new uninstall parser
func NewUninstallParser() *UninstallParser {
	return &UninstallParser{
		BaseParser: NewBaseParser("uninstall", "1.0.0"),
	}
}

// Parse parses the raw output from "adb uninstall" command
func (p *UninstallParser) Parse(output string) (*model.UninstallResponse, error) {
	output = strings.TrimSpace(output)
	
	uninstallResult := model.UninstallResult{
		Success: false,
		Message: output,
	}
	
	// Check for success indicators
	if strings.Contains(output, "Success") {
		uninstallResult.Success = true
		uninstallResult.Message = "Uninstallation successful"
	} else if strings.Contains(output, "Failure") {
		uninstallResult.Success = false
		uninstallResult.Message = "Uninstallation failed: " + output
	}
	
	return &model.UninstallResponse{
		UninstallResult: uninstallResult,
	}, nil
}

// Validate checks if the parsed result is valid
func (p *UninstallParser) Validate(result *model.UninstallResponse) error {
	// Uninstall response is always valid
	return nil
}

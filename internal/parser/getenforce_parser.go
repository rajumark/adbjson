package parser

import (
	"adbjson/internal/model"
	"strings"
)

// GetenforceParser parses adb shell getenforce output
type GetenforceParser struct {
	*BaseParser
}

// NewGetenforceParser creates a new getenforce parser
func NewGetenforceParser() *GetenforceParser {
	return &GetenforceParser{
		BaseParser: NewBaseParser("getenforce", "1.0.0"),
	}
}

// Parse parses the raw output from "adb shell getenforce" command
func (p *GetenforceParser) Parse(output string) (*model.SELinuxStatusResponse, error) {
	status := strings.TrimSpace(output)
	
	return &model.SELinuxStatusResponse{
		SELinuxStatus: model.SELinuxStatus{Status: status},
	}, nil
}

// Validate checks if the parsed result is valid
func (p *GetenforceParser) Validate(result *model.SELinuxStatusResponse) error {
	// SELinux status response is always valid
	return nil
}

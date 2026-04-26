package parser

import (
	"adbjson/internal/model"
	"strings"
)

// DevPathParser parses adb get-devpath output
type DevPathParser struct{}

// NewDevPathParser creates a new device path parser
func NewDevPathParser() *DevPathParser {
	return &DevPathParser{}
}

// Parse parses the raw output from "adb get-devpath" command
func (p *DevPathParser) Parse(output string) (*model.DevPathResponse, error) {
	devPath := strings.TrimSpace(output)
	
	return &model.DevPathResponse{
		DevPath: devPath,
	}, nil
}

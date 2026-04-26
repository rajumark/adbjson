package parser

import (
	"adbjson/internal/model"
	"strings"
)

// ScreenDensityParser parses adb shell wm density output
type ScreenDensityParser struct{}

// NewScreenDensityParser creates a new screen density parser
func NewScreenDensityParser() *ScreenDensityParser {
	return &ScreenDensityParser{}
}

// Parse parses the raw output from "adb shell wm density" command
func (p *ScreenDensityParser) Parse(output string) (*model.ScreenDensityResponse, error) {
	// Output format: "Physical density: 390"
	parts := strings.Split(output, ":")
	if len(parts) < 2 {
		return &model.ScreenDensityResponse{
			PhysicalDensity: "",
		}, nil
	}
	
	physicalDensity := strings.TrimSpace(parts[1])
	
	return &model.ScreenDensityResponse{
		PhysicalDensity: physicalDensity,
	}, nil
}

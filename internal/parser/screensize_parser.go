package parser

import (
	"adbjson/internal/model"
	"strings"
)

// ScreenSizeParser parses adb shell wm size output
type ScreenSizeParser struct{}

// NewScreenSizeParser creates a new screen size parser
func NewScreenSizeParser() *ScreenSizeParser {
	return &ScreenSizeParser{}
}

// Parse parses the raw output from "adb shell wm size" command
func (p *ScreenSizeParser) Parse(output string) (*model.ScreenSizeResponse, error) {
	// Output format: "Physical size: 1080x2400"
	parts := strings.Split(output, ":")
	if len(parts) < 2 {
		return &model.ScreenSizeResponse{
			PhysicalSize: "",
		}, nil
	}
	
	physicalSize := strings.TrimSpace(parts[1])
	
	return &model.ScreenSizeResponse{
		PhysicalSize: physicalSize,
	}, nil
}

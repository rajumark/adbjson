package parser

import (
	"adbjson/internal/model"
	"strings"
)

// DevicesParser parses adb devices output
type DevicesParser struct{}

// NewDevicesParser creates a new devices parser
func NewDevicesParser() *DevicesParser {
	return &DevicesParser{}
}

// Parse parses the raw output from "adb devices" command
func (p *DevicesParser) Parse(output string) (*model.DevicesResponse, error) {
	lines := strings.Split(output, "\n")
	devices := []model.Device{}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Skip header line
		if strings.HasPrefix(line, "List of devices attached") {
			continue
		}
		
		// Split by tab
		parts := strings.Split(line, "\t")
		
		// Handle malformed lines safely
		if len(parts) < 2 {
			continue
		}
		
		device := model.Device{
			ID:     strings.TrimSpace(parts[0]),
			Status: strings.TrimSpace(parts[1]),
		}
		
		devices = append(devices, device)
	}
	
	return &model.DevicesResponse{
		Devices: devices,
	}, nil
}
